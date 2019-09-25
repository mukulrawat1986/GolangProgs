package arrays

// Sum returns the sum of a slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll returns a slice which contains the sum of slices
func SumAll(numbersToSum ...[]int) []int {
	return []int{}
}
