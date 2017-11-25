package docstore

import "sync"

var entityTypes = map[string]*entityType{}
var entityTypesMutex = &sync.Mutex{}

type CommandHandler func(doc interface{}, request interface{}) (resp interface{})

type entityType struct {
	handlers      map[string]CommandHandler
	handlersMutex *sync.Mutex
}

func (store *entityType) Handler(commandType string, handler CommandHandler) *entityType {
	store.handlersMutex.Lock()
	defer store.handlersMutex.Unlock()
	store.handlers[commandType] = handler
	return store
}

func (store *entityType) getHandler(commandType string) CommandHandler {
	store.handlersMutex.Lock()
	defer store.handlersMutex.Unlock()
	return store.handlers[commandType]
}

func AddEntityType(entityType string) *entityType {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	if entityTypes[entityType] != nil {
		panic("already has store for " + entityType)
	}
	store := &entityType{
		handlers:      map[string]CommandHandler{},
		handlersMutex: &sync.Mutex{},
	}
	entityTypes[entityType] = store
	return store
}

func getEntityType(entityType string) *entityType {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	return entityTypes[entityType]
}