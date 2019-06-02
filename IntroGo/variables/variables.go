package main

import "fmt"

const (
	message = "%d %d %d\n"
	answer1 = iota + 11
	answer2
	answer3
)

func main() {
	fmt.Printf(message, answer1, answer2, answer3)
}
