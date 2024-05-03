package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

// new deck type extends slice of strings
type deck []string

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
