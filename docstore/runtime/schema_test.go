package runtime

import (
	"testing"
	"github.com/v2pro/quokka/docstore/pb"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"reflect"
)

func Test_pb(t *testing.T) {
	charStream := antlr.NewInputStream(`
	message Doc {
		string nested_msg = 1;
	}
	`)
	lexer := pb.NewProtobuf3Lexer(charStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := pb.NewProtobuf3Parser(tokenStream)
	msg := parser.Message()
	fmt.Println(msg.GetChildCount())
	fmt.Println(msg.GetChild(0))
	fmt.Println(msg.GetChild(1).GetChild(0))
	fmt.Println(reflect.TypeOf(msg.GetChild(2)))
}