package array_and_slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("make the sum of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	numbers1 := []int{1, 5, 3, 21, 67, 54, 43, 2, 324, 4365, 34576, 54}
	numbers2 := []int{1, 23, 4, 4236, 5437, 2, 3, 32, 41, 26, 32, 674, 75664, 3}
	numbers3 := []int{233, 234, 1, 23, 25, 346, 2}
	numbers4 := []int{234, 23, 6, 64, 68, 5, 34, 6, 23, 34}

	for b.Loop() {
		SumAll(
			numbers1,
			numbers2,
			numbers3,
			numbers4)
	}
}
