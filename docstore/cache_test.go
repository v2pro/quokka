package docstore

import (
	"testing"
	"context"
	"github.com/stretchr/testify/require"
	"github.com/json-iterator/go"
	"net/http"
	"io/ioutil"
)

func Test_query_entity(t *testing.T) {
	reset("user").AddCommand("create",
		func(doc interface{}, request interface{}) (resp interface{}) {
			return nil
		}, nil, nil)
	StartNode(context.TODO(), "127.0.0.1:9879")
	queryAndExpectError(t, "http://127.0.0.1:9879/docstore/user/entities/123", ErrEntityMissing)
}

func queryAndExpectError(t *testing.T, url string, errorNumber int) {
	should := require.New(t)
	resp, err := http.Get(url)
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(errorNumber, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}
