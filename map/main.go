package main

import "fmt"

func main() {
	// maps are great for dynamic additions, like adding a new color
	// we want them to be tightly coupled, i.e. "these are all colors"
	// a struct is more tightly coupled in the "if we only work with the initially created key/value pairs"
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}

	colors["yellow"] = "#123456"
	delete(colors, "yellow")

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %c is %s\n", color, hex)
	}
}
