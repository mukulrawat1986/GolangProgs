package main

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

// Hello function returns a string greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	} else if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
