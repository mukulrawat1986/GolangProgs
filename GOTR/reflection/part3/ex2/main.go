package main

import "github.com/mukulrawat1986/GolangProgs/GOTR/reflection/part3/ex2/foo"

type Person struct {
	name string
	age int
}

func main() {
	foo.Println(Person{name: "John Doe"})
	foo.Println(&Person{name: "John Doe"})
}
