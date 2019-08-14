package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	queue := make(chan string)

	var w sync.WaitGroup

	// wait for two goroutines
	w.Add(2)

	go func(queue chan string, w *sync.WaitGroup) {
		defer w.Done()
		fmt.Println("func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("func 1")
		queue <- "done func 1"
	}(queue, &w)

	go func(queue chan string, w *sync.WaitGroup) {
		defer w.Done()
		fmt.Println("func 2")
		time.Sleep(1 * time.Second)
		fmt.Println("func 2")
		queue <- "done func 2"
	}(queue, &w)

	go func(queue chan string) {
		for text := range queue {
			fmt.Println(text)
		}
	}(queue)
	w.Wait()
}
