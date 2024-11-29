package main

import (
	"fmt"
	"strings"
	"testing"
)

var (
	cases = []struct {
		Number int
		Words  string
	}{
		{Number: 0, Words: "zero"},
		{Number: 1, Words: "one"},
		{Number: 2, Words: "two"},
		{Number: 3, Words: "three"},
		{Number: 4, Words: "four"},
		{Number: 5, Words: "five"},
		{Number: 6, Words: "six"},
		{Number: 7, Words: "seven"},
		{Number: 8, Words: "eight"},
		{Number: 9, Words: "nine"},
		{Number: 10, Words: "ten"},
		{Number: -10, Words: "negative ten"},
		{Number: 11, Words: "eleven"},
		{Number: 12, Words: "twelve"},
		{Number: 13, Words: "thirteen"},
		{Number: 14, Words: "fourteen"},
		{Number: 15, Words: "fifteen"},
		{Number: 16, Words: "sixteen"},
		{Number: 17, Words: "seventeen"},
		{Number: 18, Words: "eightteen"},
		{Number: 19, Words: "nineteen"},
		{Number: 20, Words: "twenty"},
		{Number: 25, Words: "twenty five"},
		{Number: -36, Words: "negative thirty six"},
		{Number: 49, Words: "fourty nine"},
		{Number: 51, Words: "fifty one"},
		{Number: 63, Words: "sixty three"},
		{Number: 77, Words: "seventy seven"},
		{Number: 84, Words: "eighty four"},
		{Number: 92, Words: "ninety two"},
		{Number: 100, Words: "one hundred"},
		{Number: 324, Words: "three hundred twenty four"},
		{Number: 641, Words: "six hundred fourty one"},
		{Number: 1071, Words: "one thousand seventy one"},
		{Number: 2413, Words: "two thousand four hundred thirteen"},
		{Number: 17356, Words: "Numbers this large are currently not supported"},
	}
)

func TestConvertNumToWords(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Number, test.Words), func(t *testing.T) {
			var result strings.Builder

			got := convertToWords(test.Number, &result)

			if got != test.Words {
				t.Errorf("got %q, want %q", got, test.Words)
			}
		})
	}
}

func TestCharacterCount(t *testing.T) {
	t.Run("Empty String", func(t *testing.T) {
		got := characterCount("")
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
