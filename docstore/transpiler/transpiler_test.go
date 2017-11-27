package transpiler

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/v2pro/quokka/docstore/runtime"
)

func Test_compile(t *testing.T) {
	fixtures := []struct {
		title        string
		src          string
		expectedDoc  interface{}
		expectedResp interface{}
	}{
		{"set string", `
	doc['hello']='world';
	return nil;
	`, runtime.NewObject("hello", "world"), nil},
		{"set object",`
	doc['hello']={1:2.1};
	return nil;
	`, runtime.NewObject("hello", runtime.NewObject("1", 2.1)), nil},
		{"get string",`
	doc['hello']='world';
	return doc['hello'];
	`, runtime.NewObject("hello", "world"), "world"},
		{"get set by dot",`
	doc.hello='world';
	return doc.hello;
	`, runtime.NewObject("hello", "world"), "world"},
		{"throw error",`
	throw 'hello';
	`, runtime.NewObject(), "hello"},
	}
	for _, fixture := range fixtures {
		t.Run(fixture.title, func(t *testing.T) {
			should := require.New(t)
			fn, err := Compile("function handle(doc, req) {\n" + fixture.src + "\n}")
			should.Nil(err)
			doc := runtime.NewObject()
			resp := fn(doc, nil)
			should.Equal(runtime.Dump(fixture.expectedDoc), runtime.Dump(doc))
			should.Equal(runtime.Dump(fixture.expectedResp), runtime.Dump(resp))
		})
	}
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
	fn, err := Compile(`
	function troubledFunc(doc, req) {
		doc['1']='1';
		req['hello'];
		return 0;
	}
	function handle(doc, req) {
		doc['1']='1';
		doc['1']='2';
		doc['1']='3';
		troubledFunc(doc, req);
		doc['1']='4';
		return null;
	}
	`)
	should.Nil(err)
	obj := runtime.NewObject()
	obj.Set("hello", "world")
	should.Panics(func() {
		fn(obj, nil)
	})
}
