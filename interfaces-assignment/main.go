package main

import "fmt"

// interfaces
type shape interface {
	getArea() float64
}

// structs
type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	triangle := triangle{
		base:   20,
		height: 10,
	}

	square := square{
		sideLength: 5,
	}

	printArea(triangle)
	printArea(square)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
