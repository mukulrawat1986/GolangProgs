package _reflection

import "reflect"

func walk(x interface{}, fn func(string)) {
	// get the concrete value stored in the interface
	val := reflect.ValueOf(x)

	// iterate over all fields in the struct val
	for i := 0; i < val.NumField(); i++ {

		// get the ith field of the struct val
		// it panics if val's Kind is not struct or i is out of range
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())

		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
