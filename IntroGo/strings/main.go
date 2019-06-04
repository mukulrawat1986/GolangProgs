package main

import "fmt"

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	// make substrings
	fmt.Printf("%s\n", atoz[:9])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[15:])
}
