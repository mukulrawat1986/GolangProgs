package main

import "github.com/mukulrawat1986/GolangProgs/GOTR/reflection/ex5/foo"

type (
	ID     string
	Person struct {
		name string
	}
)

func main() {
	foo.Println(true)
	foo.Println(2010)
	foo.Println(9.14)
	foo.Println(7 + 7i)
	foo.Println("Hello, World!")
	foo.Println(ID("fkdjfd"))
	foo.Println([5]byte{})
	foo.Println([]byte{})
	foo.Println(map[string]int{})
	foo.Println(Person{name: "Anna"})
	foo.Println(&Person{name: "Anna"})
	foo.Println(make(chan int))
	foo.Println(nil)
	var i interface{}

	foo.Println(i)
}
