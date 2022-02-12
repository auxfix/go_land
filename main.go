package main

import (
	"flag"
	"os"
	"github.com/auxfix/go_land/evenOdd"
	"github.com/auxfix/go_land/deckPack"
)

func main() {
	os.Mkdir("working_files", 0666)

	readOrWriteFlagPtr := flag.String("way", "write", "read/write")
	fileNamePtr := flag.String("file", "deck_file", "filename")
	shflPtr := flag.Bool("shfl", false, "true/false")
	evenOddPtr := flag.Bool("eod", false, "true/false")
	flag.Parse()

	if(*evenOddPtr) {
		evenOdd.CreateAndPrintEvenOdd()
		os.Exit(0)
	}

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
