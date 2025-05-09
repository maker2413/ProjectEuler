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

func (h Hand) Print() string {
	s := ""

	for i, c := range h.cards {
		if i != 4 {
			s += fmt.Sprintf("%s, ", c.Print())
		} else {
			s += fmt.Sprintf("%s", c.Print())
		}
	}

	return s
}

func (h *Hand) Sort() {
	lowest := h.LowestCard()
	h2 := Hand{}

	h2.cards[0] = h.cards[lowest]
}

func (h Hand) LowestCard() (position int) {
	lowest := h.cards[0].rank

	for i, c := range h.cards {
		if c.rank < lowest {
			position = i
			lowest = c.rank
		} else if c.rank == lowest && c.suite < h.cards[position].suite {
			position = i
		}
	}

	return position
}

func (h Hand) HighestCard() (position int) {
	highest := h.cards[position].rank

	for i, c := range h.cards {
		if c.rank > highest {
			position = i
			highest = c.rank
		} else if c.rank == highest && c.suite > h.cards[position].suite {
			position = i
		}
	}

	return position
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
