package interfaces

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

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
