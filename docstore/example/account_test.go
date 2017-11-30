package example

import (
	"testing"
	"github.com/v2pro/quokka/docstore"
	"net/http"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"github.com/json-iterator/go"
	"github.com/v2pro/quokka/kvstore/memkv"
	"context"
	"bytes"
	"github.com/v2pro/quokka/docstore/runtime"
)

func init() {
	docstore.Entity("Account", `
	struct Doc {
		1: i64 amount
		2: string account_type
	}
	`).Command("create", `
	function handle(doc, req) {
		doc.amount = req.amount;
		doc.account_type = req.account_type;
		return {};
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
			log_trace("event!charge", "amount", doc.amount, "charge", req.charge);
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
}

func TestMain(m *testing.M) {
	memkv.ResetKVStore()
	docstore.StartNode(context.TODO(), "127.0.0.1:2515")
	m.Run()
	docstore.StopNode(context.TODO())
}

func Test_account(t *testing.T) {
	execAndExpectSuccess(t, "http://127.0.0.1:2515/docstore/Account/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 100, "account_type", "vip"))
	execAndExpectSuccess(t, "http://127.0.0.1:2515/docstore/Account/charge",
		"EntityId", "123", "CommandRequest", runtime.NewObject("charge", 10))
	execAndExpectError(t, "http://127.0.0.1:2515/docstore/Account/charge", docstore.ErrBusinessRuleViolated,
		"EntityId", "123", "CommandRequest", runtime.NewObject("charge", 101))
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
