package runtime

import (
	"testing"
	"github.com/stretchr/testify/require"
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
	obj.Update("hello", "world")
	should.Equal("world", obj.data["hello"])
	should.Equal("world", obj.updated["hello"])
	delta, err := deltaJson.MarshalToString(obj)
	should.Nil(err)
	should.Equal(`{"__updated__":{"hello":"world"}}`, delta)
	should.Nil(deltaJson.UnmarshalFromString(delta, &origObj))
	should.Equal(dump(obj), dump(origObj))
}

func Test_patch_one_level(t *testing.T) {
	should := require.New(t)
	root := NewObject()
	leaf := NewObject()
	leaf.data["origKey"] = "origValue"
	root.data["leaf"] = leaf
	origRoot := clone(root)
	leaf.Update("hello", "world")
	delta, err := deltaJson.MarshalToString(root)
	should.Nil(err)
	should.Equal(`{"__patched__":{"leaf":{"__updated__":{"hello":"world"}}}}`, delta)
	should.Nil(deltaJson.UnmarshalFromString(delta, &origRoot))
	should.Equal(dump(root), dump(origRoot))
}