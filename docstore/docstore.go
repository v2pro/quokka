package docstore

import (
	"time"
	"sync"
	"github.com/v2pro/quokka/docstore/runtime"
	"errors"
)

type event struct {
	EventId     uint64
	BaseEventId uint64
	EntityId    string
	Version     uint64
	CommandId   string
	CommandName string
	Request     []byte
	Response    []byte
	State       []byte
	Delta       []byte
	CommittedAt time.Time
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

func Exec(entityType string, commandType string, entityId string, req []byte) (resp []byte) {
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
	if "create" == commandType {
		doc := runtime.NewObject()
		resp := handler(doc, reqObj)
		err, _ := resp.(error)
		if err != nil {
			return replyError(err)
		}
		return replySuccess(resp)
	}
	return nil
}

func replySuccess(resp interface{}) []byte {
	stream := runtime.Json.BorrowStream(nil)
	defer runtime.Json.ReturnStream(stream)
	stream.WriteObjectStart()
	stream.WriteObjectField("errno")
	stream.WriteInt(0)
	stream.WriteMore()
	stream.WriteObjectField("data")
	stream.WriteVal(resp)
	stream.WriteObjectEnd()
	return stream.Buffer()
}

func replyError(err error) []byte {
	panic(err)
}
