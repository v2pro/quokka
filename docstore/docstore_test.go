package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/quokka/kvstore"
	"context"
	"net"
)

var fakeMasterAddr = &net.TCPAddr{
	IP: net.IPv4(127, 127,127, 127),
	Port: 127,
}

func Test_create_object(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return "hello"
		})
	entityId := newID().String()
	resp := exec(context.TODO(),fakeMasterAddr, &execRequest{
		EntityType: "user",
		CommandType: "create",
		EntityId: entityId,
	})
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

	entityId := newID().String()
	resp := exec(context.TODO(), fakeMasterAddr, &execRequest{
		EntityType: "user",
		CommandType: "create",
		EntityId: entityId,
	})
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	resp = exec(context.TODO(), fakeMasterAddr,&execRequest{
		EntityType: "user",
		CommandType: "get",
		EntityId: entityId,
	})
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