package main

// Sum returns the sum of the array
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll function takes varying number of int slices, returning a new slice
// containing the total for each slice passed in
func SumAll(numbersToSum ...[]int) []int {
	return []int{}
}

func main() {

}
