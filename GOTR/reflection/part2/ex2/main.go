package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = 3.14

	// using reflect get the Type and Value
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("x : Type = %v, Value = %v\n", t, v)

	goo := x
	fmt.Printf("goo: Type = %T, Value = %v\n", goo, goo)

	x = &struct{ name string }{}

	t = reflect.TypeOf(x)
	v = reflect.ValueOf(x)

	fmt.Printf("x : Type = %v, Value = %v\n", t, v)

	hoo := x
	fmt.Printf("hoo: Type = %T, Value = %v\n", hoo, hoo)

	var r interface{}
	fmt.Printf("r : Type : %T, Value : %v\n", r, r)
}
