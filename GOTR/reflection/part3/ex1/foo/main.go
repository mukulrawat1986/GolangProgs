package foo

import (
	"fmt"
	"reflect"
)

func Println(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	switch t.Kind() {
	case reflect.Bool:
		fmt.Printf("boolean value '%v'\n", v)

	case reflect.Int:
		fmt.Printf("Integer value '%v'\n", v)

	case reflect.Float64:
		fmt.Printf("Float value '%v'\n", v)

	case reflect.Complex128:
		fmt.Printf("Complex value '%v'\n", v)

	case reflect.String:
		fmt.Printf("String value '%v'\n", v)

	case reflect.Struct:
		printStructExpanded(x)

	case reflect.Chan:
		fmt.Println(v.String())

	case reflect.Array:
		fmt.Printf("Array value '%v'\n", v)

	case reflect.Slice:
		fmt.Printf("Slice value '%v'\n", v)

	case reflect.Map:
		fmt.Printf("Map value '%v'\n", v)

	default:
		fmt.Println("Unknown type")
	}

	fmt.Printf("\n")
}

func printStructExpanded(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t.Kind() != reflect.Struct {
		fmt.Printf("Err: Not a struct, expected struct value of 'kind' struct\n")
		return
	}

	n := t.NumField()
	s := "%v: %v"
	fmt.Printf("'%v'{", t)
	for i:=0; i<n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)
		fmt.Printf(s, tt.Name, vv)
		s = ", " + "%v: %v"
	}
	fmt.Printf("}\n")
}
