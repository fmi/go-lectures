package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
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
			defer wg.Done()
			once.Do(onceBody)
			once.Do(anotherBody)
		}()
	}
	wg.Wait()
}
