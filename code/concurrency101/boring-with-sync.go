package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // HL

	boring := func(msg string) {
		defer wg.Done() // HL
		for i := 0; i < 8; i++ {
			fmt.Println(msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}

	wg.Add(1) // HL
	go boring("boring!")
	fmt.Println("Waiting for the boring function to do its work")
	wg.Wait() // HL
}
