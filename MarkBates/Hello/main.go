package main

import "fmt"

func Greet(name string) {
	fmt.Println("Hello, " + name)
}

func GreetNames(names []string) {
	for _, name := range names {
		Greet(name)
	}
}

func main() {
	names := []string{
		"Mark",
		"Rachel",
		"Dylan",
		"Leo",
	}

	GreetNames(names)
}
