package main

import "github.com/mukulrawat1986/GolangProgs/GOTR/reflection/part3/ex1/foo"

type Person struct {
	name string
	age int
	sex string
	height float64
	weight float64
}

type ID string

func main() {
	foo.Println(true)
	foo.Println(2010)
	foo.Println(9.15)
	foo.Println(7 + 7i)
	foo.Println("Hello, World!")
	foo.Println(ID("19590925"))
	foo.Println([5]byte{})
	foo.Println([]byte{})
	foo.Println(map[string]int{})
	foo.Println(Person{name: "Jane Doe", age: 18, sex: "female", height: 1.78, weight: 145.5})
	foo.Println(Person{name: "John Doe", age: 18, sex: "male", height: 1.82, weight: 178.7})
	foo.Println(make(chan int))
}