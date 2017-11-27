package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/quokka/kvstore"
	"github.com/rs/xid"
	"github.com/v2pro/quokka/kvstore/memkv"
)

func Test_create_object(t *testing.T) {
	should := require.New(t)
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return "hello"
		}, nil, nil)
	entityId := xid.New().String()
	processor := newCommandProcessor(0)
	resp := processor.exec(&command{
		EntityType:  "user",
		CommandType: "create",
		EntityId:    entityId,
	})
	should.Equal("", jsoniter.Get(resp, "errmsg").ToString())
	should.Equal("hello", jsoniter.Get(resp, "data").ToString())
	should.True(jsoniter.Get(debugGet(0, 1), "s").ToBool())
}

func Test_get_object(t *testing.T) {
	should := require.New(t)
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			runtime.AsObj(doc).Set("hello", "world")
			return nil
		}, nil, nil).AddCommand("get",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return doc
		}, nil, nil)

	entityId := xid.New().String()
	processor := newCommandProcessor(0)
	resp := processor.exec(&command{
		EntityType:  "user",
		CommandType: "create",
		EntityId:    entityId,
	})
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	resp = processor.exec(&command{
		EntityType:  "user",
		CommandType: "get",
		EntityId:    entityId,
	})
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	should.Equal("world", jsoniter.Get(resp, "data", "hello").ToString())
}

func reset(entityType string) *entityTypeDef {
	memkv.ResetKVStore()
	entityTypes = map[string]*entityTypeDef{}
	return AddEntity(entityType, nil)
}

func debugGet(partition uint64, rowKey uint64) []byte {
	row, err := kvstore.Get(partition, rowKey)
	if err != nil {
		panic(err)
	}
	return row
}
