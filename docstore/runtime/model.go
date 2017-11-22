package runtime

import (
	"github.com/json-iterator/go"
	"reflect"
	"unsafe"
	"io"
	"fmt"
)

// unmarshal will only return following types
// float64
// string
// bool
// *DArray
// *DObject

var json = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

var emptyInterfaceType = reflect.TypeOf((*interface{})(nil)).Elem()

func init() {
	json.RegisterExtension(&jsonExtension{})
}

type DObject struct {
	data map[string]interface{}
	updated map[string]interface{}
	patched map[string]interface{}
	removed map[string]interface{}
}

func NewObject() *DObject {
	return &DObject{
		data: map[string]interface{}{},
	}
}

type DArray struct {
	data []interface{}
}

type jsonExtension struct {
	jsoniter.DummyExtension
}

func (extension *jsonExtension) CreateDecoder(typ reflect.Type) jsoniter.ValDecoder {
	if typ == emptyInterfaceType {
		return &emptyInterfaceDecoder{}
	}
	return nil
}

type emptyInterfaceDecoder struct {
}

func (decoder *emptyInterfaceDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	nextValueType := iter.WhatIsNext()
	switch nextValueType {
	case jsoniter.ObjectValue:
		obj := &DObject{}
		iter.ReadVal(&obj.data)
		*(*interface{})(ptr) = obj
	case jsoniter.ArrayValue:
		arr := &DArray{}
		iter.ReadVal(&arr.data)
		*(*interface{})(ptr) = arr
	case jsoniter.StringValue:
		*(*interface{})(ptr) = iter.ReadString()
	case jsoniter.BoolValue:
		*(*interface{})(ptr) = iter.ReadBool()
	case jsoniter.NilValue:
		iter.ReadNil()
		*(*interface{})(ptr) = nil
	case jsoniter.NumberValue:
		number := iter.SkipAndReturnBytes()
		subIter := json.BorrowIterator(number)
		intVal := subIter.ReadInt64()
		if subIter.Error == io.EOF {
			json.ReturnIterator(subIter)
			*(*interface{})(ptr) = intVal
		} else {
			json.ReturnIterator(subIter)
			subIter = json.BorrowIterator(number)
			*(*interface{})(ptr) = subIter.ReadFloat64()
			if subIter.Error != nil && subIter.Error != io.EOF {
				iter.ReportError("readNumber", subIter.Error.Error())
			}
			json.ReturnIterator(subIter)
		}
	default:
		iter.ReportError("readEmptyInterface", fmt.Sprintf("unexpected value type: %v", nextValueType))
	}
}

func Get(obj interface{}, path ...interface{}) interface{} {
	for _, elem := range path {
		switch typedElem := elem.(type) {
		case int:
			obj = obj.(*DArray).data[typedElem]
		case string:
			obj = obj.(*DObject).data[typedElem]
		}
	}
	return obj
}
