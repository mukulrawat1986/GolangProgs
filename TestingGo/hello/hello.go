package main

import "fmt"

const (
	french             = "French"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	} else if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
