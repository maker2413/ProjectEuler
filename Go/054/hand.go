package main

import (
	"fmt"
)

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

// Sort Hand from lowest value to highest
func (h *Hand) Sort() {
	for i := range h.cards {
		for j := i; j < len(h.cards); j++ {
			if h.cards[j].IsLower(h.cards[i]) {
				temp := h.cards[i]
				h.cards[i] = h.cards[j]
				h.cards[j] = temp
			}
		}
	}
}

// Reverse sort Hand from highest value to lowest
func (h *Hand) RevSort() {
	for i := range h.cards {
		for j := i; j < len(h.cards); j++ {
			if h.cards[j].IsHigher(h.cards[i]) {
				temp := h.cards[i]
				h.cards[i] = h.cards[j]
				h.cards[j] = temp
			}
		}
	}
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

func (h Hand) NextHighestCard(current Rank) (position int) {
	highest := h.cards[position].rank

	for i, c := range h.cards {
		if c.rank > highest && c.rank != current {
			position = i
			highest = c.rank
		}
	}

	return position
}

func (h Hand) HasStraight() (bool, Rank) {
	h.Sort()

	lowestRank := h.cards[0].rank

	for i, card := range h.cards {
		if card.rank != Rank(i)+lowestRank {
			return false, UndefinedRank
		}
	}

	return true, lowestRank
}

func (h Hand) HasFlush() (bool, Rank) {
	suite := h.cards[0].suite

	for _, card := range h.cards {
		if card.suite != suite {
			return false, UndefinedRank
		}
	}

	return true, h.cards[h.HighestCard()].rank
}

func (h Hand) HasStraightFlush() (bool, Rank) {
	hasFlush, _ := h.HasFlush()
	if !hasFlush {
		return false, UndefinedRank
	}

	hasStraight, r := h.HasStraight()
	if !hasStraight {
		return false, UndefinedRank
	}

	return true, r
}

func (h Hand) HasRoyalFlush() (bool, Rank) {
	hasFlush, _ := h.HasFlush()
	if !hasFlush {
		return false, UndefinedRank
	}

	h.Sort()

	for i, card := range h.cards {
		if card.rank != Rank(i+9) {
			return false, UndefinedRank
		}
	}

	return true, Rank(h.cards[0].suite)
}

func (h Hand) HasFourOfAKind() (bool, Rank) {
	h.Sort()

	i := 0
	kind := h.cards[0].rank
	if h.cards[0].rank != h.cards[1].rank {
		i = 1
		kind = h.cards[1].rank
	}

	for j := range 4 {
		if h.cards[j+i].rank != kind {
			return false, UndefinedRank
		}
	}

	return true, kind
}

func (h Hand) HasThreeOfAKind() (bool, Rank) {
	h.Sort()

	for i := range len(h.cards) - 2 {
		kind := h.cards[i].rank
		if kind == h.cards[i+1].rank && kind == h.cards[i+2].rank {
			return true, kind
		}
	}

	return false, UndefinedRank
}

func (h Hand) HasPair() (bool, Rank) {
	h.Sort()

	for i := range len(h.cards) - 1 {
		kind := h.cards[i].rank
		if kind == h.cards[i+1].rank {
			return true, kind
		}
	}

	return false, UndefinedRank
}

func (h Hand) HasTwoPair() (bool, Rank) {
	h.Sort()

	pair := Rank(0)

	for i := range len(h.cards) - 1 {
		kind := h.cards[i].rank
		if kind == h.cards[i+1].rank {
			pair = kind

			break
		}
	}

	for j := range len(h.cards) - 1 {
		if pair != Rank(0) {
			pair2 := h.cards[j].rank
			if pair2 == h.cards[j+1].rank && pair2 != pair {
				if pair > pair2 {
					return true, pair
				} else {
					return true, pair2
				}
			}
		}
	}

	return false, UndefinedRank
}

func (h Hand) HasFullHouse() (bool, Rank) {
	h.Sort()

	if h.cards[0].rank == h.cards[1].rank && h.cards[0].rank == h.cards[2].rank {
		if h.cards[3].rank == h.cards[4].rank {
			return true, h.cards[0].rank
		}
	}

	if h.cards[2].rank == h.cards[3].rank && h.cards[2].rank == h.cards[4].rank {
		if h.cards[0].rank == h.cards[1].rank {
			return true, h.cards[2].rank
		}
	}

	return false, UndefinedRank
}

func (h Hand) HasHighestCard(o Hand) bool {
	h.RevSort()
	o.RevSort()

	for i := range len(h.cards) {
		if h.cards[i].rank > o.cards[i].rank {
			return true
		}
	}

	return false
}

func (h Hand) WinningHand() (PokerHand, Rank) {
	p := UndefinedPokerHand
	r := UndefinedRank

	hasPair, pairR := h.HasPair()
	if hasPair {
		p = OnePair
		r = pairR
	}
	hasTwoPair, twoPairR := h.HasTwoPair()
	if hasTwoPair {
		p = TwoPairs
		r = twoPairR
	}
	hasThreeOfAKind, threeR := h.HasThreeOfAKind()
	if hasThreeOfAKind {
		p = ThreeOfAKind
		r = threeR
	}
	hasStraight, straightR := h.HasStraight()
	if hasStraight {
		p = Straight
		r = straightR
	}
	hasFlush, flushR := h.HasFlush()
	if hasFlush {
		p = Flush
		r = flushR
	}
	hasFullHouse, fullHouseR := h.HasFullHouse()
	if hasFullHouse {
		p = FullHouse
		r = fullHouseR
	}
	hasFourOfAKind, fourR := h.HasFourOfAKind()
	if hasFourOfAKind {
		p = FourOfAKind
		r = fourR
	}
	hasStraightFlush, straightFlushR := h.HasStraightFlush()
	if hasStraightFlush {
		p = StraightFlush
		r = straightFlushR
	}
	hasRoyalFlush, royalR := h.HasRoyalFlush()
	if hasRoyalFlush {
		p = RoyalFlush
		r = royalR
	}

	return p, r
}
