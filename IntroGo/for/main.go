package main

import "fmt"

func main() {
	var stop bool

	i := 0

	for !stop {
		fmt.Printf("Hello, World!\n")
		i++
		if i == 10 {
			stop = true
		}
	}
}
