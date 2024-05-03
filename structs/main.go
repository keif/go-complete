package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)

	alex.firstName = "Alexis"
	alex.lastName = "Anderson, II"

	fmt.Println(alex)
}
