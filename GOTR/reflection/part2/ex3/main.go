package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}
	x = &struct{ name string }{}

	t0 := reflect.TypeOf(x)
	v0 := reflect.ValueOf(x)
	fmt.Printf("x: type= %v, value = %v\n", t0, v0)

	x = new(string)

	t1 := reflect.TypeOf(x)
	v1 := reflect.ValueOf(x)
	fmt.Printf("x: type = %v, value = %v\n", t1, v1)

	// What kind or category is the type a member of?
	fmt.Printf("t0: type = %v, kind = %v\n", t0, t0.Kind())
	fmt.Printf("t1: type = %v, kind = %v\n", t1, t1.Kind())

	x = [5]int{}
	fmt.Printf("x: type = %v, kind = %v\n", reflect.TypeOf(x), reflect.TypeOf(x).Kind())

	x = [10]float32{}
	fmt.Printf("x: type = %v, kind = %v\n", reflect.TypeOf(x), reflect.TypeOf(x).Kind())

	x = []int{}
	fmt.Printf("x: type = %v, kind = %v\n", reflect.TypeOf(x), reflect.TypeOf(x).Kind())

	x = []string{}
	fmt.Printf("x: type = %v, kind = %v\n", reflect.TypeOf(x), reflect.TypeOf(x).Kind())

}
