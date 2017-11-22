package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
)

func Test_create_object(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return "hello"
		})
	resp := Exec("user", "create", newID().String(), []byte(`{}`))
	should.Equal(0, jsoniter.Get(resp, "errno").ToInt())
	should.Equal("hello", jsoniter.Get(resp, "data").ToString())
}

func reset(entityType string) *Store {
	stores = map[string]*Store{}
	return AddStore(entityType)
}
