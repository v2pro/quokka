package docstore

import (
	"github.com/v2pro/quokka/kvstore"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/spaolacci/murmur3"
	"github.com/v2pro/plz/countlog"
	"errors"
)

var eventJson = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

func loadEntity(partition uint64, entityType string, entityId string) (*entity, error) {
	eventId, partitionVersion := getEventId(partition, entityId)
	if eventId == 0 {
		countlog.Trace("event!backtrace.can not find event for entity",
			"partition", partition,
			"entityType", entityType,
			"entityId", entityId)
	}
	encodedEvent, err := kvstore.Get(partition, eventId)
	if err != nil {
		countlog.Trace("event!backtrace.event not found",
			"err", err,
			"entityType", entityType,
			"entityId", entityId,
			"partition", partition,
			"eventId", eventId)
		return nil, err
	}
	var event partialEvent
	if err = eventJson.Unmarshal(encodedEvent, &event); err != nil {
		countlog.Trace("event!backtrace.failed to unmarshal event",
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
			countlog.Error("event!backtrace.can not find event with state",
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
		baseEventId:      baseEventId,
		entityType:       entityType,
		entityId:         entityId,
		partition:        partition,
		version:          version,
		partitionVersion: partitionVersion,
		doc:              doc,
	}
	return entity, nil
}

type partialEvent struct {
	BaseEventId uint64
	Version     uint64
	State       []byte
	Delta       []byte
}

type entity struct {
	baseEventId      uint64
	entityType       string
	entityId         string
	partition        uint64
	version          uint64
	partitionVersion uint64
	doc              interface{}
}

func hashToPartition(entityId string) uint64 {
	hash := murmur3.Sum64([]byte(entityId))
	partitionIndex := hash % kvstore.PartitionsCount
	return partitionIndex
}
