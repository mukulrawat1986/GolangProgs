package main

import (
	"fmt"
	"time"
)

func main() {

	queue := make(chan string)

	go func(queue chan string) {
		fmt.Println("func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("func 1")
		queue <- "done func 1"
	}(queue)

	go func(queue chan string) {
		fmt.Println("func 2")
		time.Sleep(1 * time.Second)
		fmt.Println("func 2")
		queue <- "done func 2"
	}(queue)

	// fmt.Println(<-queue)
	// fmt.Println(<-queue)

	for text := range queue {
		fmt.Println(text)
	}
}
