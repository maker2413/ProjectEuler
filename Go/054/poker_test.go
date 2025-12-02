package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoker(t *testing.T) {
	t.Run("WinningHand", func(t *testing.T) {
		tests := []struct {
			input    PokerGame
			expected int
		}{
			{
				input: PokerGame{
					p1: ConvertToHand([5]string{"5H", "5C", "6S", "7S", "KD"}),
					p2: ConvertToHand([5]string{"2C", "3S", "8S", "8D", "TD"}),
				},
				expected: 2,
			},
			{
				input: PokerGame{
					p1: ConvertToHand([5]string{"5D", "8C", "9S", "JS", "AC"}),
					p2: ConvertToHand([5]string{"2C", "5C", "7D", "8S", "QH"}),
				},
				expected: 1,
			},
			{
				input: PokerGame{
					p1: ConvertToHand([5]string{"2D", "9C", "AS", "AH", "AC"}),
					p2: ConvertToHand([5]string{"3D", "6D", "7D", "TD", "QD"}),
				},
				expected: 2,
			},
			{
				input: PokerGame{
					p1: ConvertToHand([5]string{"4D", "6S", "9H", "QH", "QC"}),
					p2: ConvertToHand([5]string{"3D", "6D", "7H", "QD", "QS"}),
				},
				expected: 1,
			},
			{
				input: PokerGame{
					p1: ConvertToHand([5]string{"2H", "2D", "4C", "4D", "4S"}),
					p2: ConvertToHand([5]string{"3C", "3D", "3S", "9D", "9D"}),
				},
				expected: 1,
			},
		}

		for _, tt := range tests {
			err := tt.input.CheckWinner()
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, tt.input.Winner,
				tt.input.p1.Print(), tt.input.p2.Print(),
				tt.input.p1PokerHand, tt.input.p2PokerHand)
		}
	})
}
