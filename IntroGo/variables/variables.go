package main

import "fmt"

var (
	message = "The answer to life is %d\n"
	answer  = 42
)

func main() {
	answer++
	fmt.Printf(message, answer)
}
