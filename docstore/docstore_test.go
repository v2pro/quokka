package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/quokka/kvstore"
)

func Test_create_object(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return "hello"
		})
	entityId := newID().String()
	resp := exec("user", "create", entityId, "", []byte(`{}`))
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	should.Equal("hello", jsoniter.Get(resp, "data").ToString())
	partition := hashToPartition(entityId)
	should.True(jsoniter.Get(debugGet(partition, 1), "State").ToBool())
	eventId, _ := getEventId(partition, entityId)
	should.Equal(uint64(1), eventId)
}

func Test_get_object(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			runtime.AsObj("", doc).Set("hello", "world")
			return nil
		}).Handler("get",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return doc
		})

	docId := newID().String()
	resp := exec("user", "create", docId, "", []byte(`null`))
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	resp = exec("user", "get", docId, "", []byte(`null`))
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	should.Equal("world", jsoniter.Get(resp, "data", "hello").ToString())
}

func reset(entityType string) *Store {
	resetMemKVStore()
	stores = map[string]*Store{}
	return AddStore(entityType)
}

func debugGet(partition uint64, rowKey uint64) []byte {
	row, err := kvstore.Get(partition, rowKey)
	if err != nil {
		panic(err)
	}
	return row
}