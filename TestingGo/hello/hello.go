package main

import "fmt"

// Hello returns a string
func Hello(name string) string {
	return "Hello, World!"
}

func main() {
	fmt.Println(Hello("Chris"))
}
