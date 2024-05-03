package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	fmt.Println(jim)
	jim.print()
	// &variable = give me the memory address of the value this variable is pointing at
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jimPointer.print()
}

// variable *word = type description - it means we're working with a pointer to a person
func (pointerToPerson *person) updateName(newFirstName string) {
	// *pointer = give me the value this memory address is pointing at
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
