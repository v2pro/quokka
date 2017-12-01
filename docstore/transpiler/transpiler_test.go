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
	return null;
	`, runtime.NewObject("hello", "world"), nil},
		{"set object", `
	doc['hello']={1:2.1};
	return null;
	`, runtime.NewObject("hello", runtime.NewObject("1", 2.1)), nil},
		{"get string", `
	doc['hello']='world';
	return doc['hello'];
	`, runtime.NewObject("hello", "world"), "world"},
		{"get set by dot", `
	doc.hello='world';
	return doc.hello;
	`, runtime.NewObject("hello", "world"), "world"},
		{"throw error", `
	throw 'hello';
	`, runtime.NewObject(), "hello"},
		{"arithmetic", `
	return (1+1-1*5) / 3;
	`, runtime.NewObject(), float64(-1)},
		{"-= and +=", `
	doc['val'] = 0;
	doc['val'] += 1;
	doc['val'] -= 1;
	return null;
	`, runtime.NewObject("val", 0), nil},
		{"number comparison", `
	doc['val'] = 1;
	if (doc['val'] > 0) {
		doc['val'] = 0;
	}
	doc['val'] >= 0;
	doc['val'] < 0;
	doc['val'] <= 0;
	return null;
	`, runtime.NewObject("val", 0), nil},
		{"string comparison", `
	doc['val'] = "hello";
	if (doc['val'] == "hello") {
		doc['val'] = "world";
	}
	return null;
	`, runtime.NewObject("val", "world"), nil},
		{"list length", `
	doc['val'] = ["hello"];
	return {"length": doc.val.length};
	`, runtime.NewObject("val", runtime.NewList("hello")),
	runtime.NewObject("length", 1)},
	}
	for _, fixture := range fixtures {
		t.Run(fixture.title, func(t *testing.T) {
			should := require.New(t)
			fn, err := Compile("function handle(doc, req) {\n" + fixture.src + "\n}")
			should.Nil(err)
			doc := runtime.NewObject()
			resp := fn(doc, nil)
			should.Equal(runtime.Dump(fixture.expectedDoc), runtime.Dump(doc))
			if err, _ := resp.(error); err != nil {
				should.Nil(err, err.Error())
			}
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
