package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	// get the concrete value stored in the interface x
	val := reflect.ValueOf(x)

	// get the value of the 0th field stored in the struct val,
	// assuming val is a struct
	field := val.Field(0)

	// return the string's underlying value as a string
	fn(field.String())
}
