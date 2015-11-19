package reflect

import (
	"fmt"
	"reflect"
)

func GetStructField(structPtr interface{}, name string) (interface{},error) {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return nil,fmt.Errorf("structPtr must be pointer to a struct, but is %T", structPtr)
	}
	v = v.Elem()
	if f := v.FieldByName(name); f.IsValid() {
		return f.Interface(),nil
	}
	return nil,fmt.Errorf("%T has no struct field '%s'", v.Interface(), name)
}

func SetStructField(structPtr interface{}, name string, value interface{}) error {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("structPtr must be pointer to a struct, but is %T", structPtr)
	}
	v = v.Elem()

	if f := v.FieldByName(name); f.IsValid() {
		i := reflect.ValueOf(value)
		f.Set(i)

	} else {
		return fmt.Errorf("%T has no struct field '%s'", v.Interface(), name)
	}

	return nil
}

func SetStructFieldStr(structPtr interface{}, name, value string) error {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("structPtr must be pointer to a struct, but is %T", structPtr)
	}
	v = v.Elem()

	if f := v.FieldByName(name); f.IsValid() {
		if f.Kind() == reflect.String {
			f.SetString(value)
		} else {
			_, err := fmt.Sscan(value, f.Addr().Interface())
			if err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("%T has no struct field '%s'", v.Interface(), name)
	}

	return nil
}




func IsZero(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Bool:
		return v.Bool() == false

	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface,reflect.Slice, reflect.Map:
		return v.IsNil()

	case reflect.Struct:
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	}

	panic(fmt.Errorf("Unknown value kind %T", value))
}


func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Bool:
		return v.Bool() == false

	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface:
		return v.IsNil()

	case reflect.Slice, reflect.Map:
		return v.Len() == 0

	case reflect.Struct:
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	}

	panic(fmt.Errorf("Unknown value kind %T", value))
}

/**
利用了反射，性能会下降，性能要求高建议不用
*/
func StructToMap(v interface{}) map[string]interface{} {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		panic(fmt.Errorf("Expected a struct, got %s", t))
	}
	result := make(map[string]interface{})
	structToMap(v, t, result)
	return result
}

func structToMap(value interface{}, t reflect.Type, result map[string]interface{}) {
	v := reflect.ValueOf(value)
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		if structField.Anonymous && structField.Type.Kind() == reflect.Struct { //匿名类
			structToMap(v.Field(i).Interface(), structField.Type, result)
		} else if !structField.Anonymous && structField.Type.Kind() == reflect.Struct { //struct类型
			result[structField.Name] = StructToMap(v.Field(i).Interface())
		} else { //普通类型
			result[structField.Name] = v.Field(i).Interface()
		}
	}
}

/**
利用了反射，性能会下降，性能要求高建议不用
*/
func StrMapToStruct(structPtr interface{}, m map[string]string, errOnMissingField bool) error {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("structPtr must be pointer to a struct, but is %T", structPtr)
	}
	v = v.Elem()

	for name, value := range m {
		if f := v.FieldByName(name); f.IsValid() {
			if f.Kind() == reflect.String {
				f.SetString(value)
			} else {
				_, err := fmt.Sscan(value, f.Addr().Interface())
				if err != nil {
					return err
				}
			}
		} else if errOnMissingField {
			return fmt.Errorf("%T has no struct field '%s'", v.Interface(), name)
		}
	}

	return nil
}

/**
利用了反射，性能会下降，性能要求高建议不用
*/
func MapToStruct(structPtr interface{}, m map[string]interface{}, errOnMissingField bool) error {
	v := reflect.ValueOf(structPtr)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("structPtr must be pointer to a struct, but is %T", structPtr)
	}
	v = v.Elem()

	for name, value := range m {
		f := v.FieldByName(name)
		fmt.Println("f:", f)
		if f.IsValid() {
			fieldValue := reflect.ValueOf(value)
			f.Set(fieldValue)
		} else if errOnMissingField {
			return fmt.Errorf("%T has no struct field '%s'", v.Interface(), name)
		}
	}
	return nil
}






