package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	// extract the underlying concrete value from the interface
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {

	// if its a string, apply the function
	case reflect.String:
		fn(val.String())

	// if its a struct, call walk on each field in struct
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}

	// if its a slice, call walk on each index in the slice
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	// if its a map, iterate over the keys and call walk for each key
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
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
