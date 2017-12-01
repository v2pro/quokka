package example

import (
	"github.com/v2pro/quokka/docstore"
	"testing"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func init() {
	docstore.Entity("Hongbao", `
	struct taken {
		1: string taken_by
		2: float64 taken_amount
	}
	struct Doc {
		1: i64 total_count
		2: float64 total_amount
		3: float64 remaining_amount
		4: list<taken> takens
	}
	`).Command("create", `
	function handle(doc, req) {
		doc.total_count = req.count;
		doc.total_amount = req.amount;
		doc.remaining_amount = req.amount;
		doc.takens = [];
		return {};
	}`, `
	struct Request {
		1: float64 amount
		2: i64 count
	}
	struct Response {
	}
	`).Command("take", `
	function handle(doc, req) {
		if (doc.takens.length == doc.total_count) {
			throw 'nothing left'
		}
		doc.takens.push({
			taken_by: req.username,
			taken_amount: 1,
		});
		return {taken_amount: 1};
	}`, `
	struct Request {
		1: string username
	}
	struct Response {
		1: float64 taken_amount
	}
	`)
}

func Test_take_hongbao(t *testing.T) {
	should := require.New(t)
	execAndExpectSuccess(t, "/Hongbao/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 131.4, "count", 3))
	resp := execAndExpectSuccess(t, "/Hongbao/take",
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "tom"))
	taken1 := jsoniter.Get(resp, "data", "taken_amount").ToFloat64()
	resp = execAndExpectError(t, "/Hongbao/take", docstore.ErrBusinessRuleViolated,
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "tom"))
	should.Equal("one user can not take twice", jsoniter.Get(resp, "errmsg").ToString())
	resp = execAndExpectSuccess(t, "/Hongbao/take",
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "jerry"))
	taken2 := jsoniter.Get(resp, "data", "taken_amount").ToFloat64()
	resp = execAndExpectSuccess(t, "/Hongbao/take",
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "donald"))
	taken3 := jsoniter.Get(resp, "data", "taken_amount").ToFloat64()
	should.Equal(131.4, taken1 + taken2 + taken3)
	resp = execAndExpectError(t, "/Hongbao/take", docstore.ErrBusinessRuleViolated,
		"EntityId", "123", "CommandRequest", runtime.NewObject("username", "lily"))
	should.Equal("nothing left", jsoniter.Get(resp, "errmsg").ToString())
}
