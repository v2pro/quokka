package docstore

import (
	"time"
	"github.com/v2pro/quokka/docstore/runtime"
	"errors"
	"github.com/v2pro/quokka/kvstore"
	"github.com/v2pro/plz/countlog"
	"github.com/json-iterator/go"
	"encoding/json"
)

var commandProcessors = make([]*commandProcessor, kvstore.PartitionsCount)

type commandProcessor struct {
	partition     uint64
	reqChan       chan *command
	lastEventId   uint64
	entityLookup  *entityLookup
	commandLookup *commandLookup
	lookupSyncer  *eventProcessor
}

type command struct {
	EntityType     string
	CommandType    string
	EntityId       string
	CommandId      string
	CommandRequest json.RawMessage
	respChan       chan []byte
}

type entityCache struct {
	startVersion uint64 // the cache covers modification since this event id
	entities     map[string]*entity
}

type commandCache struct {
	startVersion uint64
	responses    map[string][]byte
}

func init() {
	for i := 0; i < len(commandProcessors); i++ {
		partition := uint64(i)
		processor :=&commandProcessor{
			partition:     partition,
			reqChan:       make(chan *command),
			entityLookup:  newEntityLookup(),
			commandLookup: newCommandLookup(),
		}
		processor.lookupSyncer = newEventProcessor(partition, processor)
		go processor.executeInBackground()
		go processor.lookupSyncer.syncInBackground()
		commandProcessors[i] = processor
	}
}

func (processor *commandProcessor) executeInBackground() {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!command_processor.executeInBackground.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	for {
		processor.executeOne()
	}
}

func (processor *commandProcessor) executeOne() {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!command_processor.executeOne.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	command := <-processor.reqChan
	resp := processor.exec(command)
	command.respChan <- resp
}

func (processor *commandProcessor) delegatedExec(execReq *command, timeout time.Duration) []byte {
	execReq.respChan = make(chan []byte, 1)
	timer := time.NewTimer(timeout)
	select {
	case processor.reqChan <- execReq:
		resp := <-execReq.respChan
		return resp
	case <-timer.C:
		return replyError(errors.New("timeout"))
	}
}

func (processor *commandProcessor) init(execReq *command) {

}

func (processor *commandProcessor) exec(execReq *command) []byte {
	partition := processor.partition
	entityId := execReq.EntityId
	req := execReq.CommandRequest
	entityType := execReq.EntityType
	commandType := execReq.CommandType
	commandId := execReq.CommandId
	var reqObj interface{}
	var err error
	if len(req) > 0 {
		err = runtime.Json.Unmarshal(req, &reqObj)
		if err != nil {
			return replyError(err)
		}
	}
	store := getEntityType(entityType)
	if store == nil {
		return replyError(errors.New("store not defined for entity type " + entityType))
	}
	handler := store.getHandler(commandType)
	if store == nil {
		return replyError(errors.New("handler not defined for command type " + commandType))
	}
	var ent *entity
	var version uint64
	if "create" == commandType {
		_, err := processor.entityLookup.getEntity(partition, entityType, entityId)
		if err != nil {
			return replyError(errors.New("entity with same id found"))
		}
		ent = &entity{
			eventId: 0,
			version: 1,
			doc:     runtime.NewObject(),
		}

	} else {
		ent, err = processor.entityLookup.getEntity(partition, entityType, entityId)
		if err != nil {
			return replyError(err)
		}
	}
	resp := handler(ent.doc, reqObj)
	err, _ = resp.(error)
	if err != nil {
		return replyError(err)
	}
	encodedResp, err := runtime.Json.Marshal(resp)
	if err != nil {
		return replyError(err)
	}
	event := &Event{
		EventId:         processor.lastEventId + 1,
		Partition:       partition,
		BaseEventId:     ent.eventId,
		EntityId:        entityId,
		Version:         version + 1,
		CommandId:       commandId,
		CommandType:     commandType,
		CommandRequest:  req,
		CommandResponse: encodedResp,
		CommittedAt:     time.Now().UnixNano(),
	}
	if version%16 == 0 {
		event.State, err = runtime.Json.Marshal(ent.doc)
	}
	event.Delta, err = runtime.DeltaJson.Marshal(ent.doc)
	if err != nil {
		return replyError(err)
	}
	encodedEvent, err := eventJson.Marshal(event)
	if err != nil {
		return replyError(err)
	}
	err = kvstore.Append(processor.partition, event.EventId, encodedEvent)
	if err != nil {
		return replyError(err)
	}
	countlog.Trace("event!docstore.stored event",
		"partition", processor.partition,
		"partitionVersion", event.EventId,
		"encodedEvent", encodedEvent)
	processor.lastEventId = event.EventId
	// update in memory lookup
	ent.eventId = event.EventId
	processor.entityLookup.cacheEntity(event.EntityId, ent, event.EventId)
	processor.commandLookup.cacheCommand(event.CommandId, encodedResp, event.EventId)
	// up kv store lookup in separate goroutine
	processor.lookupSyncer.enqueue(event)
	return replySuccess(encodedResp)
}


func (processor *commandProcessor) LoadOffset(partition uint64) (uint64, error) {
	offset, err := kvstore.GetMonotonic(partition, "offset")
	if err != nil {
		return 0, err
	}
	return offset, nil
}

func (processor *commandProcessor) CommitOffset(partition uint64, offset uint64) error {
	return kvstore.SetMonotonic(partition, "offset", offset)
}

func (processor *commandProcessor) Sync(event *Event) error {
	go forwardEventInBackground(event)
	err := processor.entityLookup.setEventId(event.Partition, event.EntityId, event.EventId)
	if err != nil {
		countlog.Error("event!core_event_handler.failed to update entity lookup",
			"err", err,
			"partition", event.Partition,
			"entityId", event.EntityId,
			"eventId", event.EventId)
		return err
	}
	err = processor.entityLookup.setEventId(event.Partition, event.CommandId, event.EventId)
	if err != nil {
		countlog.Error("event!core_event_handler.failed to update command lookup",
			"err", err,
			"partition", event.Partition,
			"entityId", event.EntityId,
			"eventId", event.EventId)
		return err
	}
	return nil
}

func replySuccess(encodedResp []byte) []byte {
	return append(append([]byte(`{"errno":0,"data":`), encodedResp...), '}')
}

func replyError(err error) []byte {
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteObjectStart()
	stream.WriteObjectField("errno")
	stream.WriteVal(1)
	stream.WriteMore()
	stream.WriteObjectField("errmsg")
	stream.WriteString(err.Error())
	stream.WriteObjectEnd()
	return stream.Buffer()
}
