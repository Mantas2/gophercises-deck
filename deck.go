//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
)

type Suit uint8

const (
	Diamond Suit = iota
	Club
	Heart
	Spade
	Joker
)

type Rank uint8

const (
	_ Rank = iota
	Ace
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
)

var suits = []Suit{Diamond, Club, Heart, Joker, Spade}
var ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	var deck []Card
	for _, s := range suits {
		for _, r := range ranks {
			card := Card{Suit: s, Rank: r}
			deck = append(deck, card)
		}
	}
	return deck
}
