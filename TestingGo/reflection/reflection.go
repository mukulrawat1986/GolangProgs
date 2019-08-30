package _reflection

import "reflect"

func walk(x interface{}, fn func(string)) {
	// get the concrete value stored in the interface
	val := reflect.ValueOf(x)

	// Get the 0th field of val, it panics if val's kind is not struct or index is out
	// of range
	field := val.Field(0)
	fn(field.String())
}
