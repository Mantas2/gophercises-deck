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
		{Suit: Diamond, Rank: Ace}, {Suit: Diamond, Rank: Two}, {Suit: Diamond, Rank: Three}, {Suit: Diamond, Rank: Four}, {Suit: Diamond, Rank: Five}, {Suit: Diamond, Rank: Six}, {Suit: Diamond, Rank: Seven}, {Suit: Diamond, Rank: Eight}, {Suit: Diamond, Rank: Nine}, {Suit: Diamond, Rank: Ten}, {Suit: Diamond, Rank: Jack}, {Suit: Diamond, Rank: Queen}, {Suit: Diamond, Rank: King},
		{Suit: Club, Rank: Ace}, {Suit: Club, Rank: Two}, {Suit: Club, Rank: Three}, {Suit: Club, Rank: Four}, {Suit: Club, Rank: Five}, {Suit: Club, Rank: Six}, {Suit: Club, Rank: Seven}, {Suit: Club, Rank: Eight}, {Suit: Club, Rank: Nine}, {Suit: Club, Rank: Ten}, {Suit: Club, Rank: Jack}, {Suit: Club, Rank: Queen}, {Suit: Club, Rank: King},
		{Suit: Heart, Rank: Ace}, {Suit: Heart, Rank: Two}, {Suit: Heart, Rank: Three}, {Suit: Heart, Rank: Four}, {Suit: Heart, Rank: Five}, {Suit: Heart, Rank: Six}, {Suit: Heart, Rank: Seven}, {Suit: Heart, Rank: Eight}, {Suit: Heart, Rank: Nine}, {Suit: Heart, Rank: Ten}, {Suit: Heart, Rank: Jack}, {Suit: Heart, Rank: Queen}, {Suit: Heart, Rank: King},
		{Suit: Joker, Rank: Ace}, {Suit: Joker, Rank: Two}, {Suit: Joker, Rank: Three}, {Suit: Joker, Rank: Four}, {Suit: Joker, Rank: Five}, {Suit: Joker, Rank: Six}, {Suit: Joker, Rank: Seven}, {Suit: Joker, Rank: Eight}, {Suit: Joker, Rank: Nine}, {Suit: Joker, Rank: Ten}, {Suit: Joker, Rank: Jack}, {Suit: Joker, Rank: Queen}, {Suit: Joker, Rank: King},
		{Suit: Spade, Rank: Ace}, {Suit: Spade, Rank: Two}, {Suit: Spade, Rank: Three}, {Suit: Spade, Rank: Four}, {Suit: Spade, Rank: Five}, {Suit: Spade, Rank: Six}, {Suit: Spade, Rank: Seven}, {Suit: Spade, Rank: Eight}, {Suit: Spade, Rank: Nine}, {Suit: Spade, Rank: Ten}, {Suit: Spade, Rank: Jack}, {Suit: Spade, Rank: Queen}, {Suit: Spade, Rank: King},
	}
	t.Run("idk man", func(t *testing.T) {
		got := New()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %s; want %s", got, want)
		}
	})
}
