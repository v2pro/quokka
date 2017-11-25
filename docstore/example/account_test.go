package example

import (
	"testing"
	"github.com/v2pro/quokka/docstore"
	"net/http"
	"time"
	"strings"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"github.com/json-iterator/go"
)

func Test_account(t *testing.T) {
	docstore.AddEntityType("Account").Handler("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}).Handler("charge",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		})
	go http.ListenAndServe("127.0.0.1:2515", docstore.Mux)
	time.Sleep(time.Second)
	post(t, "http://127.0.0.1:2515/exec", `
{
	"EntityType": "Account",
	"EntityId": "123",
	"CommandType": "create"
}
	`)
	post(t, "http://127.0.0.1:2515/exec", `
{
	"EntityType": "Account",
	"EntityId": "123",
	"CommandType": "charge",
	"CommandRequest": 1
}
	`)
}

func post(t *testing.T, url string, req string) {
	should := require.New(t)
	resp, err := http.Post(url, "application/json", strings.NewReader(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
}