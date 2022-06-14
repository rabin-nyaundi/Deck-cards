package main

func main() {

	cards := deck{"card 1", newCard()}

	cards = append(cards, "another card")

	cards.print()

}

func newCard() string {
	return "The returned string"
}
