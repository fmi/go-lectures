package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func waitForMe(ch chan int) {
	defer wg.Done()
	for {
		value, ok := <-ch
		if ok {
			fmt.Println(value)
		} else {
			fmt.Println("My channel has closed and all I got is this:", value)
			return
		}
	}
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
