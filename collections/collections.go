package collections

import (
	"reflect"
)

func Contains(collection interface{}, element interface{}) bool {
	collectionValue := reflect.ValueOf(collection)
	for i := 0; i < collectionValue.Len(); i++ {
		if collectionValue.Index(i).Interface() == element {
			return true
		}
	}
	return false
}
