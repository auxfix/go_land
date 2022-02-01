package main

import (
	"fmt"
	"strconv"
)

func main() {
	cards := deck{newCard(1), newCard(2), newCard(3)}

	cards = append(cards, newCard(6))

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard(ndx int) string {
	return "__Five of Diamonds(" + strconv.Itoa(ndx) + ")  "
}
