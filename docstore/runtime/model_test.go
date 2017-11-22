package runtime

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func Test_object(t *testing.T) {
	should := require.New(t)
	var obj interface{}
	encoded := `{"hello":"world"}`
	should.Nil(json.UnmarshalFromString(encoded, &obj))
	should.Equal("world", obj.(*DObject).data["hello"])
	reEncoded, err := json.MarshalToString(obj)
	should.Nil(err)
	should.Equal(encoded, reEncoded)
}

func Test_array(t *testing.T) {
	should := require.New(t)
	var obj interface{}
	encoded := `[1,2,3]`
	should.Nil(json.UnmarshalFromString(encoded, &obj))
	should.Equal(int64(1), obj.(*DArray).data[0])
	should.Equal(int64(2), obj.(*DArray).data[1])
	should.Equal(int64(3), obj.(*DArray).data[2])
	reEncoded, err := json.MarshalToString(obj)
	should.Nil(err)
	should.Equal(encoded, reEncoded)
}

func Test_float(t *testing.T) {
	should := require.New(t)
	var obj interface{}
	should.Nil(json.Unmarshal([]byte(`1e+1`), &obj))
	should.Equal(float64(10), obj.(float64))
}

func Test_nested(t *testing.T) {
	should := require.New(t)
	var obj interface{}
	should.Nil(json.Unmarshal([]byte(`{"hello":[true, null]}`), &obj))
	should.Equal(true, Get(obj, "hello", 0))
	should.Equal(nil, Get(obj, "hello", 1))
}