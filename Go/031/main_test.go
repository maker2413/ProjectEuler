package main

import "testing"

func TestPossibleDenominations(t *testing.T) {
	t.Run("No coins, no total", func(t *testing.T) {
		got := possibleDenominations([]int{}, 0)
		want := 0

		if got != want {
			t.Errorf("Expected: %d, got: %d", want, got)
		}
	})
	t.Run("Given coins, no total", func(t *testing.T) {
		got := possibleDenominations([]int{}, 0)
		want := 0

		if got != want {
			t.Errorf("Expected: %d, got: %d", want, got)
		}
	})
	t.Run("Given coins, that can't make total", func(t *testing.T) {
		got := possibleDenominations([]int{5}, 2)
		want := 0

		if got != want {
			t.Errorf("Expected: %d, got: %d", want, got)
		}
	})
	t.Run("Given coins, one total", func(t *testing.T) {
		got := possibleDenominations([]int{5, 6, 9}, 5)
		want := 1

		if got != want {
			t.Errorf("Expected: %d, got: %d", want, got)
		}
	})
	t.Run("Given coins, give total", func(t *testing.T) {
		got := possibleDenominations([]int{1, 2, 3}, 4)
		want := 4

		if got != want {
			t.Errorf("Expected: %d, got: %d", want, got)
		}
	})
}
