package main

import "fmt"

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main() {
	// creating a slice and adding strings to it
	// create a slice of 0 length (means there is nothing in it),
	// and capacity of 4
	words := make([]string, 0, 4)

	fmt.Printf("%d %d\n", len(words), cap(words))
	words = append(words, "The")
	words = append(words, "Quick")
	words = append(words, "Brown")
	words = append(words, "Fox")
	fmt.Printf("%d %d\n", len(words), cap(words))

	printer(words)

	words = append(words, "Jumps")
	fmt.Printf("%d %d\n", len(words), cap(words))
}
