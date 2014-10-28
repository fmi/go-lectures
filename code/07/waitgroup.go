package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		wg.Add(1) // Increment the WaitGroup counter.
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			defer wg.Done()               // Decrement the counter when the goroutine completes.
			content, err := http.Get(url) // Fetch the URL.
			if err == nil {
				fmt.Println(url, content.Status)
			} else {
				fmt.Println(url, "has failed")
			}
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
