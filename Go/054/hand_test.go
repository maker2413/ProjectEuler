package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHand(t *testing.T) {
	h := ConvertToHand([5]string{"JS", "9C", "TH", "2S", "AC"})
	assert.NotNil(t, h)

	t.Run("Print", func(t *testing.T) {
		assert.Equal(t,
			"Jack of Spades, Nine of Clubs, Ten of Hearts, Two of Spades, Ace of Clubs",
			h.Print())
	})

	t.Run("HighestCard", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected int
		}{
			{input: ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}), expected: 4},
			{input: ConvertToHand([5]string{"JS", "AD", "TS", "AS", "AC"}), expected: 3},
		}

		for _, tt := range tests {
			assert.Equal(t, tt.expected, tt.input.HighestCard())
		}
	})

	t.Run("LowestCard", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected int
		}{
			{input: ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}), expected: 2},
			{input: ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}), expected: 3},
		}

		for _, tt := range tests {
			assert.Equal(t, tt.expected, tt.input.LowestCard())
		}
	})

	t.Run("Sort", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected Hand
		}{
			{
				input:    ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}),
				expected: ConvertToHand([5]string{"TS", "JS", "QS", "KS", "AS"}),
			},
			{
				input:    ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}),
				expected: ConvertToHand([5]string{"TC", "TS", "JS", "AC", "AD"}),
			},
		}

		for _, tt := range tests {
			tt.input.Sort()

			assert.Equal(t, tt.expected, tt.input)
		}
	})

	t.Run("RevSort", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected Hand
		}{
			{
				input:    ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}),
				expected: ConvertToHand([5]string{"AS", "KS", "QS", "JS", "TS"}),
			},
			{
				input:    ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}),
				expected: ConvertToHand([5]string{"AD", "AC", "JS", "TS", "TC"}),
			},
		}

		for _, tt := range tests {
			tt.input.RevSort()

			assert.Equal(t, tt.expected, tt.input)
		}
	})

	t.Run("HasStraight", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"5S", "4H", "7S", "6D", "3C"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasStraight()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasFlush", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasFlush()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasStraightFlush", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "QS", "TS", "KS", "9S"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "QS", "AS", "KS", "9S"}),
				expected: false,
			},
			{
				input:    ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}),
				expected: false,
			},
			{
				input:    ConvertToHand([5]string{"2S", "5S", "4S", "6S", "3S"}),
				expected: true,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasStraightFlush()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasRoyalFlush", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "QS", "TS", "KS", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}),
				expected: false,
			},
			{
				input:    ConvertToHand([5]string{"JS", "9S", "TS", "KS", "AS"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasRoyalFlush()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasFourOfAKind", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JD", "JH", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JH", "JD", "TC"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "AD", "TS", "TC", "AC"}),
				expected: false,
			},
			{
				input:    ConvertToHand([5]string{"KC", "9C", "9S", "KS", "9H"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasFourOfAKind()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasThreeOfAKind", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JD", "JH", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JH", "JD", "TC"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"7S", "8C", "TS", "TC", "TD"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JH", "TD", "TC"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"KC", "9C", "9S", "KS", "TH"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasThreeOfAKind()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasPair", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JD", "JH", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"9S", "8C", "JH", "TD", "TC"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"KC", "9C", "8S", "JS", "TH"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasPair()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasTwoPair", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JD", "AH", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"8S", "8C", "JH", "TD", "TC"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"KC", "9C", "8S", "JS", "TH"}),
				expected: false,
			},
			{
				input:    ConvertToHand([5]string{"KC", "9C", "9S", "JS", "TH"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasTwoPair()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasFullHouse", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JD", "AH", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JH", "TD", "TC"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"KC", "9C", "9S", "KS", "TH"}),
				expected: false,
			},
			{
				input:    ConvertToHand([5]string{"KC", "9C", "7S", "KS", "KH"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			result, _ := tt.input.HasFullHouse()
			assert.Equal(t, tt.expected, result, tt.input.Print())
		}
	})

	t.Run("HasHighestCard", func(t *testing.T) {
		tests := []struct {
			input    Hand
			opponent Hand
			expected bool
		}{
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JD", "AH", "AS"}),
				opponent: ConvertToHand([5]string{"JS", "TC", "JD", "AH", "AS"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "QC", "TD", "AH", "AS"}),
				opponent: ConvertToHand([5]string{"JD", "QH", "TS", "AD", "9S"}),
				expected: true,
			},
			{
				input:    ConvertToHand([5]string{"JS", "JC", "JD", "AH", "AS"}),
				opponent: ConvertToHand([5]string{"JS", "JC", "JD", "AH", "AS"}),
				expected: false,
			},
		}

		for _, tt := range tests {
			assert.Equal(t, tt.expected, tt.input.HasHighestCard(tt.opponent),
				tt.input.Print(), tt.opponent.Print())
		}
	})

	t.Run("WinningHand", func(t *testing.T) {
		tests := []struct {
			input    Hand
			expected PokerHand
		}{
			{
				input:    ConvertToHand([5]string{"JS", "QS", "KS", "TS", "AS"}),
				expected: RoyalFlush,
			},
			{
				input:    ConvertToHand([5]string{"JS", "AH", "AD", "AC", "AS"}),
				expected: FourOfAKind,
			},
			{
				input:    ConvertToHand([5]string{"JS", "JH", "JD", "AC", "AS"}),
				expected: FullHouse,
			},
			{
				input:    ConvertToHand([5]string{"5D", "8C", "9S", "JS", "AC"}),
				expected: UndefinedPokerHand,
			},
		}

		for _, tt := range tests {
			winHand, winRank := tt.input.WinningHand()
			assert.Equal(t, tt.expected, winHand, tt.input.Print(), pokerHandName[winHand])
			assert.NotNil(t, winRank)
		}
	})
}
