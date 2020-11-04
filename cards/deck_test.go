package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	t.Run("deck length is 52", func(t *testing.T) {
		if len(d) != 52 {
			t.Errorf("Expected deck of length 52 got %v", len(d))
		}
	})

	t.Run("first card is Ace of Spades", func(t *testing.T) {
		if d[0] != "Ace of Spades" {
			t.Errorf("Expected first card to be \"Ace of Spades\" but go %v", d[0])
		}
	})

	t.Run("last card is King of Clubs", func(t *testing.T) {
		if d[len(d)-1] != "King of Clubs" {
			t.Errorf("Expected last card to be \"King of Clubs\" but go %v", d[len(d)-1])
		}
	})

}

func TestSaveAndRemoveDeck(t *testing.T) {
	filename := "/tmp/_decktesting.deck"
	os.Remove(filename)
	d := newDeck()
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	t.Run("load_deck_52_length", func(t *testing.T) {
		if len(loadedDeck) != 52 {
			t.Errorf("Expected deck length of 52 but got %v", len(loadedDeck))
		}
	})
	os.Remove(filename)
}
