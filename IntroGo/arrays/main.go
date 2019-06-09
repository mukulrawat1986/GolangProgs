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
	words := make([]string, 4)

	words[0] = "The"
	words[1] = "Quick"
	words[2] = "Brown"
	words[3] = "Fox"

	printer(words)
}
