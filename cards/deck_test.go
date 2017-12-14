package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length to be 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}
	if d[len(d)-1] != "Queen of Clubs" {
		t.Errorf("Expected last card to be Queen of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadeddeck := newDeckFromFile("_decktesting")

	if len(loadeddeck) != 52 {
		t.Errorf("Expected 52 card in deck, but got %v", len(loadeddeck))
	}
	os.Remove("_decktesting")
}
