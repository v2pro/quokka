package docstore

import (
	"time"
	"sync"
	"github.com/v2pro/quokka/docstore/runtime"
	"errors"
	"github.com/v2pro/quokka/kvstore"
)

type event struct {
	BaseEventId uint64
	EntityId    string
	Version     uint64
	CommandId   string
	CommandType string
	Request     []byte
	Response    []byte
	State       []byte
	Delta       []byte
	CommittedAt int64
}

type Handler func(doc interface{}, request interface{}) (resp interface{})

var stores = map[string]*Store{}
var storesMutex = &sync.Mutex{}

type Store struct {
	handlers      map[string]Handler
	handlersMutex *sync.Mutex
}

func (store *Store) Handler(commandType string, handler Handler) *Store {
	store.handlersMutex.Lock()
	defer store.handlersMutex.Unlock()
	store.handlers[commandType] = handler
	return store
}

func (store *Store) getHandler(commandType string) Handler {
	store.handlersMutex.Lock()
	defer store.handlersMutex.Unlock()
	return store.handlers[commandType]
}

func AddStore(entityType string) *Store {
	storesMutex.Lock()
	defer storesMutex.Unlock()
	if stores[entityType] != nil {
		panic("already has store for " + entityType)
	}
	store := &Store{
		handlers:      map[string]Handler{},
		handlersMutex: &sync.Mutex{},
	}
	stores[entityType] = store
	return store
}

func getStore(entityType string) *Store {
	storesMutex.Lock()
	defer storesMutex.Unlock()
	return stores[entityType]
}

func Exec(entityType string, commandType string, entityId string, commandId string, req []byte) []byte {
	var reqObj interface{}
	err := runtime.Json.Unmarshal(req, &reqObj)
	if err != nil {
		return replyError(err)
	}
	store := getStore(entityType)
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
		partition := hashToPartition(entityId)
		eventId, partitionVersion := getEventId(partition, entityId)
		if eventId != 0 {
			return replyError(errors.New("entity with same id found"))
		}
		ent = &entity{
			baseEventId:      0,
			entityType:       entityType,
			entityId:         entityId,
			partition:        partition,
			version:          1,
			partitionVersion: partitionVersion,
			doc:              runtime.NewObject(),
		}

	} else {
		ent, err = loadEntity(entityType, entityId)
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
	event := &event{
		BaseEventId: ent.baseEventId,
		EntityId:    entityId,
		Version:     version + 1,
		CommandId:   commandId,
		CommandType: commandType,
		Request:     req,
		Response:    encodedResp,
		CommittedAt: time.Now().UnixNano(),
	}
	if version%16 == 0 {
		event.State, err = runtime.Json.Marshal(ent.doc)
	} else {
		event.Delta, err = runtime.DeltaJson.Marshal(ent.doc)
	}
	if err != nil {
		return replyError(err)
	}
	encodedEvent, err := eventJson.Marshal(event)
	if err != nil {
		return replyError(err)
	}
	err = kvstore.Append(ent.partition, ent.partitionVersion+1, encodedEvent)
	if err != nil {
		return replyError(err)
	}
	return replySuccess(encodedResp)
}

type entityId struct {
	entityType string
	entityId   string
}

func replySuccess(encodedResp []byte) []byte {
	return append(append([]byte(`{"errno":0,"data":`), encodedResp...), '}')
}

func replyError(err error) []byte {
	panic(err)
}
