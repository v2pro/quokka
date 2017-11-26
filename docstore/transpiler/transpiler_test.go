package transpiler

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/quokka/docstore/runtime"
)

func Test_set_field(t *testing.T) {
	should := require.New(t)
	fn, err := compile(`
	doc['hello']='world';
	return nil;
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	fn(obj, nil)
	should.Equal("world", runtime.Get(obj, "hello"))
}

func Test_get_field(t *testing.T) {
	should := require.New(t)
	fn, err := compile(`
	return doc['hello'];
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	obj.Set("hello", "world")
	should.Equal("world", fn(obj, nil))
}
