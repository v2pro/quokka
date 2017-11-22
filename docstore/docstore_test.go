package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/docstore/runtime"
)

func Test_create_object(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return "hello"
		})
	resp := Exec("user", "create", newID().String(), "", []byte(`{}`))
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	should.Equal("hello", jsoniter.Get(resp, "data").ToString())
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
	Exec("user", "create", docId, "", []byte(`null`))
	resp := Exec("user", "get", docId, "", []byte(`null`))
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	should.Equal("hello", jsoniter.Get(resp, "data").ToString())
}

func reset(entityType string) *Store {
	resetMemKVStore()
	stores = map[string]*Store{}
	return AddStore(entityType)
}
