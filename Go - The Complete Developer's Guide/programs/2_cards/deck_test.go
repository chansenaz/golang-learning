package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Two of Clubs" {
		t.Errorf("Expected first card to be Two of Clubs, but got %s", d[0])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fn := "_deckTesting"
	os.Remove(fn)

	d := newDeck()
	d.saveToFile(fn)

	loadedDeck := newDeckFromFile(fn)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}
}
