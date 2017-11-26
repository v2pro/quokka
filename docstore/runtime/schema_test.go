package runtime

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_thrift(t *testing.T) {
	should := require.New(t)
	obj := NewObject()
	obj.Schema = NewSchemaByThriftIDL(`
	struct Doc {
		1: required string hello;
		2: required string world;
	}
	`)
	should.Panics(func() {
		obj.Set("hello", 1)
	})
}

func Test_validate_direct_field(t *testing.T) {
	should := require.New(t)
	obj := NewObject()
	obj.Schema = &Schema{
		Fields: map[string]*Schema{
			"hello": mustBeString,
		},
	}
	should.Panics(func() {
		obj.Set("hello", 1)
	})
}

func Test_validate_indirect_field(t *testing.T) {
	should := require.New(t)
	obj := NewObject()
	obj.Schema = &Schema{
		Fields: map[string]*Schema{
			"hello": {
				Fields: map[string]*Schema{
					"world": mustBeString,
				},
			},
		},
	}
	should.Panics(func() {
		obj.Set("hello", NewObject("world", 1))
	})
}