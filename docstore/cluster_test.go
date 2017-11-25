package docstore

import (
	"testing"
	"net/http"
	"time"
	"strings"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"github.com/json-iterator/go"
	"net"
	"github.com/v2pro/quokka/kvstore"
)

func Test_write_to_master(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		})
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
	resp, err := http.Post("http://127.0.0.1:2515/docstore/exec", "application/json",
		strings.NewReader(`
{
	"EntityType": "user",
	"CommandType": "create",
	"EntityId": "123"
}
		`))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
}

func Test_write_to_slave(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		})
	entityId := "123"
	partition := kvstore.HashToPartition(entityId)
	should.Nil(setPartitionServers(partition, &partitionServers{
		Master: &net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 2515,
		},
	}))
	go http.ListenAndServe("127.0.0.1:2515", Mux)
	go http.ListenAndServe("127.0.0.1:2516", Mux)
	time.Sleep(time.Second)
	resp, err := http.Post("http://127.0.0.1:2516/docstore/exec", "application/json",
		strings.NewReader(`
{
	"EntityType": "user",
	"CommandType": "create",
	"EntityId": "123"
}
		`))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(2515, jsoniter.Get(body, "hint_master", "Port").ToInt())
}
