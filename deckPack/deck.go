package deckPack

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"math/rand"
	"time"
)

type deck []card

type card struct {
	suit string
	value string
}

func (crd card) GetStr() string {
	return  crd.value + " of " + crd.suit
}

func NewDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, card{suit, value})
		}
	}

	return cards
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Printf("%v %+vi\n", i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) ToString() string {
	var s []string
	for _, crd := range d {
		s = append(s, crd.value + " of " + crd.suit )
	}
	return strings.Join(s , ",")
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
	var dk deck
	for _ , st := range s {
		tempStr := strings.Split(st, " of ")
		dk  = append(dk, card{tempStr[0], tempStr[1]})
	}

	return dk
}


func (d deck) Shuffle() deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		swapI := r.Intn(len(d) -1)
		d[i], d[swapI] = d[swapI], d[i]
	}

	return d
} 
