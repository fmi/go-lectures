package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once // HL
	var wg sync.WaitGroup

	onceBody := func() {
		fmt.Println("Only once")
	}
	anotherBody := func() {
		fmt.Println("Another")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(onceBody)    // HL
			once.Do(anotherBody) // HL
			wg.Done()
		}()
	}
	wg.Wait()
}
