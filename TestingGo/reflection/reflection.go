package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	// get the concrete value stored in the interface
	val := reflect.ValueOf(x)

	// we are assuming its a struct and looping over the number of fields
	for i := 0; i < val.NumField(); i++ {
		// get the individual field value
		field := val.Field(i)

		// check if the specific type of field is a String
		if field.Kind() == reflect.String {
			// apply the function to this field, assuming its a string
			fn(field.String())
		}

	}
}

func main() {
}
