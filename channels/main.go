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

	c := make(chan string)

	for _, link := range links {
		// "go" being added means run this function inside its own thread go routine
		// and runs checkLink inside it
		// concurrency - when code is running, if one thread blocks, another one is picked up and worked on. Single CPU
		// parallelism - multiple threads executed at "exact" same time. Requires multiple CPUs
		//go checkLink(link)
		go checkLink(link, c)
	}

	// infinite loop - whenever we pass in a value to a channel
	// receiving the value is a blocking operation
	//for {
	//	go checkLink(<-c, c)
	//}
	// equivalent to the above
	// wait for the channel to return some value, assign to l
	// the "verbose" option XD
	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) bool {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s might be down.\n", link)
		c <- link
		return false
	}

	fmt.Printf("%s is up.\n", link)
	c <- link
	return true
}
