package reflect

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	String  string
	Int     int
	Uint8   uint8
	Float32 float32
	Bool    bool
	Ints   []int
}

func TestSetStructField(t *testing.T) {
	structPtr := new(TestStruct)

	SetStructField(structPtr,"String","111")
	fmt.Println(structPtr)
	i1,_ := GetStructField(structPtr,"String")
	fmt.Println("i1:", i1)

	SetStructField(structPtr,"Int",111)
	i2,_ :=GetStructField(structPtr,"Int")
	fmt.Println("i2:",i2 )

	ints := []int{1,2,3,4,5}
	SetStructField(structPtr,"Ints",ints)
	i3,_ :=GetStructField(structPtr,"Ints")
	fmt.Println("i3:",i3 )
}

func TestIsZero(t *testing.T) {
	var i int
	assert.Equal(t,true,IsZero(i));
	i1 := 0
	assert.Equal(t,true,IsZero(i1));
	var s string
	assert.Equal(t,true,IsZero(s));
	s1 := 0
	assert.Equal(t,true,IsZero(s1));
	var ts TestStruct
	assert.Equal(t,true,IsZero(ts));
	ts1 := new(TestStruct)
	assert.Equal(t,false,IsZero(ts1))
	var tss []TestStruct
	assert.Equal(t,true,IsZero(tss))
	tss1 := make([]TestStruct,0)
	assert.Equal(t,false,IsZero(tss1))

	var m map[string]interface{}
	assert.Equal(t,true,IsZero(m))
	m1 := make(map[string]interface{})
	assert.Equal(t,false,IsZero(m1))

	var obj interface{} = make(map[string]interface{})
	assert.Equal(t,false,IsZero(obj))
}

func TestIsEmpty(t *testing.T) {
	var i int
	assert.Equal(t,true,IsEmpty(i));
	i1 := 0
	assert.Equal(t,true,IsEmpty(i1));
	var s string
	assert.Equal(t,true,IsEmpty(s));
	s1 := 0
	assert.Equal(t,true,IsEmpty(s1));
	var ts TestStruct
	assert.Equal(t,true,IsEmpty(ts));
	ts1 := new(TestStruct)
	assert.Equal(t,false,IsEmpty(ts1))
	var tss []TestStruct
	assert.Equal(t,true,IsEmpty(tss))
	tss1 := make([]TestStruct,0)
	assert.Equal(t,true,IsEmpty(tss1))

	var m map[string]interface{}
	assert.Equal(t,true,IsEmpty(m))
	m1 := make(map[string]interface{})
	assert.Equal(t,true,IsEmpty(m1))

	var obj interface{} = make(map[string]interface{})
	assert.Equal(t,true,IsEmpty(obj))


}
type A struct {
	X int
}
type B struct {
	Z A
	Y int
}

func TestStructToMap(t *testing.T) {
	b := B{
		Z: A{
			X: 11,
		},
		Y: 21,
	}

	m := StructToMap(b)
	test(t, m["Y"], 21)

	var m1 map[string]interface{} = m["Z"].(map[string]interface{})
	test(t, m1["X"], 11)

}

func test(t *testing.T, test interface{}, expected interface{}) {
	res := (test)
	if res != expected {
		t.Log("Case ToString: expected ", expected, " when result is ", res)
		t.FailNow()
	}
}

func TestStrMapToStruct(t *testing.T) {
	structPtr := new(TestStruct)
	m := map[string]string{
		"String":  "Hello World",
		"Int":     "666",
		"Uint8":   "234",
		"Float32": "0.01",
		"Bool":    "true",
	}
	err := StrMapToStruct(structPtr, m, true)
	if err != nil {
		t.Fatal(err)
	}
	check(t,structPtr)

	m["NotExisting"] = "xxx"

	structPtr = new(TestStruct)
	err = StrMapToStruct(structPtr, m, true)
	if err == nil {
		t.Fail()
	}

	structPtr = new(TestStruct)
	err = StrMapToStruct(structPtr, m, false)
	if err != nil {
		t.Fatal(err)
	}
	check(t,structPtr)
}

func TestMapToStruct(t *testing.T) {
	structPtr := new(TestStruct)
	m := map[string]interface{}{
		"String":  "Hello World",
		"Int":     666,
		"Uint8":   uint8(234),
		"Float32": float32(0.01),
		"Bool":    true,
		"Ints":    []int{1, 2},
	}
	fmt.Println("m:", m)
	err := MapToStruct(structPtr, m, true)
	if err != nil {
		t.Fatal(err)
	}
	check(t, structPtr)

	m["NotExisting"] = "xxx"

	structPtr = new(TestStruct)
	err = MapToStruct(structPtr, m, true)
	if err == nil {
		t.Fail()
	}

	structPtr = new(TestStruct)
	err = MapToStruct(structPtr, m, false)
	if err != nil {
		t.Fatal(err)
	}
	check(t, structPtr)
	ints := structPtr.Ints
	if  ints[0] != 1 ||
	ints[1] != 2 {
		t.Fatalf("Invalid values: %#v", structPtr)
	}
}

func check(t *testing.T, structPtr *TestStruct) {
	if structPtr.String != "Hello World" ||
	structPtr.Int != 666 ||
	structPtr.Uint8 != 234 ||
	structPtr.Float32 != 0.01 ||
	structPtr.Bool != true {
		t.Fatalf("Invalid values: %#v", structPtr)
	}
}
