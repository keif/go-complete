package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, number := range numbers {
		fmt.Printf("The number %d is %s\n", number, isEvenOrOdd(number))
	}
}

func isEvenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	}
	return "Odd"
}
