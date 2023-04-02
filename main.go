import "testing"


package main


func main() { // Has to be first matching package main
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Two of Clubs" {
		t.Errorf("Expected last card of Two of Clubs, but got %v", d[len(d)-1])
	}
}

