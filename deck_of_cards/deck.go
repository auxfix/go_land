package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for i, suit := range cardSuits {
		for j, value := range cardSuits{
			cards = append(cards, suit" of "value)
		}
	}
}


func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
