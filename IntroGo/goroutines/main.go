package main

import (
	"fmt"
)

func emit(c chan int) {
	for {

	}
}

func main() {
	wordChannel := make(chan string)

	go emit(wordChannel)
	go emit(wordChannel)

	for word := range wordChannel {
		fmt.Printf("%s\n", word)
	}

}
