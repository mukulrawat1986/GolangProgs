package arrays

// Sum returns the sum of an array of integers of length 5
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
