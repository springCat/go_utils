package converter

import (
	"testing"
)

func TestToInt64(t *testing.T) {
	tests := []string{"1000", "-123", "abcdef", "100000000000000000000000000000000000000000000"}
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
	tests := []string{"1000", "-123", "abcdef", "100000000000000000000000000000000000000000000"}
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

func TestToBool(t *testing.T) {
	tests := []string{"true", "1", "True", "false", "0", "abcdef"}
	expected := []bool{true, true, true, false, false, false}
	for i := 0; i < len(tests); i++ {
		res, _ := ToBool(tests[i])
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
}

func testToJson(t *testing.T, test interface{}, expected string) {
	res, err := ToJson(test)
	if res != expected || err != nil {
		t.Log("Case ToString: expected ", expected, " when result is ", res)
		t.FailNow()
	}
}

func TestToJson(t *testing.T) {

	testToJson(t, map[string]string{"a": "a", "b": "b"}, "{\"a\":\"a\",\"b\":\"b\"}")

	type user struct {
		Name  string  `json:"name"`
		Age   int64   `json:"age"`
		Money float64 `json:"money"`
	}

	u := user{
		Name:  "小明",
		Age:   12,
		Money: 1000.0,
	}

	u1 := user{
		Name:  "小明1",
		Age:   13,
		Money: 1001.0,
	}

	us := make([]user, 0)
	us = append(us, u)
	us = append(us, u1)

	testToJson(t, u, "{\"name\":\"小明\",\"age\":12,\"money\":1000}")
	testToJson(t, us, "[{\"name\":\"小明\",\"age\":12,\"money\":1000},{\"name\":\"小明1\",\"age\":13,\"money\":1001}]")
}

func testToXml(t *testing.T, test interface{}, expected string) {
	res, err := ToXml(test)
	if res != expected || err != nil {
		t.Log("Case ToString: expected ", expected, " when result is ", res)
		t.FailNow()
	}
}

func TestToXml(t *testing.T) {
	//	m := map[string]string{"a":"a","b":"b"}
	//	testToXml(t,m,"1")

	type user struct {
		Name  string  `xml:"name"`
		Age   int64   `xml:"age"`
		Money float64 `xml:"money"`
	}

	u := user{
		Name:  "小明",
		Age:   12,
		Money: 1000.0,
	}

	u1 := user{
		Name:  "小明1",
		Age:   13,
		Money: 1001.0,
	}

	us := make([]user, 0)
	us = append(us, u)
	us = append(us, u1)

	testToXml(t, u, "<user><name>小明</name><age>12</age><money>1000</money></user>")
	testToXml(t, us, "<user><name>小明</name><age>12</age><money>1000</money></user><user><name>小明1</name><age>13</age><money>1001</money></user>")

}
