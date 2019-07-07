package main

import "reflect"

func walk(x interface{}, fn func(input string)) {

	// get the underlying concrete valye stored in the interface
	value := reflect.ValueOf(x)

	// get the first and only field
	field := value.Field(0)

	// apply the function to the string field
	fn(field.String())
}

func main() {
}
