//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand"
	"sort"
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

func New(opts ...func([]Card) []Card) []Card {
	var deck []Card
	for _, s := range suits {
		if s == Joker {
			continue
		}
		for _, r := range ranks {
			card := Card{Suit: s, Rank: r}
			deck = append(deck, card)
		}
	}
	for _, opt := range opts {
		deck = opt(deck)
	}
	return deck
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

func Shuffle(cards []Card) []Card {
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})
	return cards
}

func Less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		if cards[i].Suit == cards[j].Suit {
			return cards[i].Rank < cards[j].Rank
		}
		return cards[i].Suit < cards[j].Suit
	}
}

func Jokers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{Suit: Joker})
		}
		return cards
	}
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var filtered []Card
		for _, crd := range cards {
			if !f(crd) {
				filtered = append(filtered, crd)
			}
		}
		return filtered
	}
}

func Multiply(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var newDeck []Card
		for i := 0; i < n; i++ {
			newDeck = append(newDeck, cards...)
		}
		return newDeck
	}
}

// func MultiplyStupid(n int) func([]Card) []Card {
// 	return func(_ []Card) []Card {
// 		var newDeck []Card
// 		for i := 0; i < n; i++ {
// 			newDeck = append(newDeck, New()...)
// 		}
// 		return newDeck
// 	}
// }
