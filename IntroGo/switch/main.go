package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!\n")

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output\n")
	case n != 14:
		fmt.Printf("Wrong number of characters\n")
	default:
		fmt.Printf("OK!\n")
	}

	// switching over a variable
	atoz := "the quick brown fox jumps over the lazy dog"

	vowels, consonants, zeds := 0, 0, 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		case 'z':
			zeds += 1
			fallthrough
		default:
			consonants += 1
		}
	}

	fmt.Printf("Vowels: %d; Consonants: %d (Zeds: %d)\n", vowels, consonants, zeds)
}
