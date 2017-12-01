package example

import (
	"github.com/v2pro/quokka/docstore"
	"testing"
	"github.com/v2pro/quokka/docstore/runtime"
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
		return {};
	}`, `
	struct Request {
		1: float64 amount
		2: i64 count
	}
	struct Response {
	}
	`)
}

func Test_take_hongbao(t *testing.T) {
	execAndExpectSuccess(t, "http://127.0.0.1:9865/docstore/Hongbao/create",
		"EntityId", "123", "CommandRequest", runtime.NewObject("amount", 131.4, "count", 3))
}
