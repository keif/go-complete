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

	bs := make([]byte, 999999)
	//bs := make([]byte, resp.ContentLength)
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	//bytesRead, err := resp.Body.Read(bs)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(1)
	//}
	//fmt.Println(string(bs[:bytesRead]))

}
