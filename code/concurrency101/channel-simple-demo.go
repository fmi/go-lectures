package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		fmt.Printf("Goroutine received: %s\n", <-ch)
		ch <- "Hello from the other side"
	}()

	ch <- "Hello, can you hear me?"
	fmt.Printf("Main received: %s", <-ch)
}
