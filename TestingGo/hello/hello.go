package main

import "fmt"

const (
	exclamation        = "!"
	englishHelloPrefix = "Hello, "
)

// Hello returns a string
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name + exclamation
}

func main() {
	fmt.Println(Hello("Chris"))
}
