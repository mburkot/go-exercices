package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %s", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)

	d := newDeck()
	d.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards from loaded file, got %d", len(loadedDeck))
	}

	os.Remove(testFile)
}

func TestDeal(t *testing.T) {
	d := newDeck()

	hand, remainingDeck := deal(d, 5)

	if len(remainingDeck) != len(d)-5 {
		t.Errorf("Expected %d cards from remaining deck, got %d", len(d)-5, len(remainingDeck))
	}

	if len(hand) != 5 {
		t.Errorf("Expected 5 cards from hand, got %d", len(hand))
	}
}
