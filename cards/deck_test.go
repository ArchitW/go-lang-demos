package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected Deck 52 got %v", len(d))
	}
	if d[len(d)-1] != "Ten of Clubs" {
		t.Errorf("Not valid card got %v", d[len(d)-1])
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Card is not Ace of Sp, got %v", d[0])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	lodeDeck := newDeckFromFile("_decktesting")
	if len(lodeDeck) != 52 {
		t.Errorf("%v", len(lodeDeck))
	}
}
