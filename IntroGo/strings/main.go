package main

import "fmt"

func main() {
	//  atoz := "the quick brown fox jumps over the lazy dog\n"

	temp := "こんにちは世界"

	fmt.Println(len(temp))

	// go through a string rune by rune
	for i, r := range temp {
		fmt.Printf("%d %c\n", i, r)
	}
}
