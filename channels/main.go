package main

import (
	"fmt"
	"net/http"
)

// Site status
const (
	UP   = 1
	DOWN = 0
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)

		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// if err != nil {
	// 	return DOWN
	// }
	// return UP
	if err != nil {
		c <- link + " is down"
	} else {
		c <- link + " is up"
	}
}
