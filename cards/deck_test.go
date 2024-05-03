package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("NewDeck returned wrong number of cards, expected 52 got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, got %s", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of Ace of Spades, got %s", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {}
