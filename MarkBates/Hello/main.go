package main

import "fmt"

func Greet(name string) {
	fmt.Println("Hello, " + name)
}

func GreetNames(names []string, suffix string) {
	for _, name := range names {
		Greet(name + suffix)
	}
}

func main() {
	names := []string{
		"Mark",
		"Rachel",
		"Dylan",
		"Leo",
	}

	go GreetNames(names, " <C> ")
	GreetNames(names, " <M> ")
}
