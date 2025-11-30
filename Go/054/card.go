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
		return Two
	case "3":
		return Three
	case "4":
		return Four
	case "5":
		return Five
	case "6":
		return Six
	case "7":
		return Seven
	case "8":
		return Eight
	case "9":
		return Nine
	case "T":
		return Ten
	case "J":
		return Jack
	case "Q":
		return Queen
	case "K":
		return King
	case "A":
		return Ace
	default:
		return Rank(UndefinedSuite)
	}
}

func convertToSuite(s string) Suite {
	switch s {
	case "S":
		return Spades
	case "H":
		return Hearts
	case "D":
		return Diamonds
	case "C":
		return Clubs
	default:
		return UndefinedSuite
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
