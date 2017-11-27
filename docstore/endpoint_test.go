package docstore

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/quokka/kvstore"
	"github.com/json-iterator/go"
	"net"
	"net/http"
	"time"
	"strings"
	"io/ioutil"
)

func Test_entity_command_url(t *testing.T) {
	should := require.New(t)
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, nil, nil)
	entityId := "123"
	partition := kvstore.HashToPartition(entityId)
	should.Nil(setPartitionServers(partition, &partitionServers{
		Master: &net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 2515,
		},
	}))
	go http.ListenAndServe("127.0.0.1:2515", Mux)
	time.Sleep(time.Second)
	resp, err := http.Post("http://127.0.0.1:2515/docstore/user/create", "application/json",
		strings.NewReader(`
{
	"EntityId": "123"
}
		`))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}
