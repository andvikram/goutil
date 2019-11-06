// Package mr provides functions to implement MapReduce
package mr

import (
	"errors"
	"reflect"
)

// MapFunc is required to map values.
// It takes a slice of any type
// and returns a slice of any type with error.
type MapFunc func(interface{}) (interface{}, error)

// Map operates on all values of a list of any type values
// as per the provided function.
//
// Parameters required are: a slice of any type and
// a function implementation of MapFunc for defining how to map.
func Map(list interface{}, funk MapFunc) (interface{}, error) {
	t := reflect.TypeOf(list)
	if t.Kind() != reflect.Slice {
		return nil, errors.New("list is not a slice")
	}
	return funk(list)
}
