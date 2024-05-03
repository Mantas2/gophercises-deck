package deck

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})

	//Output:
	//Ace of Hearts
}

func TestNew(t *testing.T) {
	want := []Card{
		{Suit: Diamond, Rank: Ace},
		{Suit: Diamond, Rank: Two}, {Suit: Diamond, Rank: Three}, {Suit: Diamond, Rank: Four}, {Suit: Diamond, Rank: Five}, {Suit: Diamond, Rank: Six}, {Suit: Diamond, Rank: Seven}, {Suit: Diamond, Rank: Eight}, {Suit: Diamond, Rank: Nine}, {Suit: Diamond, Rank: Ten}, {Suit: Diamond, Rank: Jack}, {Suit: Diamond, Rank: Queen}, {Suit: Diamond, Rank: King},
		{Suit: Club, Rank: Ace}, {Suit: Club, Rank: Two}, {Suit: Club, Rank: Three}, {Suit: Club, Rank: Four}, {Suit: Club, Rank: Five}, {Suit: Club, Rank: Six}, {Suit: Club, Rank: Seven}, {Suit: Club, Rank: Eight}, {Suit: Club, Rank: Nine}, {Suit: Club, Rank: Ten}, {Suit: Club, Rank: Jack}, {Suit: Club, Rank: Queen}, {Suit: Club, Rank: King},
		{Suit: Heart, Rank: Ace}, {Suit: Heart, Rank: Two}, {Suit: Heart, Rank: Three}, {Suit: Heart, Rank: Four}, {Suit: Heart, Rank: Five}, {Suit: Heart, Rank: Six}, {Suit: Heart, Rank: Seven}, {Suit: Heart, Rank: Eight}, {Suit: Heart, Rank: Nine}, {Suit: Heart, Rank: Ten}, {Suit: Heart, Rank: Jack}, {Suit: Heart, Rank: Queen}, {Suit: Heart, Rank: King},
		{Suit: Spade, Rank: Ace}, {Suit: Spade, Rank: Two}, {Suit: Spade, Rank: Three}, {Suit: Spade, Rank: Four}, {Suit: Spade, Rank: Five}, {Suit: Spade, Rank: Six}, {Suit: Spade, Rank: Seven}, {Suit: Spade, Rank: Eight}, {Suit: Spade, Rank: Nine}, {Suit: Spade, Rank: Ten}, {Suit: Spade, Rank: Jack}, {Suit: Spade, Rank: Queen}, {Suit: Spade, Rank: King},
	}
	t.Run("idk man", func(t *testing.T) {
		got := New()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s; want %s", got, want)
		}
	})
	t.Run("idk sorted", func(t *testing.T) {
		cards := New(DefaultSort)
		exp := Card{Rank: Ace, Suit: Diamond}
		if cards[0] != exp {
			t.Error("Expected Ace of Diamonds first. Received:", cards[0])
		}
	})

	// t.Run("idk shuffled", func(t *testing.T) {
	// 	cards := New(Shuffle)
	// 	exp := Card{Rank: Two, Suit: Diamond}
	// 	if cards[0] != exp {
	// 		t.Error("Expected to be random each time. Received:", cards)
	// 	}
	// })

	t.Run("idk jokers", func(t *testing.T) {
		cards := New(Jokers(7))
		count := 0
		for _, c := range cards {
			if c.Suit == Joker {
				count++
			}
		}
		if count != 7 {
			t.Error("Expected 7, Received:", cards)
		}
	})

	t.Run("idk filter", func(t *testing.T) {
		wantFiltered := []Card{
			{Suit: Diamond, Rank: Ace},
			{Suit: Diamond}, {Suit: Diamond}, {Suit: Diamond, Rank: Four}, {Suit: Diamond, Rank: Five}, {Suit: Diamond, Rank: Six}, {Suit: Diamond, Rank: Seven}, {Suit: Diamond, Rank: Eight}, {Suit: Diamond, Rank: Nine}, {Suit: Diamond, Rank: Ten}, {Suit: Diamond, Rank: Jack}, {Suit: Diamond, Rank: Queen}, {Suit: Diamond, Rank: King},
			{Suit: Club, Rank: Ace}, {Suit: Club}, {Suit: Club}, {Suit: Club, Rank: Four}, {Suit: Club, Rank: Five}, {Suit: Club, Rank: Six}, {Suit: Club, Rank: Seven}, {Suit: Club, Rank: Eight}, {Suit: Club, Rank: Nine}, {Suit: Club, Rank: Ten}, {Suit: Club, Rank: Jack}, {Suit: Club, Rank: Queen}, {Suit: Club, Rank: King},
			{Suit: Heart, Rank: Ace}, {Suit: Heart}, {Suit: Heart}, {Suit: Heart, Rank: Four}, {Suit: Heart, Rank: Five}, {Suit: Heart, Rank: Six}, {Suit: Heart, Rank: Seven}, {Suit: Heart, Rank: Eight}, {Suit: Heart, Rank: Nine}, {Suit: Heart, Rank: Ten}, {Suit: Heart, Rank: Jack}, {Suit: Heart, Rank: Queen}, {Suit: Heart, Rank: King},
			{Suit: Spade, Rank: Ace}, {Suit: Spade}, {Suit: Spade}, {Suit: Spade, Rank: Four}, {Suit: Spade, Rank: Five}, {Suit: Spade, Rank: Six}, {Suit: Spade, Rank: Seven}, {Suit: Spade, Rank: Eight}, {Suit: Spade, Rank: Nine}, {Suit: Spade, Rank: Ten}, {Suit: Spade, Rank: Jack}, {Suit: Spade, Rank: Queen}, {Suit: Spade, Rank: King},
		}
		// got := New(Filter(Card{Rank: Three}))
		filter := func(card Card) bool {
			return card.Rank == Two || card.Rank == Three
		}
		got := New(Filter(filter))
		for _, c := range got {
			if c.Rank == Two || c.Rank == Three {
				t.Errorf("got %s; want: %s", got, wantFiltered)
			}
		}
	})

	t.Run("idk multiply", func(t *testing.T) {
		cards := New(Multiply(1))
		if len(cards) != 52 {
			t.Error("Expected 1 set, Received:", len(cards))
		}
	})
}
