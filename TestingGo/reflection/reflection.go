package _reflection

import "reflect"

func walk(x interface{}, fn func(string)) {

	val := getValue(x)

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

func getValue(x interface{}) reflect.Value {
	// get the concrete value stored in the interface
	val := reflect.ValueOf(x)

	// check if the concrete value is of pointer type
	if val.Kind() == reflect.Ptr {
		// get the value pointed to by the pointer
		val = val.Elem()
	}

	return val
}
