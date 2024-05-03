package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

// new deck type extends slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (d deck) is a receiver on a function
// any variable of type 'deck' now gets access to the 'print' method
// every variable of type 'deck' can call the method on itself
// go doesn't use 'this' or 'self' just 'what it is' (d in this scenario)
// the receiver is generally referred to it in a single char or two (the first letter - it's a convention, not a rule)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func (d deck) newDeckFromFile(fileName string) (deck, error) {
	bs, err := os.Open(fileName)
	if err != nil {
		// Opt 1: log the error and return a call to newDeck()
		// Opt 2: log the error, quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
