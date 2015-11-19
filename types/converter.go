package types

import (
	"fmt"
	"reflect"
	"strconv"
)

func StrToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 0, 64)
	return i, err
}

func StrToInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 0, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), err
}

func StrToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	return f, err
}

func StrToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return float32(f), err
}

func ByteToInt64(b []byte) (int64,error) {
	s := string(b)
	return StrToInt64(s)
}

func Int64ToByte(i int64) ([]byte,error) {
	var b []byte
	return strconv.AppendInt(b,i,10),nil
}

func ToString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	rv := reflect.ValueOf(src)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(rv.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(rv.Uint(), 10)
	case reflect.Float64:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 64)
	case reflect.Float32:
		return strconv.FormatFloat(rv.Float(), 'g', -1, 32)
	case reflect.Bool:
		return strconv.FormatBool(rv.Bool())
	}
	return fmt.Sprintf("%v", src)
}
