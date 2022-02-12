package deckPack

import "testing"

func TestNewDeck(t *testing.T) {
	expLeng := 16
	dck := NewDeck()

	if len(dck) != expLeng {
		t.Errorf("Expected length of deck %v but got %v", expLeng, len(dck))
	}
}