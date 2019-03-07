package ArraySlice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	assertSum := func(t *testing.T, got, want int, numbers []int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertSum(t, got, want, numbers)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		assertSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	assertSlice(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}
	assertSlice(t, got, want)
}

func assertSlice(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
