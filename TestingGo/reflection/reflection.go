package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {

	val := reflect.ValueOf(x)

	// Numfield returns the number of fields in the struct v, panics if v's kind is not Struct
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fn(field.String())
	}
}
