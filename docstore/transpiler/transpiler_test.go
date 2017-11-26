package transpiler

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/quokka/docstore/runtime"
	"fmt"
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

func Test_report_error(t *testing.T) {
	should := require.New(t)
	defer func() {
		recovered := recover()
		should.Contains(fmt.Sprintf("%v", recovered), "req['hello']")
	}()
	fn, err := compile(`req['1']='2';
	req['1']='2';
	req['1']='2';
	req['1']='2';
	req['hello'];
	req['1']='2';
	return req['hello'];
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	obj.Set("hello", "world")
	fn(obj, nil)
}
