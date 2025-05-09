package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHand(t *testing.T) {
	h := ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"})
	assert.NotNil(t, h)

	t.Run("Print", func(t *testing.T) {
		assert.Equal(t,
			"Jack of Spades, Queen of Spades, Ten of Spades, King of Spades, Ace of Spades",
			h.Print())
	})

	t.Run("HighestCard", func(t *testing.T) {
		assert.Equal(t, 4, h.HighestCard())

		h2 := ConvertToHand([5]string{"JS", "AD", "TS", "AS", "AC"})
		assert.Equal(t, 3, h2.HighestCard())
	})

	t.Run("LowestCard", func(t *testing.T) {
		assert.Equal(t, 2, h.LowestCard())

		h2 := ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"})
		assert.Equal(t, 3, h2.LowestCard())
	})

	t.Run("Sort", func(t *testing.T) {})

	t.Run("IsFlush", func(t *testing.T) {
		assert.True(t, h.IsFlush(), "expected hand to be flush")

		h2 := h
		h2.cards[2].suite++
		assert.False(t, h2.IsFlush(), "expected hand not to be flush")
	})

	t.Run("IsStraightFlush", func(t *testing.T) {
		assert.True(t, h.IsStraightFlush(), "expected hand to be straight flush")

		h2 := h
		h2.cards[2].suite++
		assert.False(t, h2.IsStraightFlush(), "expected hand to not be straight flush")

		h3 := ConvertToHand([5]string{"1D", "3D", "5D", "4D", "2D"})
		assert.True(t, h3.IsStraightFlush(), "expected hand to be straight flush")

		// h4 := ConvertToHand([5]string{"1D", "6D", "5D", "4D", "2D"})
		// assert.False(t, h4.IsStraightFlush(), "expected hand to not be straight flush")
	})
}
