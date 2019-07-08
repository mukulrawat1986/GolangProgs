package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	// extract the underlying concrete value from the interface
	val := getValue(x)

	switch val.Kind() {

	// if its a struct, call walk on each field in struct
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}

	// if its a slice, call walk on each index in the slice
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}

	// if its a string, apply the function
	case reflect.String:
		fn(val.String())
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
