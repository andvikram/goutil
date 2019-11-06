// Package mr provides functions to implement ReduceReduce
package mr

import (
	"errors"
	"reflect"
)

// ReduceFunc is required to reduce values.
// It takes a slice of any type and returns
// a slice reduced values of the same type with error.
type ReduceFunc func(interface{}) (interface{}, error)

// Reduce operates on all values of a list of any type values
// as per the provided function.
//
// Parameters required are: a slice of any type and
// a function implementation of ReduceFunc for defining how to reduce.
func Reduce(list interface{}, funk ReduceFunc) (interface{}, error) {
	t := reflect.TypeOf(list)
	if t.Kind() != reflect.Slice {
		return nil, errors.New("list is not a slice")
	}
	return funk(list)
}
