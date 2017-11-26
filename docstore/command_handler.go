package docstore

import (
	"sync"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/quokka/docstore/transpiler"
	"github.com/v2pro/plz/countlog"
)

var entityTypes = map[string]*entityTypeDef{}
var entityTypesMutex = &sync.Mutex{}

type CommandHandler func(doc interface{}, request interface{}) (resp interface{})

type commandDef struct {
	requestSchema  *runtime.Schema
	responseSchema *runtime.Schema
	handler        CommandHandler
}

type entityTypeDef struct {
	schema           *runtime.Schema
	commandDefs      map[string]*commandDef
	commandDefsMutex *sync.Mutex
}

func (typ *entityTypeDef) Command(commandType string, jsHandler string, thriftIDL string) *entityTypeDef {
	handler, err := transpiler.Compile(jsHandler)
	if err != nil {
		countlog.Error("event!command_handler.failed to transpile javascript",
			"err", err, "jsHandler", jsHandler)
		return typ
	}
	schemas := runtime.ThriftSchemas(thriftIDL)
	typ.AddCommand(commandType, handler, schemas["Request"], schemas["Response"])
	return typ
}

func (typ *entityTypeDef) AddCommand(commandType string, handler CommandHandler,
	requestSchema *runtime.Schema, responseSchema *runtime.Schema) {
	typ.commandDefsMutex.Lock()
	defer typ.commandDefsMutex.Unlock()
	typ.commandDefs[commandType] = &commandDef{
		requestSchema:  requestSchema,
		responseSchema: responseSchema,
		handler:        handler,
	}
}

func (typ *entityTypeDef) getCommandDef(commandType string) *commandDef {
	typ.commandDefsMutex.Lock()
	defer typ.commandDefsMutex.Unlock()
	return typ.commandDefs[commandType]
}

func Entity(entityType string, thriftIDL string) *entityTypeDef {
	return AddEntity(entityType, runtime.ThriftSchemas(thriftIDL)["Doc"])
}

func AddEntity(entityType string, schema *runtime.Schema) *entityTypeDef {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	if entityTypes[entityType] != nil {
		panic("already has store for " + entityType)
	}
	typ := &entityTypeDef{
		schema:           schema,
		commandDefs:      map[string]*commandDef{},
		commandDefsMutex: &sync.Mutex{},
	}
	entityTypes[entityType] = typ
	return typ
}

func getEntityType(entityType string) *entityTypeDef {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	return entityTypes[entityType]
}
