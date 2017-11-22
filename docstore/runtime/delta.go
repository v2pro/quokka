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

var dobjectType = reflect.TypeOf((*DObject)(nil)).Elem()

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

type deltaJsonExtension struct {
	jsoniter.DummyExtension
}

func (extension *deltaJsonExtension) CreateEncoder(typ reflect.Type) jsoniter.ValEncoder {
	if dobjectType == typ {
		return &objectEncoder{}
	}
	return nil
}

type objectEncoder struct {
}

func (encoder *objectEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	return false
}

func (encoder *objectEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
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

func (encoder *objectEncoder) EncodeInterface(val interface{}, stream *jsoniter.Stream) {
	jsoniter.WriteToStream(val, stream, encoder)
}
