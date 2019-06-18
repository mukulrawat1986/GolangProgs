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

func Println(x interface{}) {
	fmt.Printf("type is '%T', value is '%v'\n", x, x)
}
