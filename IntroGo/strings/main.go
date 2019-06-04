package main

import "fmt"

func main() {
	atoz := `the quick brown fox jumps over the lazy dog\n`

	fmt.Printf("%s\n", atoz)

	// go through a string rune by rune
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	// len returns the number of bytes in a string, which in case of ascii
	// is the length of the string
	fmt.Printf("%d\n", len(atoz))
}
