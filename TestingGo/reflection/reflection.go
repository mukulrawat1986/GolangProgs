package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	// extract the underlying concrete value from the interface
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {

	// if its a string, apply the function
	case reflect.String:
		fn(val.String())

	// if its a struct, call walk on each field in struct
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field

	// if its a slice, call walk on each index in the slice
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index

	}

	// iterate over the struct or slice
	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
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
