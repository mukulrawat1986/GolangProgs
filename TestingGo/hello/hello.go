package main

import "fmt"

const (
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

// Hello function returns a string greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	} else if language == spanish {
		return spanishHelloPrefix + name
	} else if language == "French" {
		return "Bonjour, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
