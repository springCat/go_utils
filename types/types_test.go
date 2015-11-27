package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
"fmt"
)

func TestBasicType(t *testing.T) {
	i := 1
	fmt.Println("i:", Help(i).IntOr(19))

	f := 1.01;
	fmt.Println("i:", Help(f).Float64Or(2.0))
}

func TestSlice(t *testing.T) {

		v := make([]int, 0)
		v = append(v, 1)
		v = append(v, 2)
		v = append(v, 3)
		v = append(v, 4)
		v = append(v, 5)

		i := Help(v).IntSlice()
		assert.Equal(t,[]int{1,2,3,4,5},i)

		v1 := make([]float32, 0.0)
		v1 = append(v1, 0.1)
		v1 = append(v1, 0.2)
		v1 = append(v1, 0.3)
		v1 = append(v1, 0.4)
		v1 = append(v1, 0.5)

	f := Help(v1).Float32Slice()
	fmt.Println("f:", f)
	assert.Equal(t,[]float32{0.1,0.2,0.3,0.4,0.5},f)
}

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
	assert.Equal(t, true, Help(i).IsZero())
	i1 := 0
	assert.Equal(t, true, Help(i1).IsZero())

	var s string
	assert.Equal(t, true, Help(s).IsZero())
	s1 := 0
	assert.Equal(t, true, Help(s1).IsZero())

	var ts TestStruct
	assert.Equal(t, true, Help(ts).IsZero())
	ts1 := new(TestStruct)
	assert.Equal(t, false, Help(ts1).IsZero())

	var tss []TestStruct
	assert.Equal(t, true, Help(tss).IsZero())
	tss1 := make([]TestStruct, 0)
	assert.Equal(t, false, Help(tss1).IsZero())

	var m map[string]interface{}
	assert.Equal(t, true, Help(m).IsZero())
	m1 := make(map[string]interface{})
	assert.Equal(t, false, Help(m1).IsZero())

	var obj interface{} = make(map[string]interface{})
	assert.Equal(t, false, Help(obj).IsZero())
}

func TestIsEmpty(t *testing.T) {
	var i int
	assert.Equal(t, true, Help(i).IsEmpty())
	i1 := 0
	assert.Equal(t, true, Help(i1).IsEmpty())

	var s string
	assert.Equal(t, true, Help(s).IsEmpty())
	s1 := 0
	assert.Equal(t, true, Help(s1).IsEmpty())

	var ts TestStruct
	assert.Equal(t, true, Help(ts).IsEmpty())
	ts1 := new(TestStruct)
	assert.Equal(t, false, Help(ts1).IsEmpty())

	var tss []TestStruct
	assert.Equal(t, true, Help(tss).IsEmpty())
	tss1 := make([]TestStruct, 0)
	assert.Equal(t, true, Help(tss1).IsEmpty())

	var m map[string]interface{}
	assert.Equal(t, true, Help(m).IsEmpty())
	m1 := make(map[string]interface{})
	assert.Equal(t, true, Help(m1).IsEmpty())

	var obj interface{} = make(map[string]interface{})
	assert.Equal(t, true, Help(obj).IsEmpty())
}
