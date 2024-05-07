package main

import "fmt"

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
		fmt.Println(link)

	}
}
