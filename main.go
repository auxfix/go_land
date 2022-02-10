package main

import (
	"github.com/auxfix/go_land/deckPack"
	"flag"
)

func main() {
	readOrWriteFlagPtr := flag.String("way", "write", "read/write")
	fileNamePtr := flag.String("file", "deck_file", "filename")
	flag.Parse()
	if *readOrWriteFlagPtr == "write" {
		dck := deckPack.NewDeck()
		dck.SaveToFile(*fileNamePtr)
	} else {
		dck := deckPack.ReadDeckFromDisk(*fileNamePtr)
		dck.Print()
	}

}
