package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
)

// Hello function returns a string greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
