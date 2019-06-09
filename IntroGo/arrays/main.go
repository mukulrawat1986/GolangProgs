package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

func main() {
	// a slice
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	// print out the length of the slice
	fmt.Printf("%d\n", len(words))

	printer(words)
	printer(words)
}
