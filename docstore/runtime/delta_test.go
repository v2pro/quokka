package runtime

import (
	"testing"
	"github.com/stretchr/testify/require"
	"fmt"
)

func clone(obj interface{}) interface{} {
	encoded, err := Json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	var cloned interface{}
	err = Json.Unmarshal(encoded, &cloned)
	if err != nil {
		panic(err)
	}
	return cloned
}

func dump(obj interface{}) string {
	// DebugJson sorts map key, so it is stable to compare
	encoded, err := DebugJson.MarshalToString(obj)
	if err != nil {
		panic(err)
	}
	return encoded
}

func Test_update_root(t *testing.T) {
	should := require.New(t)
	obj := NewObject()
	obj.data["origKey"] = "origValue"
	origObj := clone(obj)
	obj.Set("hello", "world")
	should.Equal("world", obj.data["hello"])
	should.Equal("world", obj.updated["hello"])
	delta, err := DeltaJson.MarshalToString(obj)
	should.Nil(err)
	should.Equal(`{"__updated__":{"hello":"world"}}`, delta)
	should.Nil(DeltaJson.UnmarshalFromString(delta, &origObj))
	should.Equal(dump(obj), dump(origObj))
}

func Test_patch_one_level(t *testing.T) {
	should := require.New(t)
	root := NewObject()
	leaf := NewObject()
	leaf.data["origKey"] = "origValue"
	root.data["leaf"] = leaf
	origRoot := clone(root)
	leaf.Set("hello", "world")
	delta, err := DeltaJson.MarshalToString(root)
	should.Nil(err)
	should.Equal(`{"__patched__":{"leaf":{"__updated__":{"hello":"world"}}}}`, delta)
	fmt.Println(dump(origRoot))
	should.Nil(DeltaJson.UnmarshalFromString(delta, &origRoot))
	fmt.Println(dump(origRoot))
	should.Equal(dump(root), dump(origRoot))
}