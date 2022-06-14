package main

import "fmt"

func main() {

	cards := newDeck()

	halfCards := len(cards)/2 + 10

	cards.print()
	fmt.Println("============all cards=============")

	hand, remainingCards := deal(cards, halfCards)

	hand.print()
	fmt.Println("=============hand cards============")

	remainingCards.print()
	fmt.Println("=============remaining cards============")

	fmt.Println(hand.toString())
	fmt.Println("=============Strings============")

	hand.saveToFile("cards")
}
