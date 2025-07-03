package iteration

import (
	"fmt"
	"testing"
)

//goland:noinspection ALL
func TestRepeat(t *testing.T) {
	t.Run("iterate a 5 times", func(t *testing.T) {
		inputChar := "a"
		repeatCount := 5

		repeated := Repeat(inputChar, repeatCount)
		//goland:noinspection SpellCheckingInspection
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q, want %q", repeated, expected)
		}
	})

	t.Run("iterate c 15 times", func(t *testing.T) {
		inputChar := "c"
		repeatCount := 7

		repeated := Repeat(inputChar, repeatCount)
		//goland:noinspection SpellCheckingInspection
		expected := "ccccccc"

		if repeated != expected {
			t.Errorf("got %q, want %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 50)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 4)
	fmt.Println(repeat)
	// output: aaaa
}
