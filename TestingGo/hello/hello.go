package main

import "fmt"

// Hello function returns a string greeting
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("World"))
}
