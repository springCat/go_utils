package reflectx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	String  string
	Int     int
	Uint8   uint8
	Float32 float32
	Bool    bool
	Ints    []int
}

func TestIsZero(t *testing.T) {
	var i int
	assert.Equal(t, true, IsZero(i))
	i1 := 0
	assert.Equal(t, true, IsZero(i1))

	var s string
	assert.Equal(t, true, IsZero(s))
	s1 := 0
	assert.Equal(t, true, IsZero(s1))

	var ts TestStruct
	assert.Equal(t, true, IsZero(ts))
	ts1 := new(TestStruct)
	assert.Equal(t, false, IsZero(ts1))

	var tss []TestStruct
	assert.Equal(t, true, IsZero(tss))
	tss1 := make([]TestStruct, 0)
	assert.Equal(t, false, IsZero(tss1))

	var m map[string]interface{}
	assert.Equal(t, true, IsZero(m))
	m1 := make(map[string]interface{})
	assert.Equal(t, false, IsZero(m1))

	var obj interface{} = make(map[string]interface{})
	assert.Equal(t, false, IsZero(obj))
}

func TestIsEmpty(t *testing.T) {
	var i int
	assert.Equal(t, true, IsEmpty(i))
	i1 := 0
	assert.Equal(t, true, IsEmpty(i1))

	var s string
	assert.Equal(t, true, IsEmpty(s))
	s1 := 0
	assert.Equal(t, true, IsEmpty(s1))

	var ts TestStruct
	assert.Equal(t, true, IsEmpty(ts))
	ts1 := new(TestStruct)
	assert.Equal(t, false, IsEmpty(ts1))

	var tss []TestStruct
	assert.Equal(t, true, IsEmpty(tss))
	tss1 := make([]TestStruct, 0)
	assert.Equal(t, true, IsEmpty(tss1))

	var m map[string]interface{}
	assert.Equal(t, true, IsEmpty(m))
	m1 := make(map[string]interface{})
	assert.Equal(t, true, IsEmpty(m1))

	var obj interface{} = make(map[string]interface{})
	assert.Equal(t, true, IsEmpty(obj))

}
