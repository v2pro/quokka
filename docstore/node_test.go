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
	"github.com/v2pro/quokka/kvstore"
	"os"
	"context"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	if thisNodeExecutor != nil {
		StopNode(context.TODO())
	}
	os.Exit(exitCode)
}

func Test_first_event(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, nil, nil)
	StartNode(context.TODO(), "127.0.0.1:2515")
	time.Sleep(time.Second)
	execAndExpectSuccess(t, "http://127.0.0.1:2515/docstore/user/create", "EntityId", "123")
}

func Test_duplicated_entity(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, nil, nil)
	StartNode(context.TODO(),"127.0.0.1:2515")
	time.Sleep(time.Second)
	execAndExpectSuccess(t, "http://127.0.0.1:2515/docstore/user/create", "EntityId", "123")
	StopNode(context.TODO())
	StartNode(context.TODO(),"127.0.0.1:2515")
	execAndExpectError(t, "http://127.0.0.1:2515/docstore/user/create", ErrDuplicatedEntity,
		"EntityId", "123")
}

func Test_cluster_commit_failure(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, nil, nil)
	StartNode(context.TODO(),"127.0.0.1:2515")
	time.Sleep(time.Second)
	kvstore.Append(context.TODO(), kvstore.HashToPartition("123"), "user", 1, []byte{})
	execAndExpectError(t, "http://127.0.0.1:2515/docstore/user/create", ErrEventLogConflict,
		"EntityId", "123")
	time.Sleep(time.Second)
	execAndExpectSuccess(t, "http://127.0.0.1:2515/docstore/user/create",
		"EntityId", "123")
}

func execAndExpectSuccess(t *testing.T, url string, kv ...interface{}) {
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

func execAndExpectError(t *testing.T, url string, errorNumber int, kv ...interface{}) {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal(errorNumber, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}
