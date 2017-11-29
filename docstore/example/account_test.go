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
	"github.com/v2pro/quokka/kvstore/memkv"
	"context"
)

func Test_account(t *testing.T) {
	memkv.ResetKVStore()
	docstore.Entity("Account", `
	struct Doc {
		1: i64 amount
		2: string account_type
	}
	`).Command("create", `
	function handle(doc, req) {
		doc.amount = req.amount;
		doc.account_type = req.account_type;
		return null;
	}
	`, `
	struct Request {
		1: i64 amount
		2: string account_type
	}
	struct Response {
	}
	`).Command("charge", `
	function handle(doc, req) {
		if (doc.account_type == 'vip') {
			if (doc.amount - req.charge < -10) {
				throw 'vip account can not below -10';
			}
		} else {
			if (doc.amount - req.charge < 0) {
				throw 'normal account can not below 0';
			}
		}
		doc.amount -= req.charge;
		return {remaining_amount: doc.amount};
	}
	`, `
	struct Request {
		1: i64 charge
	}
	struct Response {
		1: i64 remaining_amount
	}
	`)
	docstore.StartNode(context.TODO(), "127.0.0.1:2515")
	time.Sleep(time.Second)
	post(t, "http://127.0.0.1:2515/docstore/Account/create", `
{
	"EntityId": "123",
	"CommandRequest": {
		"amount": 100,
		"account_type": "vip"
	}
}
	`)
	post(t, "http://127.0.0.1:2515/docstore/Account/charge", `
{
	"EntityId": "123",
	"CommandId": "acvx",
	"CommandRequest": {
		"charge": 10
	}
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
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
}
