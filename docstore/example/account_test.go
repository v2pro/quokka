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

func TestMain(m *testing.M) {
	memkv.ResetKVStore()
	docstore.StartNode(context.TODO(), "127.0.0.1:9865")
	m.Run()
	docstore.StopNode(context.TODO())
}

func defineAccount() {
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
	}`, `
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
	}`, `
	struct Request {
		1: i64 charge
	}
	struct Response {
		1: i64 remaining_amount
	}
	`)
}

func Test_account_minimum_amount_rule(t *testing.T) {
	defineAccount()
	execAndExpectSuccess(t, "/Account/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 100, "account_type", "vip"))
	execAndExpectSuccess(t, "/Account/charge",
		"EntityId", "123", "CommandRequest", runtime.NewObject("charge", 10))
	execAndExpectError(t, "/Account/charge", docstore.ErrBusinessRuleViolated,
		"EntityId", "123", "CommandRequest", runtime.NewObject("charge", 101))
}

func Test_charge_idempotence(t *testing.T) {
	defineAccount()
	should := require.New(t)
	execAndExpectSuccess(t, "/Account/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 100, "account_type", "vip"))
	resp := execAndExpectSuccess(t, "/Account/charge",
		"EntityId", "123", "CommandId", "xcvf", "CommandRequest", runtime.NewObject("charge", 10))
	should.Equal(90, jsoniter.Get(resp, "data", "remaining_amount").ToInt())
	resp = execAndExpectSuccess(t, "/Account/charge",
		"EntityId", "123", "CommandId", "xcvf", "CommandRequest", runtime.NewObject("charge", 10))
	should.Equal(90, jsoniter.Get(resp, "data", "remaining_amount").ToInt())
}

func execAndExpectSuccess(t *testing.T, url string, kv ...interface{}) []byte {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post("http://127.0.0.1:9865/docstore" + url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal("", jsoniter.Get(body, "errmsg").ToString())
	should.Equal(0, jsoniter.Get(body, "errno").MustBeValid().ToInt())
	return body
}

func execAndExpectError(t *testing.T, url string, errorNumber int, kv ...interface{}) []byte {
	should := require.New(t)
	req, err := runtime.Json.Marshal(runtime.NewObject(kv...))
	should.Nil(err)
	resp, err := http.Post("http://127.0.0.1:9865/docstore" + url, "application/json", bytes.NewBuffer(req))
	should.Nil(err)
	body, err := ioutil.ReadAll(resp.Body)
	should.Nil(err)
	should.Equal(errorNumber, jsoniter.Get(body, "errno").MustBeValid().ToInt())
	return body
}
