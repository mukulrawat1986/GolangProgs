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

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
