package deckPack

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
	"time"
)

type deck []string

func NewDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) ToString() string {

	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func ReadDeckFromDisk(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs),",");
	d := deck(s)	
	return d
}


func (d deck) Shuffle() deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		swapI := r.Intn(len(d) -1)
		d[i], d[swapI] = d[swapI], d[i]
	}

	return d
} 
