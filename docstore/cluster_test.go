package docstore

import (
	"testing"
	"net/http"
	"time"
	"strings"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"github.com/json-iterator/go"
)

func Test_write_to_master(t *testing.T) {
	should := require.New(t)
	reset("user").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
		return nil
	})
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
	should.Equal(0, jsoniter.Get(body, "errno").ToInt())
}