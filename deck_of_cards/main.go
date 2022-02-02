package main

func main() {
	cards := deck{"two", "three", "four"}

	cards = append(cards, "One")

	cards.print()
}
