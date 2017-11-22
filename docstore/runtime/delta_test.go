package runtime

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_update_root(t *testing.T) {
	should := require.New(t)
	obj := NewObject()
	obj.Update("hello", "world")
	should.Equal("world", obj.data["hello"])
	should.Equal("world", obj.updated["hello"])
	delta, err := deltaJson.MarshalToString(obj)
	should.Nil(err)
	should.Equal(`{"__updated__":{"hello":"world"}}`, delta)
}

func Test_patch_one_level(t *testing.T) {
	should := require.New(t)
	root := NewObject()
	leaf := NewObject()
	root.data["leaf"] = leaf
	leaf.Update("hello", "world")
	delta, err := deltaJson.MarshalToString(root)
	should.Nil(err)
	should.Equal(`{"__patched__":{"leaf":{"__updated__":{"hello":"world"}}}}`, delta)
}