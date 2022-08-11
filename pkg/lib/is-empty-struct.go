package lib

import (
	"errors"
	"reflect"
)

// heavy invoke
func IsEmptyStruct(someStruct interface{}) (bool, error) {
	value := reflect.ValueOf(someStruct)
	if value.Kind() != reflect.Struct {
		return true, errors.New("value is not a struct")
	}
	return value.IsZero(), nil
}
