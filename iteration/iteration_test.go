package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("iterate a 5 times", func(t *testing.T) {
		inputChar := "a"
		repeatCount := 5

		repeated := Repeat(inputChar, repeatCount)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q, want %q", repeated, expected)
		}
	})

	t.Run("iterate c 15 times", func(t *testing.T) {
		inputChar := "c"
		repeatCount := 7

		repeated := Repeat(inputChar, repeatCount)
		expected := "ccccccc"

		if repeated != expected {
			t.Errorf("got %q, want %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 15)
	}
}
