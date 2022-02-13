package deckPack

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expLeng := 16
	dck := NewDeck()

	if len(dck) != expLeng {
		t.Errorf("Expected length of deck %v but got %v", expLeng, len(dck))
	}

	if dck[0].GetStr() != "One of Spades" {
		t.Errorf("Expected first card be One o Spades but got %v", dck[0].GetStr())
	}

	if dck[len(dck) - 1].GetStr() != "Four of Clubs" {
		t.Errorf("Expected last card be Four of Clubs but got %v", dck[len(dck) - 1].GetStr())
	}
}

func TestSaveDeckToDriveAndReadFromIt(t *testing.T) {
	tfn := "_deckfile"

	os.Remove(tfn)

	dck := NewDeck()
	dck.SaveToFile(tfn)
	nDck := ReadDeckFromDisk(tfn)

	if len(nDck) != 16 {
		t.Errorf("Readed deck length expected 16 but got %v", len(nDck))
	}

	os.Remove(tfn)
}