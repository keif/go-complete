package main

import "fmt"

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
