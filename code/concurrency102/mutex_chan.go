package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		count   int
		countWg sync.WaitGroup
	)
	ch := make(chan struct{}, 1) // HL

	worker := func() {
		ch <- struct{}{} // HL
		count += 1
		<-ch // HL
		countWg.Done()
	}

	for i := 0; i < 8; i++ {
		countWg.Add(1)
		go worker()
	}

	countWg.Wait()
	fmt.Println("counter:", count)
}
