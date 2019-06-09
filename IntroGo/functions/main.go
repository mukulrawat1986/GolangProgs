package main

import "fmt"

func printer(msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	printer("Hello", "this", "is", "dog")
}
