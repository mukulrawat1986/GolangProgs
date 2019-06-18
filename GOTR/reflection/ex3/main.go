package main

import "fmt"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	Println(true)
	Println(2010)
	Println(9.14)
	Println(7 + 7i)
	Println("Hello, World!")
	Println(ID("fkdjfd"))
	Println([5]byte{})
	Println([]byte{})
	Println(map[string]int{})
	Println(Person{name: "Anna"})
	Println(&Person{name: "Anna"})
	Println(make(chan int))
	Println(nil)
	var i interface{}

	Println(i)
}

// Println is my simple print function
// Example of Type Assertion
func Println(x interface{}) {
	// x.Type for type information
	// x.Value for the value assigned
	if v, ok := x.(ID); !ok {
		fmt.Printf("'%T' is not the type that I want\n", x)
	} else {
		fmt.Printf("x has type ID, which I defined. The value is '%v'\n", v)
	}
}
