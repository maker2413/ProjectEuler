package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToRank(t *testing.T) {
	tests := []string{
		"undefined",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"T",
		"J",
		"Q",
		"K",
		"A",
	}

	for i, test := range tests {
		assert.Equal(t, Rank(i), convertToRank(test), "expected correct rank to be returned")
	}
}

func TestConvertToSuite(t *testing.T) {
	tests := []string{
		"undefined",
		"C",
		"D",
		"H",
		"S",
	}

	for i, test := range tests {
		assert.Equal(t, Suite(i), convertToSuite(test), "expected correct suit to be returned")
	}
}

func TestCard(t *testing.T) {
	c := ConvertToCard("TD")
	assert.Equal(t, Card{rank: 10, suite: 2}, c, "expected a ten of diamonds to be created")

	t.Run("Print", func(t *testing.T) {
		assert.Equal(t, "Ten of Diamonds", c.Print())
	})

	t.Run("IsLower", func(t *testing.T) {
		// TD < KS
		assert.True(t, c.IsLower(ConvertToCard("KS")))

		// TD < TS
		assert.True(t, c.IsLower(ConvertToCard("TS")))

		// TD > 2S
		assert.False(t, c.IsLower(ConvertToCard("2S")))

		// TD == TD
		assert.False(t, c.IsLower(ConvertToCard("TD")))
	})

	t.Run("IsHigher", func(t *testing.T) {
		// TD > TC
		assert.True(t, c.IsHigher(ConvertToCard("TC")))

		// TD > 2S
		assert.True(t, c.IsHigher(ConvertToCard("2S")))

		// TD < KS
		assert.False(t, c.IsHigher(ConvertToCard("KS")))

		// TD == TD
		assert.False(t, c.IsHigher(ConvertToCard("TD")))
	})
}
