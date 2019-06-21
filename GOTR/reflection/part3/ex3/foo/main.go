package foo

import (
	"fmt"
	"reflect"
	"strings"
)

func Println(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	// fmt.Printf("%v %v\n", t, v)

	switch t.Kind() {
	case reflect.Struct:
		s := printStructExpanded(x)
		fmt.Println(s)
	case reflect.Ptr:
		v1 := reflect.Indirect(v) // get the value being pointed by pointer
		s := printStructExpanded(v1.Interface())
		fmt.Printf("&%v\n", s)
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println()
}

func printStructExpanded(x interface{}) string {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	// fmt.Printf("%v: %v\n", t, v)
	if t.Kind() != reflect.Struct {
		fmt.Print("Err: Not a struct, expected struct value of 'kind' struct\n")
		return ""
	}

	n := t.NumField()
	s := "%v: %v"
	res := strings.Builder{}
	fmt.Fprintf(&res, "%v{", t)
	for i:= 0; i < n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)

		fmt.Fprintf(&res, s, tt.Name, vv)
		s = ", " + "%v: %v"
	}
	fmt.Fprintln(&res, "}")

	return res.String()
}
