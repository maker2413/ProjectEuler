package main

type PokerHand int

const (
	UndefinedPokerHand PokerHand = iota
	HighCard
	OnePair
	TwoPairs
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

var pokerHandName = map[PokerHand]string{
	UndefinedPokerHand: "Undefined Hand",
	HighCard:           "Highest value card",
	OnePair:            "One pair",
	TwoPairs:           "Two pairs",
	ThreeOfAKind:       "Three of a kind",
	Straight:           "Straight",
	Flush:              "Flush",
	FullHouse:          "Full house",
	FourOfAKind:        "Four of a kind",
	StraightFlush:      "Straight Flush",
	RoyalFlush:         "Royal Flush",
}

type PokerGame struct {
	p1              Hand
	p1PokerHand     PokerHand
	p1PokerHandRank Rank
	p2              Hand
	p2PokerHand     PokerHand
	p2PokerHandRank Rank
	Winner          int
}

func (pg *PokerGame) CheckWinner() error {
	pg.p1PokerHand, pg.p1PokerHandRank = pg.p1.WinningHand()
	pg.p2PokerHand, pg.p2PokerHandRank = pg.p2.WinningHand()

	if pg.p1PokerHand > pg.p2PokerHand {
		pg.Winner = 1
	} else if pg.p2PokerHand > pg.p1PokerHand {
		pg.Winner = 2
	} else {
		if pg.p1PokerHandRank > pg.p2PokerHandRank {
			pg.Winner = 1
		} else if pg.p1PokerHandRank < pg.p2PokerHandRank {
			pg.Winner = 2
		} else {
			pg.p1PokerHandRank = pg.p1.cards[pg.p1.HighestCard()].rank
			pg.p2PokerHandRank = pg.p2.cards[pg.p2.HighestCard()].rank

			if pg.p1PokerHandRank > pg.p2PokerHandRank {
				pg.Winner = 1
			} else if pg.p2PokerHandRank > pg.p1PokerHandRank {
				pg.Winner = 2
			} else {
				pg.p1PokerHandRank = pg.p1.cards[pg.p1.NextHighestCard(pg.p1PokerHandRank)].rank
				pg.p2PokerHandRank = pg.p2.cards[pg.p2.NextHighestCard(pg.p2PokerHandRank)].rank

				if pg.p1PokerHandRank > pg.p2PokerHandRank {
					pg.Winner = 1
				} else if pg.p2PokerHandRank > pg.p1PokerHandRank {
					pg.Winner = 2
				} else {
					pg.p1PokerHandRank = pg.p1.cards[pg.p1.NextHighestCard(pg.p1PokerHandRank)].rank
					pg.p2PokerHandRank = pg.p2.cards[pg.p2.NextHighestCard(pg.p2PokerHandRank)].rank

					if pg.p1PokerHandRank > pg.p2PokerHandRank {
						pg.Winner = 1
					} else if pg.p2PokerHandRank > pg.p1PokerHandRank {
						pg.Winner = 2
					}
				}
			}
		}
	}

	return nil
}
