package _reflection

import "reflect"

func walk(x interface{}, fn func(string)) {

	val := getValue(x)

	// check the type of val
	switch val.Kind() {

	// if value is of struct type
	case reflect.Struct:
		// iterate over the fields of the struct
		for i := 0; i < val.NumField(); i++ {
			//  call the walk function for each field
			walk(val.Field(i).Interface(), fn)
		}

	// if value if of slice type
	case reflect.Slice:
		// iterate over the elements of the slice
		for i := 0; i < val.Len(); i++ {
			// call the walk function for each element
			walk(val.Index(i).Interface(), fn)
		}

	// if value is of string type
	case reflect.String:
		fn(val.String())
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
