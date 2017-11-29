package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"net/http"
	"time"
	"io/ioutil"
	"github.com/v2pro/quokka/docstore/runtime"
	"bytes"
)

func Test_cluster_first_node(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, nil, nil)
	StartNode("127.0.0.1:2515")
	time.Sleep(time.Second)
	post(t, "http://127.0.0.1:2515/docstore/user/create", "EntityId", "123")
}

func post(t *testing.T, url string, kv ...interface{}) {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}
