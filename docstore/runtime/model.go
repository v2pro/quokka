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

var Json = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

var DebugJson = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
	SortMapKeys:                   true,
}.Froze()

var emptyInterfaceType = reflect.TypeOf((*interface{})(nil)).Elem()
var dobjectType = reflect.TypeOf((*DObject)(nil)).Elem()
var darrayType = reflect.TypeOf((*DArray)(nil)).Elem()

func init() {
	Json.RegisterExtension(&jsonExtension{})
	DebugJson.RegisterExtension(&jsonExtension{})
}

type DObject struct {
	Schema  *Schema `json:"-"`
	data    map[string]interface{}
	updated map[string]interface{}
	patched map[string]interface{}
	removed map[string]interface{}
}

func AsBool(val interface{}) bool {
	return val.(bool)
}

func AsObj(val interface{}) Object {
	return val.(Object)
}

func NewObject(kv ...interface{}) *DObject {
	data := map[string]interface{}{}
	for i := 0; i < len(kv); i += 2 {
		data[kv[i].(string)] = kv[i+1]
	}
	return &DObject{
		data: data,
	}
}

type DArray struct {
	data []interface{}
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

type jsonExtension struct {
	jsoniter.DummyExtension
}

func (extension *jsonExtension) CreateDecoder(typ reflect.Type) jsoniter.ValDecoder {
	if typ == emptyInterfaceType {
		return &emptyInterfaceDecoder{}
	}
	return nil
}

func (extension *jsonExtension) CreateEncoder(typ reflect.Type) jsoniter.ValEncoder {
	if typ == dobjectType {
		return &objectEncoder{}
	} else if typ == darrayType {
		return &arrayEncoder{}
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
		subIter := Json.BorrowIterator(number)
		intVal := subIter.ReadInt()
		if subIter.Error == io.EOF {
			Json.ReturnIterator(subIter)
			*(*interface{})(ptr) = intVal
		} else {
			Json.ReturnIterator(subIter)
			subIter = Json.BorrowIterator(number)
			*(*interface{})(ptr) = subIter.ReadFloat64()
			if subIter.Error != nil && subIter.Error != io.EOF {
				iter.ReportError("readNumber", subIter.Error.Error())
			}
			Json.ReturnIterator(subIter)
		}
	default:
		iter.ReportError("readEmptyInterface", fmt.Sprintf("unexpected value type: %v", nextValueType))
	}
}

type objectEncoder struct {
}

func (encoder *objectEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func (encoder *objectEncoder) EncodeInterface(val interface{}, stream *jsoniter.Stream) {
	jsoniter.WriteToStream(val, stream, encoder)
}

func (encoder *objectEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DObject)(ptr)
	stream.WriteVal(obj.data)
}

type arrayEncoder struct {
}

func (encoder *arrayEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func (encoder *arrayEncoder) EncodeInterface(val interface{}, stream *jsoniter.Stream) {
	jsoniter.WriteToStream(val, stream, encoder)
}

func (encoder *arrayEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DArray)(ptr)
	stream.WriteVal(obj.data)
}

func Dump(obj interface{}) string {
	// DebugJson sorts map key, so it is stable to compare
	encoded, err := DebugJson.MarshalToString(obj)
	if err != nil {
		panic(err)
	}
	return encoded
}
