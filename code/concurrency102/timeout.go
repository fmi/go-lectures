package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(url string) <-chan *http.Response {
	ch := make(chan *http.Response)

	go func() {
		if response, err := http.Get(url); err == nil {
			ch <- response
		}
	}()

	return ch
}

func main() {
	select {
	case resp := <-fetch("https://www.google.com/search?q=golang"):
		fmt.Printf("received %v from google\n", resp.Status)
	case resp := <-fetch("https://www.bing.com/search?q=golang"):
		fmt.Printf("received %v from bing\n", resp.Status)
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("timed out\n")
	}
}
