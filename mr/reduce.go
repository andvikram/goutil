// mr = MapReduce

package mr

import "errors"

// ReduceIntFunc is required to reduce int values.
type ReduceIntFunc func(int, int) int

// ReduceStringFunc is required to reduce string values.
type ReduceStringFunc func(string, string) string

// Reduce is a function to operate on all values of a list of any primitive type values
// as per the provided or default function.
// ----------
// Parameters required are: an initial value, a list of values to reduce and
// a function implementation for defining how to reduce,
// for all of which the value types must be the same.
// However, function implementation is optional and can be replaced with nil,
// in which case the default implementation is to fold with addition.
func Reduce(list, initVal, funk interface{}) (interface{}, error) {
	var ok bool

	switch v := initVal.(type) {
	case int:
		var l []int
		var f ReduceIntFunc
		if l, ok = list.([]int); !ok {
			return nil, errors.New("parameter values must be of the same type")
		}
		if funk == nil {
			var reduceFunc ReduceIntFunc
			reduceFunc = func(x, y int) int {
				return x + y
			}
			return reduceInt(l, v, reduceFunc), nil
		}
		if f, ok = funk.(ReduceIntFunc); !ok {
			return nil, errors.New("parameter values must be of the same type")
		}
		return reduceInt(l, v, f), nil
	case string:
		var l []string
		var f ReduceStringFunc
		var ok bool
		if l, ok = list.([]string); !ok {
			return nil, errors.New("parameter values must be of the same type")
		}
		if funk == nil {
			var reduceFunc ReduceStringFunc
			reduceFunc = func(x, y string) string {
				return x + y
			}
			return reduceString(l, v, reduceFunc), nil
		}
		if f, ok = funk.(ReduceStringFunc); !ok {
			return nil, errors.New("parameter values must be of the same type")
		}
		return reduceString(l, v, f), nil
	default:
		return nil, errors.New("initVal is invalid")
	}
}

func reduceInt(list []int, memo int, funk ReduceIntFunc) interface{} {
	for i := 0; i < len(list); i++ {
		memo = funk(memo, list[i])
	}
	return memo
}

func reduceString(list []string, memo string, funk ReduceStringFunc) interface{} {
	for i := 0; i < len(list); i++ {
		memo = funk(memo, list[i])
	}
	return memo
}
