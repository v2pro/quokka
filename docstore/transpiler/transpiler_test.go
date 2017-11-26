package transpiler

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/quokka/docstore/runtime"
)

func Test_set_string(t *testing.T) {
	should := require.New(t)
	fn, err := Compile(`
	function handle(doc, req) {
		doc['hello']='world';
		return nil;
	}
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	fn(obj, nil)
	should.Equal("world", runtime.Get(obj, "hello"))
}

func Test_set_object(t *testing.T) {
	should := require.New(t)
	fn, err := Compile(`
	doc['hello']={1:2.1};
	return nil;
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	fn(obj, nil)
	should.Equal(2.1, runtime.Get(obj, "hello", "1"))
}

func Test_get_field(t *testing.T) {
	should := require.New(t)
	fn, err := Compile(`
	return doc['hello'];
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	obj.Set("hello", "world")
	should.Equal("world", fn(obj, nil))
}

func Test_sub_function(t *testing.T) {
	should := require.New(t)
	fn, err := Compile(`
	function handle(doc, req) {
		updateDoc(doc);
		return nil;
	}
	function updateDoc(doc) {
		doc['hello']='world'
	}
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	fn(obj, nil)
	should.Equal("world", runtime.Get(obj, "hello"))
}

func Test_report_error(t *testing.T) {
	should := require.New(t)
	fn, err := Compile(`req['1']='2';
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

	should.Panics(func() {
		fn(obj, nil)
	})
}
