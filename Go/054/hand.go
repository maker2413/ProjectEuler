package main

import "fmt"

type Hand struct {
	cards [5]Card
}

func ConvertToHand(input [5]string) Hand {
	newHand := Hand{}

	for i, c := range input {
		newHand.cards[i] = ConvertToCard(c)
	}

	return newHand
}

func (h Hand) Print() {
	fmt.Print("This hand has:")

	for _, c := range h.cards {
		fmt.Print(" ")
		c.Print()
		fmt.Print(",")
	}

	fmt.Print("\n")
}

func (h Hand) IsFlush() bool {
	suite := h.cards[0].suite

	for _, card := range h.cards {
		if card.suite != suite {
			return false
		}
	}

	return true
}

func (h Hand) IsStraightFlush() bool {
	if !h.IsFlush() {
		return false
	}

	h.Sort()

	return true
}

func (h Hand) IsRoyalFlush() bool {
	suite := h.cards[0].suite

	for i, card := range h.cards {
		if card.rank != Rank(i+10) || card.suite != suite {
			return false
		}
	}

	return true
}

func (h Hand) Sort() {}
