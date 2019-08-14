package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("func 1")
	}()

	go func() {
		fmt.Println("func 2")
		time.Sleep(1 * time.Second)
		fmt.Println("func 2")
	}()

	time.Sleep(4 * time.Second)
}
