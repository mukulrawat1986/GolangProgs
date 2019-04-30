package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
