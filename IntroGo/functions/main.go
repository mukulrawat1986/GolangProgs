package main

import "fmt"

func printer(msg string) (err error) {
	_, err = fmt.Printf("Hello, World\n")
	return
}

func main() {
	printer("Hello, World!")
}
