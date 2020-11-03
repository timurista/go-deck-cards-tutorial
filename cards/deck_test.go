package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck of length 52 got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be \"Ace of Spades\" but go %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be \"King of Clubs\" but go %v", d[len(d)-1])
	}
}
