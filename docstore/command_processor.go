package docstore

import (
	"time"
	"github.com/v2pro/quokka/docstore/runtime"
	"errors"
	"github.com/v2pro/quokka/kvstore"
	"github.com/v2pro/plz/countlog"
	"github.com/json-iterator/go"
	"encoding/json"
	"context"
)

const ErrUnknown = 1000
const ErrEventLogConflict = 1001 // the master is no longer in charge, should find out who is the master
const ErrDuplicatedEntity = 1002 // create entity with same entity id but with different command id
var commandProcessors = make([]map[string]*commandProcessor, kvstore.PartitionsCount)

type commandProcessor struct {
	partitionId   uint64
	entityType    string
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
	IsPromoting    bool // update topo when command executed successfully
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
	for partitionId := uint64(0); partitionId < kvstore.PartitionsCount; partitionId++ {
		commandProcessors[partitionId] = map[string]*commandProcessor{}
	}
}

func newCommandProcessor(partitionId uint64, entityType string) *commandProcessor {
	processor := &commandProcessor{
		partitionId:   partitionId,
		entityType:    entityType,
		reqChan:       make(chan *command),
		entityLookup:  newEntityLookup(),
		commandLookup: newCommandLookup(),
	}
	processor.lookupSyncer = newEventProcessor(partitionId, processor)
	return processor
}

func (processor *commandProcessor) init(ctx context.Context) error {
	offset, err := processor.entityLookup.kvstoreLookup.getOffset(ctx, processor.partitionId, processor.entityType)
	if err != nil {
		return err
	}
	iter, err := kvstore.Scan(ctx, processor.partitionId, processor.entityType, offset+1)
	if err != nil {
		return err
	}
	for {
		rows, err := iter()
		if err != nil {
			return err
		}
		if rows == nil {
			processor.lastEventId = offset
			return nil
		}
		for _, row := range rows {
			offset = row.Key
			var event Event
			err = runtime.Json.Unmarshal(row.Value, &event)
			if err != nil {
				countlog.Error("event!command_processor.failed to unmarshal event",
					"err", err,
					"encodedEvent", row.Value)
				return err
			}
			processor.commandLookup.cacheCommand(event.CommandId, event.CommandResponse, event.EventId)
			err := processor.replayEvent(ctx, &event)
			if err != nil {
				return err
			}
		}
	}
}

func (processor *commandProcessor) replayEvent(ctx context.Context, event *Event) error {
	var err error
	if event.State != nil {
		var doc = runtime.NewObject()
		err = runtime.Json.Unmarshal(event.State, doc)
		if err != nil {
			countlog.Error("event!command_processor.failed to unmarshal state",
				"err", err,
				"state", event.State)
			return err
		}
		processor.entityLookup.cacheEntity(event.EntityId, &entity{
			eventId: event.EventId,
			version: event.Version,
			doc:     doc,
		}, event.EventId)
		return nil
	}
	entity, _ := processor.entityLookup.memLookup.getCacheValue(event.EntityId).(*entity)
	if entity == nil {
		entity, err = loadEntity(ctx, processor.partitionId, processor.entityType, event.EntityId, event.BaseEventId)
		if err != nil {
			return err
		}
	}
	err = runtime.DeltaJson.Unmarshal(event.Delta, entity.doc)
	if err != nil {
		countlog.Error("event!command_processor.failed to apply delta",
			"err", err,
			"state", event.State)
		return err
	}
	processor.entityLookup.cacheEntity(event.EntityId, entity, event.EventId)
	return nil
}

func (processor *commandProcessor) executeInBackground(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			return
		}
		processor.executeOne(ctx)
	}
}

func (processor *commandProcessor) executeOne(ctx context.Context) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			countlog.Fatal("event!command_processor.executeOne.panic", "err", recovered,
				"stacktrace", countlog.ProvideStacktrace)
		}
	}()
	select {
	case command := <-processor.reqChan:
		resp := processor.exec(ctx, command)
		command.respChan <- resp
	case <-ctx.Done():
		return
	}
}

func (processor *commandProcessor) delegatedExec(cmd *command, timeout time.Duration) []byte {
	cmd.respChan = make(chan []byte, 1)
	timer := time.NewTimer(timeout)
	select {
	case processor.reqChan <- cmd:
		resp := <-cmd.respChan
		return resp
	case <-timer.C:
		return replyError(errors.New("timeout"))
	}
}

