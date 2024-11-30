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
		Count  int
	}{
		{Number: 0, Words: "zero", Count: 4},
		{Number: 1, Words: "one", Count: 3},
		{Number: 2, Words: "two", Count: 3},
		{Number: 3, Words: "three", Count: 5},
		{Number: 4, Words: "four", Count: 4},
		{Number: 5, Words: "five", Count: 4},
		{Number: 6, Words: "six", Count: 3},
		{Number: 7, Words: "seven", Count: 5},
		{Number: 8, Words: "eight", Count: 5},
		{Number: 9, Words: "nine", Count: 4},
		{Number: 10, Words: "ten", Count: 3},
		{Number: -10, Words: "negative ten", Count: 11},
		{Number: 11, Words: "eleven", Count: 6},
		{Number: 12, Words: "twelve", Count: 6},
		{Number: 13, Words: "thirteen", Count: 8},
		{Number: 14, Words: "fourteen", Count: 8},
		{Number: 15, Words: "fifteen", Count: 7},
		{Number: 16, Words: "sixteen", Count: 7},
		{Number: 17, Words: "seventeen", Count: 9},
		{Number: 18, Words: "eighteen", Count: 8},
		{Number: 19, Words: "nineteen", Count: 8},
		{Number: 20, Words: "twenty", Count: 6},
		{Number: 25, Words: "twenty five", Count: 10},
		{Number: 30, Words: "thirty", Count: 6},
		{Number: -36, Words: "negative thirty six", Count: 17},
		{Number: 40, Words: "forty", Count: 5},
		{Number: 49, Words: "forty nine", Count: 9},
		{Number: 50, Words: "fifty", Count: 5},
		{Number: 51, Words: "fifty one", Count: 8},
		{Number: 60, Words: "sixty", Count: 5},
		{Number: 63, Words: "sixty three", Count: 10},
		{Number: 70, Words: "seventy", Count: 7},
		{Number: 77, Words: "seventy seven", Count: 12},
		{Number: 80, Words: "eighty", Count: 6},
		{Number: 84, Words: "eighty four", Count: 10},
		{Number: 90, Words: "ninety", Count: 6},
		{Number: 92, Words: "ninety two", Count: 9},
		{Number: 100, Words: "one hundred", Count: 10},
		{Number: 115, Words: "one hundred and fifteen", Count: 20},
		{Number: 249, Words: "two hundred and forty nine", Count: 22},
		{Number: 342, Words: "three hundred and forty two", Count: 23},
		{Number: 500, Words: "five hundred", Count: 11},
		{Number: 641, Words: "six hundred and forty one", Count: 21},
		{Number: 967, Words: "nine hundred and sixty seven", Count: 24},
		{Number: 1000, Words: "one thousand", Count: 11},
		{Number: 1071, Words: "one thousand seventy one", Count: 21},
		{Number: 2413, Words: "two thousand four hundred and thirteen", Count: 33},
		{Number: 17356, Words: "Numbers this large are currently not supported", Count: 40},
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
	t.Run("Single Word", func(t *testing.T) {
		got := characterCount("Maker")
		want := 5

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("Two Words", func(t *testing.T) {
		got := characterCount("Hello World")
		want := 10

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("Hyphenated Word", func(t *testing.T) {
		got := characterCount("twenty-one")
		want := 9

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d has a character count of: %q", test.Number, test.Count), func(t *testing.T) {
			var result strings.Builder

			got := characterCount(convertToWords(test.Number, &result))

			if got != test.Count {
				t.Errorf("got %d, want %d", got, test.Count)
			}
		})
	}
}
