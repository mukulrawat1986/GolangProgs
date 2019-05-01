package main

// Sum will take an array of numbers and return the sum
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {

}
