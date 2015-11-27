package types

import (
	"testing"
	"time"
"fmt"
	"encoding/json"
)

func TestToInt64(t *testing.T) {
	tests := []string{"1000", "-123", "abcdef", "0"}
	expected := []int64{1000, -123, 0, 0}
	for i := 0; i < len(tests); i++ {
		result, _ := ToInt64(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestToInt32(t *testing.T) {
	tests := []string{"1000", "-123", "abcdef", "0"}
	expected := []int32{1000, -123, 0, 0}
	for i := 0; i < len(tests); i++ {
		result, _ := ToInt32(tests[i])
		if result != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", result)
			t.FailNow()
		}
	}
}

func TestToFloat64(t *testing.T) {
	tests := []string{"", "123", "-.01", "10.", "string", "1.23e3", ".23e10"}
	expected := []float64{0, 123, -0.01, 10.0, 0, 1230, 0.23e10}
	for i := 0; i < len(tests); i++ {
		res, _ := ToFloat64(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func TestToFloat32(t *testing.T) {
	tests := []string{"", "123", "-.01", "10.", "string", "1.23e3", ".23e10"}
	expected := []float32{0, 123, -0.01, 10.0, 0, 1230, 0.23e10}
	for i := 0; i < len(tests); i++ {
		res, _ := ToFloat32(tests[i])
		if res != expected[i] {
			t.Log("Case ", i, ": expected ", expected[i], " when result is ", res)
			t.FailNow()
		}
	}
}

func TestToString(t *testing.T) {
	m := map[interface{}]interface{}{
		"str123":  "str123",
		123:       "123",
		12.3:      "12.3",
		true:      "true",
		1.5 + 10i: "(1.5+10i)",
	}

	for k, v := range m {
		res := ToString(k)
		if res != v {
			t.Log("Case ", k, ": expected ", v, " when result is ", res)
			t.FailNow()
		}
	}

	type Hobby struct {
		Name string
		Desc string
		Start time.Time
	}

	type man struct {
		Id   int `json:"-"`
		Name string
		Sex  string
		Hobbys []Hobby
 	}



	user := &man{1,"小明","男",[]Hobby{{"画画","huahua",time.Now()},{"doti","dotinow",time.Now()}}}

	fmt.Println("ToString(user):", ToString(user))


	b,_ := json.Marshal(user)
	fmt.Println("ToString(user):", ToString(b))

}
