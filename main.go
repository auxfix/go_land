package main

import (
	"github.com/auxfix/go_land/deckPack"
)

func main() {
	dck := deckPack.NewDeck()
	dck.SaveToFile("my_not_areally_first_file")
}
