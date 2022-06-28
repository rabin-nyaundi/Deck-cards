package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue + " of "+cardSuit )
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	result := strings.Join([]string(d), ",")
	return result
}

func (d deck) saveToFile(filename string) error {

	// convert deck to string
	// convert string[] to bytes

	err := ioutil.WriteFile(filename, []byte(d.toString()), 0644)

	if err != nil {
		return err
	}
	return nil
}

func newDeckFromFile(filename string) deck {
	byte, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	stringBytes := strings.Split(string(byte), ",")

	d := deck(stringBytes)
	return d
}

func (d deck) shuffle() {
	// define the source for which random numbers will be created
	source := rand.NewSource(time.Now().UnixNano())

	// create random numbers
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn((len(d) - 1))
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
