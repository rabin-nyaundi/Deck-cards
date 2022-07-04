package main

import "fmt"

var marks map[string]int

func main() {
	fmt.Println("We learn about maps in Go lang today.")

	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	colors["green"] = "#4bf745"

	PrintColors(colors, "blue")

	fmt.Println("\n************************")

	marks = make(map[string]int)
	marks["Data Structures and algos"] = 78
	marks["Linear Programming"] = 87
	marks["Probability and statistics"] = 89
	marks["Information Systems"] = 67
	marks["automata theory"] = 94

	PrintMarks(marks, "automata theory")
}

func PrintColors(m map[string]string, colr string) {
	for color, hex := range m {
		m[colr] = "#4bc"
		fmt.Println("Hex code for", color, "is", hex)
	}
}

func PrintMarks(m map[string]int, unit1 string) {
	for unit, score := range m {
		m[unit1] = 56
		fmt.Println("Score for", unit, "is", score)
	}
}
