package converter

import (
	"fmt"
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

func ToString(obj interface{}) string {
	res := fmt.Sprintf("%v", obj)
	return string(res)
}



