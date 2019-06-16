package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x) // returns a new value initialized to the concrete value stored in the
							  // interface i

	field := val.Field(0) // returns the ith field of the struct, panics if its not a struct or i
							  // is out of range

	fn(field.String()) // returns the underlying value as a string, panics if its not a string
}
