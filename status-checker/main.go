package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://udemy.com",
	}

	// make a channel of type string
	c := make(chan string)

	for _, link := range links {
		// making additional routine with keyword 'go', only in front of function calls.
		go makeRequest(link, c)

		// print the value as soon as recieved
		//fmt.Println(<-c)
	}
	// making use of function literal
	for l := range c {
		//go makeRequest(l, c)
		// function literal
		go func(link string) { // because we are passing the link in the bottom parenthesis
			time.Sleep(5 * time.Second)
			makeRequest(link, c)
		}(l) //<-- parenthesis to execute function literal
	}
	/*
		for {
			go makeRequest(<-c, c) // blocking line of code
		}
	*/
}

func makeRequest(st string, c chan string) {
	_, err := http.Get(st)
	if err != nil {
		fmt.Println("site", st, "is down")
		//time.Sleep(5 * time.Second)
		c <- st
		return
	}
	// send the value to channel
	fmt.Println("site", st, "is up")
	//time.Sleep(5 * time.Second)
	c <- st
}
