package main

import "fmt"

type Rank int

type Suite int

const (
	UndefinedRank Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	UndefinedSuite Suite = iota
	Clubs
	Diamonds
	Hearts
	Spades
)

var rankName = map[Rank]string{
	UndefinedRank: "Undefined",
	Two:           "Two",
	Three:         "Three",
	Four:          "Four",
	Five:          "Five",
	Six:           "Six",
	Seven:         "Seven",
	Eight:         "Eight",
	Nine:          "Nine",
	Ten:           "Ten",
	Jack:          "Jack",
	Queen:         "Queen",
	King:          "King",
	Ace:           "Ace",
}

var suiteName = map[Suite]string{
	UndefinedSuite: "Undefined",
	Clubs:          "Clubs",
	Diamonds:       "Diamonds",
	Hearts:         "Hearts",
	Spades:         "Spades",
}

type Card struct {
	rank  Rank
	suite Suite
}

func ConvertToCard(input string) Card {
	newCard := Card{
		rank:  convertToRank(string(input[0])),
		suite: convertToSuite(string(input[1])),
	}

	return newCard
}

func convertToRank(r string) Rank {
	switch r {
	case "2":
		return Rank(1)
	case "3":
		return Rank(2)
	case "4":
		return Rank(3)
	case "5":
		return Rank(4)
	case "6":
		return Rank(5)
	case "7":
		return Rank(6)
	case "8":
		return Rank(7)
	case "9":
		return Rank(8)
	case "T":
		return Rank(9)
	case "J":
		return Rank(10)
	case "Q":
		return Rank(11)
	case "K":
		return Rank(12)
	case "A":
		return Rank(13)
	default:
		return Rank(0)
	}
}

func convertToSuite(s string) Suite {
	switch s {
	case "S":
		return Suite(4)
	case "H":
		return Suite(3)
	case "D":
		return Suite(2)
	case "C":
		return Suite(1)
	default:
		return Suite(0)
	}
}

func (c Card) Print() string {
	return fmt.Sprintf("%s of %s", rankName[c.rank], suiteName[c.suite])
}

func (c Card) IsLower(c2 Card) bool {
	if c.rank < c2.rank {
		return true
	} else if c.rank == c2.rank && c.suite < c2.suite {
		return true
	}

	return false
}

func (c Card) IsHigher(c2 Card) bool {
	if c.rank > c2.rank {
		return true
	} else if c.rank == c2.rank && c.suite > c2.suite {
		return true
	}

	return false
}
