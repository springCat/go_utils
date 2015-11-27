package types

import (
	"errors"
	"reflect"
"fmt"
)

type Type struct {
	reflect.Value
}

var TypeError = errors.New("type error")

func Help(v interface{}) *Type{
	return &Type{
		Value:reflect.ValueOf(v),
	}
}

func (t *Type) IsZero() bool {
	switch t.Value.Kind() {
	case reflect.String:
		return t.Value.Len() == 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return t.Value.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return t.Value.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return t.Value.Float() == 0

	case reflect.Complex64,reflect.Complex128:
		return t.Value.Complex() == 0+0i

	case reflect.Bool:
		return t.Value.Bool() == false

	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface,reflect.Slice, reflect.Map:
		return t.Value.IsNil()

	case reflect.Struct:
		return reflect.DeepEqual(t.Value.Interface(), reflect.Zero(t.Value.Type()).Interface())
	}
	panic(fmt.Errorf("Unknown value kind %T", t.Value))
}

func (t *Type) IsEmpty() bool {

	switch t.Value.Kind() {
	case reflect.String:
		return t.Value.Len() == 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return t.Value.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return t.Value.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return t.Value.Float() == 0

	case reflect.Complex64,reflect.Complex128:
		return t.Value.Complex() == 0+0i

	case reflect.Bool:
		return t.Value.Bool() == false

	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface:
		return t.Value.IsNil()

	case reflect.Slice, reflect.Map:
		return t.Value.Len() == 0

	case reflect.Struct:
		return reflect.DeepEqual(t.Value.Interface(), reflect.Zero(t.Value.Type()).Interface())
	}

	panic(fmt.Errorf("Unknown value kind %T", t.Value))
}

func (t *Type) Uint64() (uint64,error) {
	if t.Value.Kind() == reflect.Int64 {
		return t.Value.Uint(), nil
	}
	return 0, TypeError
}

func (t *Type) Uint64Or(value uint64) uint64 {
	i,err := t.Uint64()
	if err != nil{
		return value
	}
	return i
}

func (t *Type) Int32() (int32,error) {
	if t.Value.Kind() == reflect.Int32 {
		return int32(t.Value.Int()), nil
	}
	return 0, TypeError
}

func (t *Type) Int32Or(value int32) int32 {
	i,err := t.Int32()
	if err != nil{
		return value
	}
	return i
}

func (t *Type) Int() (int,error) {
	if t.Value.Kind() == reflect.Int {
		return int(t.Value.Int()), nil
	}
	return 0, TypeError
}

func (t *Type) IntOr(value int) int {
	i,err := t.Int()
	if err != nil{
		return value
	}
	return i
}

func (t *Type) Int64() (int64,error) {
	if t.Value.Kind() == reflect.Int64 {
		return t.Value.Int(), nil
	}
	return 0, TypeError
}

func (t *Type) Int64Or(value int64) int64 {
	i,err := t.Int64()
	if err != nil{
		return value
	}
	return i
}

func (t *Type) Float32() (float32,error) {
	if t.Value.Kind() == reflect.Float32 {
		return float32(t.Value.Float()), nil
	}
	return 0.0, TypeError
}

func (t *Type) Float32Or(value float32) float32 {
	f,err := t.Float32()
	if err != nil{
		return value
	}
	return f
}

func (t *Type) Float64() (float64,error) {
	if t.Value.Kind() == reflect.Float64 {
		return t.Value.Float(), nil
	}
	return 0.0, TypeError
}

func (t *Type) Float64Or(value float64) float64 {
	f,err := t.Float64()
	if err != nil{
		return value
	}
	return f
}

func (t *Type) Bool() (bool,error) {
	if t.Value.Kind() == reflect.Bool {
		return t.Value.Bool(), nil
	}
	return false, TypeError
}

func (t *Type) BoolOr(value bool) bool {
	b,err := t.Bool()
	if err != nil{
		return value
	}
	return b
}

func (t *Type) String() (string,error) {
	if t.Value.Kind() == reflect.String {
		return t.Value.String(), nil
	}
	return "", TypeError
}

func (t *Type) StringOr(value string) string {
	s,err := t.String()
	if err != nil{
		return value
	}
	return s
}

func (t *Type) Interface() interface{} {
	return t.Value.Interface()
}

func (t *Type) UIntSlice() []uint {
	slice :=  t.Slice(reflect.Uint)
	if slice != nil{
		return slice.([]uint)
	}
	return nil
}

func (t *Type) UInt64Slice() []uint64 {
	slice :=  t.Slice(reflect.Uint64)
	if slice != nil{
		return slice.([]uint64)
	}
	return nil
}

func (t *Type) IntSlice() []int {
	slice :=  t.Slice(reflect.Int)
	if slice != nil{
		return slice.([]int)
	}
	return nil
}

func (t *Type) Int32Slice() []int32 {
	slice :=  t.Slice(reflect.Int32)
	if slice != nil{
		return slice.([]int32)
	}
	return nil
}

func (t *Type) Int64Slice() []int64 {
	slice :=  t.Slice(reflect.Int64)
	if slice != nil{
		return slice.([]int64)
	}
	return nil
}

func (t *Type) Float32Slice() []float32 {
	slice :=  t.Slice(reflect.Float32)
	if slice != nil{
		return slice.([]float32)
	}
	return nil
}

func (t *Type) Float64Slice() []float64 {
	slice :=  t.Slice(reflect.Float64)
	if slice != nil{
		return slice.([]float64)
	}
	return nil
}

func (t *Type) Slice(sliceKind reflect.Kind) interface{} {
	kind :=t.Value.Kind()
	if kind == reflect.Slice  {
		if !t.Value.IsNil() {
			if t.Value.Index(0).Kind() == sliceKind {
				return t.Value.Interface()
			}
		}
	}else if kind == reflect.Ptr {
		if !t.Value.IsNil() {
			if t.Value.Elem().Index(0).Kind() == sliceKind {
				return t.Value.Elem().Interface()
			}
		}
	}
	return nil
}
