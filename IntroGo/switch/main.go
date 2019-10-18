package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!\n")

	// switch in general
	switch {
	case err != nil:
		os.Exit(1)

	case n == 0:
		fmt.Printf("No bytes output")

	case n != 14:
		fmt.Printf("Wrong number of characters")

	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")
}
