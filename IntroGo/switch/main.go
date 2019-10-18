package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!\n")

	n = 0

	// switch in general
	switch {
	case err != nil:
		os.Exit(1)

	case n == 0:
		fmt.Printf("No bytes output\n")
		fallthrough

	case n != 14:
		fmt.Printf("Wrong number of characters\n")

	default:
		fmt.Printf("OK!\n")
	}
}
