package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("cards")
	cards, _ := newDeckFromFile("cards")
	cards.print()
}
