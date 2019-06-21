package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

type Proto struct {
	src   string
	dst   string
	size  uint
	magic int
}

func main() {
	printStructInfo(3.14)
	printStructInfo(Person{})
	printStructInfo(Proto{})
}

func printStructInfo(x interface{}) {
	t := reflect.TypeOf(x)

	fmt.Printf("-------Kind - %v-------\n", t.Kind())

	if t.Kind() != reflect.Struct {
		fmt.Printf("Err: Not a struct, expected struct value of 'kind' struct\n")
		return
	}

	n := t.NumField()
	fmt.Printf("struct of type %v has %v fields.\n", t, n)
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		fmt.Printf("Field %d: %v, name: %v, type: %v\n", i, tt, tt.Name, tt.Type)
	}

}
