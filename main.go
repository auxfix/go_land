package main

import (
	"flag"
	"os"
	"github.com/auxfix/go_land/evenOdd"
	"github.com/auxfix/go_land/deckPack"
	"github.com/auxfix/go_land/bot"
)

func main() {
	os.Mkdir("working_files", 0666)

	modeFlagPtr := flag.String("mode", "evenOdd", "evenOdd/readWrite/bot",)
	readOrWriteFlagPtr := flag.String("way", "write", "read/write")
	fileNamePtr := flag.String("file", "deck_file", "filename")
	shflPtr := flag.Bool("shfl", false, "true/false")
	languagePtr := flag.String("leng", "en", "en/ru")

	flag.Parse()

	if(*modeFlagPtr == "evenOdd") {
		evenOdd.CreateAndPrintEvenOdd()
		os.Exit(0)
	} else if(*modeFlagPtr == "readWrite") {
		if *readOrWriteFlagPtr == "write" {
			dck := deckPack.NewDeck()
			if(*shflPtr) { dck = dck.Shuffle() }
			dck.SaveToFile("./working_files/" + *fileNamePtr)
		} else {
			dck := deckPack.ReadDeckFromDisk("./working_files/" + *fileNamePtr)
			if(*shflPtr) { dck = dck.Shuffle() }
			dck.Print()
		}
		os.Exit(0)
	} else if(*modeFlagPtr == "bot") {
		if(*languagePtr == "en"){
			bt := bot.EngBot{}
			bot.TypeGreetings(bt)
		} else {
			bt := bot.RuBot{}
			bot.TypeGreetings(bt)
		} 
	
		os.Exit(0);
	}



}
