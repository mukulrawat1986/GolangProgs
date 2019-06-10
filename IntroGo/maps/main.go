package main

import "fmt"

func main() {
	dayMonths := make(map[string]int)

	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

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
