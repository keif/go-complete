package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.bing.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.golang.org",
		"http://www.amazon.com",
	}

	for _, link := range links {
		// "go" being added means run this function inside its own thread go routine
		// and runs checkLink inside it
		go checkLink(link)
	}
}

func checkLink(link string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("The link %s might be down.\n", link)
		return false
	}

	fmt.Printf("The link %s is up.\n", link)
	return true
}
