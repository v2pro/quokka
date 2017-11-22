package transpiler

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/quokka/docstore/runtime"
)

func Test_set_field(t *testing.T) {
	should := require.New(t)
	fn, err := compile(`doc['hello']='world'`)
	should.Nil(err)
	obj := runtime.NewObject()
	fn(obj)
	should.Equal("world", runtime.Get(obj, "hello"))
}
