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

	words = append(words, "The")
	words = append(words, "Quick")
	words = append(words, "Brown")
	words = append(words, "Fox")

	printer(words)

	newWords := make([]string, 4)
	copy(newWords, words)
	printer(newWords)

}
