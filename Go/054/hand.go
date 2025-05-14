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

func (h Hand) HasStraight() bool {
	h.Sort()

	lowestRank := h.cards[0].rank

	for i, card := range h.cards {
		if card.rank != Rank(i)+lowestRank {
			return false
		}
	}

	return true
}

func (h Hand) StraightRank() Rank {
	h.Sort()

	lowestRank := h.cards[0].rank

	return lowestRank
}

func (h Hand) HasFlush() bool {
	suite := h.cards[0].suite

	for _, card := range h.cards {
		if card.suite != suite {
			return false
		}
	}

	return true
}

func (h Hand) HasStraightFlush() bool {
	if !h.HasFlush() {
		return false
	}

	if !h.HasStraight() {
		return false
	}

	return true
}

func (h Hand) StraightFlushRank() Rank {
	return h.StraightRank()
}

func (h Hand) HasRoyalFlush() bool {
	if !h.HasFlush() {
		return false
	}

	h.Sort()

	for i, card := range h.cards {
		if card.rank != Rank(i+9) {
			return false
		}
	}

	return true
}

func (h Hand) HasFourOfAKind() bool {
	h.Sort()

	i := 0
	kind := h.cards[0].rank
	if h.cards[0].rank != h.cards[1].rank {
		i = 1
		kind = h.cards[1].rank
	}

	for j := range 4 {
		if h.cards[j+i].rank != kind {
			return false
		}
	}

	return true
}

func (h Hand) FourOfAKindRank() Rank {
	h.Sort()

	kind := h.cards[0].rank
	if h.cards[0].rank != h.cards[1].rank {
		kind = h.cards[1].rank
	}

	return kind
}

func (h Hand) HasThreeOfAKind() bool {
	h.Sort()

	for i := range len(h.cards) - 2 {
		kind := h.cards[i].rank
		if kind == h.cards[i+1].rank && kind == h.cards[i+2].rank {
			return true
		}
	}

	return false
}

func (h Hand) ThreeOfAKindRank() Rank {
	h.Sort()

	for i := range len(h.cards) - 2 {
		kind := h.cards[i].rank
		if kind == h.cards[i+1].rank && kind == h.cards[i+2].rank {
			return kind
		}
	}

	return UndefinedRank
}

func (h Hand) HasPair() bool {
	h.Sort()

	for i := range len(h.cards) - 1 {
		kind := h.cards[i].rank
		if kind == h.cards[i+1].rank {
			return true
		}
	}

	return false
}

func (h Hand) PairRank() Rank {
	for i := range len(h.cards) - 1 {
		pair := h.cards[i].rank
		if pair == h.cards[i+1].rank {
			return pair
		}
	}

	return UndefinedRank
}

func (h Hand) HasTwoPair() bool {
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
				return true
			}
		}
	}

	return false
}

func (h Hand) TwoPairRank() Rank {
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
					return pair
				} else {
					return pair2
				}
			}
		}
	}

	return UndefinedRank
}

func (h Hand) HasFullHouse() bool {
	h.Sort()

	if h.cards[0].rank == h.cards[1].rank && h.cards[0].rank == h.cards[2].rank {
		if h.cards[3].rank == h.cards[4].rank {
			return true
		}
	}

	if h.cards[2].rank == h.cards[3].rank && h.cards[2].rank == h.cards[4].rank {
		if h.cards[0].rank == h.cards[1].rank {
			return true
		}
	}

	return false
}

func (h Hand) HasFullHouseRank() Rank {
	h.Sort()

	if h.cards[0].rank == h.cards[1].rank && h.cards[0].rank == h.cards[2].rank {
		if h.cards[3].rank == h.cards[4].rank {
			return h.cards[0].rank
		}
	}

	if h.cards[2].rank == h.cards[3].rank && h.cards[2].rank == h.cards[4].rank {
		if h.cards[0].rank == h.cards[1].rank {
			return h.cards[2].rank
		}
	}

	return UndefinedRank
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

	if h.HasPair() {
		p = OnePair
		r = h.PairRank()
	}
	if h.HasTwoPair() {
		p = TwoPairs
		r = h.TwoPairRank()
	}
	if h.HasThreeOfAKind() {
		p = ThreeOfAKind
		r = h.ThreeOfAKindRank()
	}
	if h.HasStraight() {
		p = Straight
		r = h.StraightRank()
	}
	if h.HasFlush() {
		p = Flush
		r = h.cards[h.HighestCard()].rank
	}
	if h.HasFullHouse() {
		p = FullHouse
		r = h.HasFullHouseRank()
	}
	if h.HasFourOfAKind() {
		p = FourOfAKind
		r = h.FourOfAKindRank()
	}
	if h.HasStraightFlush() {
		p = StraightFlush
		r = h.StraightFlushRank()
	}
	if h.HasRoyalFlush() {
		p = RoyalFlush
		r = Ace
	}

	return p, r
}
