package main

import (
	"flag"
	"os"

	"github.com/auxfix/go_land/deckPack"
)

func main() {
	os.Mkdir("working_files", 0666)

	readOrWriteFlagPtr := flag.String("way", "write", "read/write")
	fileNamePtr := flag.String("file", "deck_file", "filename")
	shflPtr := flag.Bool("shfl", false, "true/false")
	flag.Parse()
	if *readOrWriteFlagPtr == "write" {
		dck := deckPack.NewDeck()
		if(*shflPtr) { dck = dck.Shuffle() }
		dck.SaveToFile("./working_files/" + *fileNamePtr)
	} else {
		dck := deckPack.ReadDeckFromDisk("./working_files/" + *fileNamePtr)
		if(*shflPtr) { dck = dck.Shuffle() }
		dck.Print()
	}

}
