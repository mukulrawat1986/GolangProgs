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

	// switching on a variable
	atoz := "the quick brown fox jumps over the lazy dog"

	vowels := 0
	consonants := 0
	zeds := 0

	// iterate over the runes by using for loop
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

	fmt.Printf("Vowels: %d Consonants: %d (Zeds: %d)\n", vowels, consonants, zeds)
}
