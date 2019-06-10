package main

import (
	"fmt"
)

func emit(c chan string) {
	words := []string{"The", "quick", "brown", "fox"}

	for _, word := range words {
		c <- word
	}

	close(c)
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)

	for {
		if word, ok := <-wordChannel; ok != false {
			fmt.Println(word)
		} else {
			break
		}
	}

	fmt.Printf("\n")
}
