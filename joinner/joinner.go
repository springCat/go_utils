package joinner

import (
	"reflect"
	"strconv"
	"strings"
)

type Join interface {
	ToJoinStr() string
}

func StructJoin(joins []Join, sep, prefix, suffix string) string {
	s := initJoinner(joins)
	for _, v := range joins {
		s = append(s, v.ToJoinStr())
	}
	return join(s, sep, prefix, suffix)
}

func IntJoin(ints []int, sep, prefix, suffix string) string {
	s := initJoinner(ints)
	for _, v := range ints {
		s = append(s, strconv.Itoa(v))
	}
	return join(s, sep, prefix, suffix)
}

func Int32Join(ints []int32, sep, prefix, suffix string) string {
	s := initJoinner(ints)
	for _, v := range ints {
		s = append(s, strconv.FormatInt(int64(v), 10))
	}
	return join(s, sep, prefix, suffix)
}

func Int64Join(int64s []int64, sep, prefix, suffix string) string {
	s := initJoinner(int64s)
	for _, v := range int64s {
		s = append(s, strconv.FormatInt(v, 10))
	}
	return join(s, sep, prefix, suffix)
}

func Float32Join(float32s []float32, sep, prefix, suffix string) string {
	s := initJoinner(float32s)
	for _, v := range float32s {
		s = append(s, strconv.FormatFloat(float64(v), 'g', -1, 32))
	}
	return join(s, sep, prefix, suffix)
}

func Float64Join(float64s []float64, sep, prefix, suffix string) string {
	s := initJoinner(float64s)
	for _, v := range float64s {
		s = append(s, strconv.FormatFloat(v, 'g', -1, 64))
	}
	return join(s, sep, prefix, suffix)
}

func initJoinner(array interface{}) []string {
	v := reflect.ValueOf(array)
	l := v.Len()
	s := make([]string, 0, l)
	return s
}

func join(s []string, sep, prefix, suffix string) string {
	innerStr := strings.Join(s, sep)
	n := len(innerStr) + len(prefix) + len(suffix)
	b := make([]byte, n)
	bp := copy(b, prefix)
	bp += copy(b[bp:], innerStr)
	copy(b[bp:], suffix)
	return string(b)
}
