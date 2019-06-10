package main

import "fmt"

func main() {
	dayMonths := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	fmt.Printf("Days in Feb: %d\n", dayMonths["Feb"])

	days, ok := dayMonths["January"]
	if !ok {
		fmt.Printf("Can't get days for January\n")
	} else {
		fmt.Printf("%d\n", days)
	}

	has28 := 0

	delete(dayMonths, "Feb")

	// iterating over a map
	for month, days := range dayMonths {
		if days == 28 {
			fmt.Println(month)
			has28++
		}
	}

	fmt.Printf("%d months have 28 days\n", has28)
}
