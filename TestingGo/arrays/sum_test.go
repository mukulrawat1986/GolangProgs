package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {

	assertCorrectSum := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d', given '%v'", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		assertCorrectSum(t, Sum(numbers), 15, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		assertCorrectSum(t, Sum(numbers), 6, numbers)
	})
}
