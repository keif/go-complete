package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := "https://www.google.com"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
