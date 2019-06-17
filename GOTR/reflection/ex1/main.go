package main

import "fmt"

type (
	ID string
	Person struct {
		name string
	}
)

func main() {
	fmt.Println(true)
	fmt.Println(2010)
	fmt.Println(9.14)
	fmt.Println(7 + 7i)
	fmt.Println("Hello, World!")
	fmt.Println(ID("fkdjfd"))
	fmt.Println([5]byte{})
	fmt.Println([]byte{})
	fmt.Println(map[string]int{})
	fmt.Println(Person{name: "Anna"})
	fmt.Println(&Person{name: "Anna"})
	fmt.Println(make(chan int))
}