func (processor *commandProcessor) exec(ctx context.Context, cmd *command) []byte {
	partition := processor.partitionId
	entityId := cmd.EntityId
	req := cmd.CommandRequest
	entityType := cmd.EntityType
	commandType := cmd.CommandType
	commandId := cmd.CommandId
	var reqObj interface{}
	var err error
	if len(req) > 0 {
		err = runtime.Json.Unmarshal(req, &reqObj)
		if err != nil {
			return replyError(err)
		}
	} else {
		req = json.RawMessage("{}")
	}
	store := getEntityType(entityType)
	if store == nil {
		return replyError(errors.New("store not defined for entity type " + entityType))
	}
	commandDef := store.getCommandDef(commandType)
	if commandDef == nil {
		return replyError(errors.New("handler not defined for command type " + commandType))
	}
	var ent *entity
	var version uint64
	if "create" == commandType {
		_, err := processor.entityLookup.getEntity(ctx, partition, entityType, entityId)
		if err == nil {
			return replyError(withErrorNumber(errors.New("entity with same id found"), ErrDuplicatedEntity))
		}
		ent = &entity{
			eventId: 0,
			version: 1,
			doc:     runtime.NewObject(),
		}

	} else {
		ent, err = processor.entityLookup.getEntity(ctx, partition, entityType, entityId)
		if err != nil {
			return replyError(err)
		}
	}
	resp := commandDef.handler(ent.doc, reqObj)
	err, _ = resp.(error)
	if err != nil {
		return replyError(err)
	}
	encodedResp, err := runtime.Json.Marshal(resp)
	if err != nil {
		return replyError(err)
	}
	event := &Event{
		Partition:       partition,
		EntityType:      entityType,
		EventId:         processor.lastEventId + 1,
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
	err = kvstore.Append(ctx, processor.partitionId, event.EntityType, event.EventId, encodedEvent)
	if err != nil {
		if err == kvstore.KeyConflictError {
			return replyError(withErrorNumber(err, ErrEventLogConflict))
		}
		return replyError(err)
	}
	countlog.Trace("event!docstore.stored event",
		"partitionId", processor.partitionId,
		"eventId", event.EventId,
		"entityId", event.EntityId,
		"encodedEvent", encodedEvent)
	processor.lastEventId = event.EventId
	// update in memory lookup
	ent.eventId = event.EventId
	processor.entityLookup.cacheEntity(event.EntityId, ent, event.EventId)
	processor.commandLookup.cacheCommand(event.CommandId, encodedResp, event.EventId)
	// up kv store lookup in separate goroutine
	processor.lookupSyncer.enqueue(event)
	if cmd.IsPromoting {
		thisNodeExecutor.Go(func(ctx context.Context) {
			topo.savePromotedMasterInBackground(ctx, partition)
		})
	}
	return replySuccess(encodedResp)
}

func (processor *commandProcessor) LoadOffset(ctx context.Context, partitionId uint64, entityType string) (uint64, error) {
	offset, err := kvstore.GetMonotonic(ctx, partitionId, entityType, "offset")
	if err != nil {
		return 0, err
	}
	return offset, nil
}

func (processor *commandProcessor) CommitOffset(ctx context.Context, partitionId uint64, entityType string, offset uint64) error {
	return kvstore.SetMonotonic(ctx, partitionId, entityType, "offset", offset)
}

func (processor *commandProcessor) Sync(ctx context.Context, event *Event) error {
	thisNodeExecutor.Go(func(ctx context.Context) {
		forwardEventInBackground(ctx, event)
	})
	err := processor.entityLookup.setEventId(ctx, event.Partition, event.EntityType, event.EntityId, event.EventId)
	if err != nil {
		countlog.Error("event!core_event_handler.failed to update entity lookup",
			"err", err,
			"partitionId", event.Partition,
			"entityId", event.EntityId,
			"eventId", event.EventId)
		return err
	}
	err = processor.entityLookup.setEventId(ctx, event.Partition, event.EntityType, event.CommandId, event.EventId)
	if err != nil {
		countlog.Error("event!core_event_handler.failed to update command lookup",
			"err", err,
			"partitionId", event.Partition,
			"entityId", event.EntityId,
			"eventId", event.EventId)
		return err
	}
	return nil
}

func replySuccess(encodedResp []byte) []byte {
	// TODO: add handled_by
	return append(append([]byte(`{"errno":0,"data":`), encodedResp...), '}')
}

type NumberedError interface {
	error
	ErrorNumber() int
}

type numberedError struct {
	errorNumber int
	err         error
}

func (err *numberedError) Error() string {
	return err.err.Error()
}

func (err *numberedError) ErrorNumber() int {
	return err.errorNumber
}

func withErrorNumber(err error, errorNumber int) *numberedError {
	return &numberedError{err: err, errorNumber: errorNumber}
}

func replyError(err error) []byte {
	stream := jsoniter.ConfigFastest.BorrowStream(nil)
	defer jsoniter.ConfigFastest.ReturnStream(stream)
	stream.WriteObjectStart()
	stream.WriteObjectField("errno")
	numberedErr, _ := err.(NumberedError)
	if numberedErr == nil {
		stream.WriteVal(ErrUnknown)
	} else {
		stream.WriteVal(numberedErr.ErrorNumber())
	}
	stream.WriteMore()
	stream.WriteObjectField("errmsg")
	stream.WriteString(err.Error())
	stream.WriteObjectEnd()
	return stream.Buffer()
}
