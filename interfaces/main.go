package main

import "fmt"

type user struct {
	name string
}

// new type of bot
type bot interface {
	// if you are a type in this program with a func called getGreeting
	// and you return a string and a honorary member of type "bot"
	// and honorary members can call getGreeting
	getGreeting() string
	getBotVersion() float64
	responseToUser(user) string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

//func printGreeting(sb spanishBot) {
//	fmt.Println(sb.getGreeting())
//}

// you can remove eb because we aren't using it in the function
// func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string {
	// super custom english greeting logic
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// super custom español greeting logic
	return "Hola!"
}
