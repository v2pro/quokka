package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/plz/countlog"
	"errors"
)

var eventJson = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

type entity struct {
	// is baseEventId before command, will be eventId after command
	eventId uint64
	version uint64
	doc     interface{}
}

type eventForLoadEntity struct {
	BaseEventId uint64 `json:"b"`
	Version     uint64 `json:"v"`
	State       []byte `json:"s"`
	Delta       []byte `json:"d"`
}

type eventForGetCommandResponse struct {
	CommandResponse []byte `json:"p"`
}

func loadEntity(partition uint64, entityType string, entityId string, eventId uint64) (*entity, error) {
	encodedEvent, err := kvstore.Get(partition, eventId)
	if err != nil {
		countlog.Trace("event!event_log.event not found",
			"err", err,
			"entityCommandHandlers", entityType,
			"entityId", entityId,
			"partition", partition,
			"eventId", eventId)
		return nil, err
	}
	var event eventForLoadEntity
	if err = eventJson.Unmarshal(encodedEvent, &event); err != nil {
		countlog.Trace("event!event_log.failed to unmarshal event",
			"err", err,
			"partition", partition,
			"eventId", eventId,
			"encodedEvent", encodedEvent)
		return nil, err
	}
	version := event.Version
	baseEventId := eventId
	var deltas [][]byte
	for event.State == nil {
		deltas = append(deltas, event.Delta)
		if event.BaseEventId == 0 {
			countlog.Error("event!event_log.can not find event with state",
				"partition", partition,
				"baseEventId", baseEventId,
				"deltas", deltas)
			return nil, errors.New("state not found")
		}
		encodedEvent, err = kvstore.Get(partition, event.BaseEventId)
		if err != nil {
			return nil, err
		}
		event.BaseEventId = 0
		event.State = nil
		event.Delta = nil
		if err = eventJson.Unmarshal(encodedEvent, &event); err != nil {
			return nil, err
		}
	}
	var doc interface{}
	if err = runtime.Json.Unmarshal(event.State, &doc); err != nil {
		return nil, err
	}
	// apply delta backwards
	for i := len(deltas) - 1; i >= 0; i-- {
		delta := deltas[i]
		if err = runtime.DeltaJson.Unmarshal(delta, &doc); err != nil {
			return nil, err
		}
	}
	entity := &entity{
		eventId: baseEventId,
		version: version,
		doc:     doc,
	}
	return entity, nil
}

// scanEvents return events in range [from, to)
func scanEvents(partition uint64, from uint64, to uint64) ([]*Event, error) {
	if to <= from {
		return []*Event{}, nil
	}
	iter, err := kvstore.Scan(partition, from)
	if err != nil {
		return nil, err
	}
	events := make([]*Event, 0, 10)
	for {
		rows, err := iter()
		if err != nil {
			return nil, err
		}
		if rows == nil {
			break
		}
		for _, row := range rows {
			eventId := row.Key
			if eventId >= to {
				return events, nil
			}
			var event Event
			err := eventJson.Unmarshal(row.Value, &event)
			if err != nil {
				return nil, err
			}
			event.EventId = eventId
			event.Partition = partition
			events = append(events, &event)
		}
	}
	if len(events) == 0 {
		return nil, errors.New("incomplete scan")
	}
	if events[len(events)-1].EventId != to-1 {
		return nil, errors.New("incomplete scan")
	}
	return events, nil
}

func getCommandResponse(partition uint64, eventId uint64) ([]byte, error) {
	encodedEvent, err := kvstore.Get(partition, eventId)
	if err != nil {
		return nil, err
	}
	var event eventForGetCommandResponse
	err = eventJson.Unmarshal(encodedEvent, &event)
	if err != nil {
		return nil, err
	}
	return event.CommandResponse, err
}
