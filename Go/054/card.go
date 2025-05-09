package main

import "fmt"

type Rank int

type Suite int

const (
	UndefinedRank Rank = iota
	One
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
	Spades
	Hearts
	Diamonds
	Clubs
)

var rankName = map[Rank]string{
	UndefinedRank: "Undefined",
	One:           "One",
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
	Spades:         "Spades",
	Hearts:         "Hearts",
	Diamonds:       "Diamonds",
	Clubs:          "Clubs",
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
	case "1":
		return Rank(1)
	case "2":
		return Rank(2)
	case "3":
		return Rank(3)
	case "4":
		return Rank(4)
	case "5":
		return Rank(5)
	case "6":
		return Rank(6)
	case "7":
		return Rank(7)
	case "8":
		return Rank(8)
	case "9":
		return Rank(9)
	case "T":
		return Rank(10)
	case "J":
		return Rank(11)
	case "Q":
		return Rank(12)
	case "K":
		return Rank(13)
	case "A":
		return Rank(14)
	default:
		return Rank(0)
	}
}

func convertToSuite(s string) Suite {
	switch s {
	case "S":
		return Suite(1)
	case "H":
		return Suite(2)
	case "D":
		return Suite(3)
	case "C":
		return Suite(4)
	default:
		return Suite(0)
	}
}

func (c Card) Print() {
	fmt.Printf("%s of %s", rankName[c.rank], suiteName[c.suite])
}
