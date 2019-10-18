package main

import "fmt"

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	// take substring
	fmt.Printf("%s\n", atoz[:9])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[15:])

	// going over the string rune by rune
	// we doing it by using a for loop
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	// number of bytes in the string
	fmt.Printf("The number of bytes in string: %d\n", len(atoz))
}
