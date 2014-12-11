package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func waitForMe(ch chan int) {
	defer wg.Done()
	for value := range ch {
		fmt.Println(value)
	}
	fmt.Println("My channel has closed")
}

func main() {
	wg.Add(1)
	ch := make(chan int)
	go waitForMe(ch)
	ch <- 64
	ch <- 128
	close(ch)
	wg.Wait()
}

// EOF OMIT
