package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		count    int
		countMtx sync.Mutex // HL
		countWg  sync.WaitGroup
	)

	worker := func() {
		countMtx.Lock() // HL
		count += 1
		countMtx.Unlock() // HL
		countWg.Done()
	}

	for i := 0; i < 8; i++ {
		countWg.Add(1)
		go worker()
	}

	countWg.Wait()
	fmt.Println("counter:", count)
}
