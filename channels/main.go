package main

import (
	"fmt"
	"net/http"
	"time"
)

// Site status
const (
	UP   = 1
	DOWN = 0
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://twitter.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for range links {
	// 	fmt.Println(<-c)
	// }

	// infinite loop
	// for {
	// 	go checkLink(<-c, c)
	// }

	// range over channel
	// means wait for channel to return
	// then assign to var l
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// if err != nil {
	// 	return DOWN
	// }
	// return UP
	if err != nil {
		fmt.Println(link + " might be down")
		c <- link
	} else {
		fmt.Println(link + " is up")
		c <- link
	}
}
