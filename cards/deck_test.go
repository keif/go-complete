package main

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck, _ := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("NewDeck returned wrong number of cards, expected 52 got %d", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, got %s", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Clubs" {
		t.Errorf("Expected first card of Ace of Spades, got %s", loadedDeck[len(loadedDeck)-1])
	}
}
