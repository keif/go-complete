package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// check that we have an argument to work with
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	// get the filename from the CLI arguments
	fileName := os.Args[1]

	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer file.Close() // close file when done

	// read the file
	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
}
