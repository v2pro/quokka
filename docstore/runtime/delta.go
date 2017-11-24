package runtime

import (
	"github.com/json-iterator/go"
	"reflect"
	"unsafe"
)

var DeltaJson = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

func init() {
	DeltaJson.RegisterExtension(&deltaJsonExtension{})
}

type Object interface {
	Set(key interface{}, value interface{})
}

func AsObj(debugInfo string, val interface{}) Object {
	obj, success := val.(Object)
	if success {
		return obj
	}
	panic(debugInfo + ": can not cast " + reflect.TypeOf(val).String())
}

func (obj *DObject) Set(key interface{}, value interface{}) {
	obj.data[key.(string)] = value
	if obj.updated == nil {
		obj.updated = map[string]interface{}{}
	}
	obj.updated[key.(string)] = value
}

func (obj *DObject) isDirty() bool {
	if obj.updated != nil {
		return true
	}
	if obj.removed != nil {
		return true
	}
	if len(obj.calcPatched()) > 0 {
		return true
	}
	return false
}

func (obj *DObject) calcPatched() map[string]interface{} {
	if obj.patched != nil {
		return obj.patched
	}
	obj.patched = map[string]interface{}{}
	for k, v := range obj.data {
		switch typedValue := v.(type) {
		case *DObject:
			if typedValue.isDirty() {
				obj.patched[k] = typedValue
			}
		}
	}
	return obj.patched
}

type deltaJsonExtension struct {
	jsoniter.DummyExtension
}

func (extension *deltaJsonExtension) CreateEncoder(typ reflect.Type) jsoniter.ValEncoder {
	if dobjectType == typ {
		return &objectDeltaEncoder{}
	}
	return nil
}

func (extension *deltaJsonExtension) CreateDecoder(typ reflect.Type) jsoniter.ValDecoder {
	if dobjectType == typ {
		return &objectDeltaDecoder{}
	}
	return nil
}

type objectDeltaEncoder struct {
}

func (encoder *objectDeltaEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func (encoder *objectDeltaEncoder) EncodeInterface(val interface{}, stream *jsoniter.Stream) {
	jsoniter.WriteToStream(val, stream, encoder)
}

func (encoder *objectDeltaEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*DObject)(ptr)
	isFirstField := true
	stream.WriteObjectStart()
	if obj.updated != nil {
		stream.WriteObjectField("u")
		subStream := Json.BorrowStream(nil) // update is not in delta format
		defer Json.ReturnStream(subStream)
		subStream.WriteVal(obj.updated)
		stream.Write(subStream.Buffer())
		isFirstField = false
	}
	patched := obj.calcPatched()
	if len(patched) > 0 {
		if !isFirstField {
			stream.WriteMore()
		}
		isFirstField = false
		stream.WriteObjectField("p")
		stream.WriteVal(patched)
	}
	stream.WriteObjectEnd()
}

type objectDeltaDecoder struct {
}

func (decoder *objectDeltaDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	obj := (*DObject)(ptr)
	iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
		switch field {
		case "u":
			iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
				input := iter.SkipAndReturnBytes()
				subIter := Json.BorrowIterator(input) // switch from DeltaJson to Json
				defer Json.ReturnIterator(subIter)
				obj.data[field] = subIter.Read()
				return true
			})
		case "p":
			iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
				fieldValue := obj.data[field]
				iter.ReadVal(&fieldValue)
				obj.data[field] = fieldValue
				return true
			})
		default:
			iter.Skip()
		}
		return true
	})
}
