package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	// extract the concrete value from inside the interface
	val := getValue(x)

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

func getValue(x interface{}) reflect.Value {
	// extract the concrete value from inside the interface
	val := reflect.ValueOf(x)

	// check if val is of pointer kind if so extract the value being pointed by it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func main() {
}
