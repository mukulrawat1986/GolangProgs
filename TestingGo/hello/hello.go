package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	spanishHelloPrefix = "Hola, "
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
