package runtime

import (
	"github.com/json-iterator/go"
	"reflect"
	"unsafe"
)

var deltaJson = jsoniter.Config{
	EscapeHTML:                    false,
	MarshalFloatWith6Digits:       true, // will lose precession
	ObjectFieldMustBeSimpleString: true, // do not unescape object field
}.Froze()

func init() {
	deltaJson.RegisterExtension(&deltaJsonExtension{})
}

func (obj *DObject) Update(key string, value interface{}) {
	obj.data[key] = value
	if obj.updated == nil {
		obj.updated = map[string]interface{}{}
	}
	obj.updated[key] = value
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
		stream.WriteObjectField("__updated__")
		subStream := json.BorrowStream(nil) // update is not in delta format
		defer json.ReturnStream(subStream)
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
		stream.WriteObjectField("__patched__")
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
		case "__updated__":
			iter.ReadMapCB(func(iter *jsoniter.Iterator, field string) bool {
				input := iter.SkipAndReturnBytes()
				subIter := json.BorrowIterator(input) // switch from deltaJson to json
				defer json.ReturnIterator(subIter)
				obj.data[field] = subIter.Read()
				return true
			})
		case "__patched__":
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
