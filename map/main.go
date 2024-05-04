package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)
	//colors["white"] = "#ffffff"
	//fmt.Println(colors["white"])

	colors := make(map[int]string)
	colors[10] = "#ffffff"

	delete(colors, 10)

	//colors := map[string]string{
	//	"red":   "#ff0000",
	//	"green": "#00ff00",
	//	"blue":  "#0000ff",
	//}

	fmt.Println(colors)
}
