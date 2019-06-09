package main

import "fmt"

func printer(w [4]string) {
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
	w[2] = "blue"
}

func main() {
	// array with unknown number of values of strings
	words := [...]string{"the", "quick", "brown", "fox"}
	printer(words)
	printer(words)
}
