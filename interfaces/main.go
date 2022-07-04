package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct {
}
type SpanishBot struct {
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func (eb EnglishBot) getGreeting() string {
	return "Hello, there"
}

func (sb SpanishBot) getGreeting() string {
	return "Hola!"
}
