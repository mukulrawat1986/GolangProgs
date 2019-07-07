package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	// get the concrete value stored in the interface
	val := reflect.ValueOf(x)

	// iterate over the number of fields stored in x, assuming x is a struct
	for i := 0; i < val.NumField(); i++ {
		// get the ith field of the struct
		field := val.Field(i)

		// switch on the kind of field
		switch field.Kind() {
		case reflect.String:
			// apply the function
			fn(field.String())

		case reflect.Struct:
			// recursively call the walk function
			walk(field.Interface(), fn)
		}
	}
}

func main() {
}
