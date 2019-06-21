package foo

import (
	"fmt"
	"reflect"
)

func Println(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("%v %v\n", t, v)

	switch t.Kind() {
	case reflect.Struct:
		printStructExpanded(x)
	case reflect.Ptr:
		v1 := reflect.Indirect(v) // get the value being pointed by pointer
		printStructExpanded(v1.Interface())
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println()
}

func printStructExpanded(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("%v: %v\n", t, v)
	if t.Kind() != reflect.Struct {
		fmt.Print("Err: Not a struct, expected struct value of 'kind' struct\n")
		return
	}

	n := t.NumField()
	s := "%v: %v"
	for i:= 0; i < n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)

		fmt.Printf(s, tt.Name, vv)
		s = ", " + "%v: %v"
	}
	fmt.Println()
}
