package docstore

import "sync"

var entityTypes = map[string]*entityCommandHandlers{}
var entityTypesMutex = &sync.Mutex{}

type CommandHandler func(doc interface{}, request interface{}) (resp interface{})

type entityCommandHandlers struct {
	handlers      map[string]CommandHandler
	handlersMutex *sync.Mutex
}

func (store *entityCommandHandlers) Handler(commandType string, handler CommandHandler) *entityCommandHandlers {
	store.handlersMutex.Lock()
	defer store.handlersMutex.Unlock()
	store.handlers[commandType] = handler
	return store
}

func (store *entityCommandHandlers) getHandler(commandType string) CommandHandler {
	store.handlersMutex.Lock()
	defer store.handlersMutex.Unlock()
	return store.handlers[commandType]
}

func AddEntityType(entityType string) *entityCommandHandlers {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	if entityTypes[entityType] != nil {
		panic("already has store for " + entityType)
	}
	store := &entityCommandHandlers{
		handlers:      map[string]CommandHandler{},
		handlersMutex: &sync.Mutex{},
	}
	entityTypes[entityType] = store
	return store
}

func getEntityType(entityType string) *entityCommandHandlers {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	return entityTypes[entityType]
}