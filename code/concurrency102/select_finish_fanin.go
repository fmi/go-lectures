package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// TALK START OMIT
func talk(msg string) <-chan string { // HL
	c := make(chan string)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(r.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// TALK END OMIT

// FANIN START OMIT
func fanIn(input1, input2 <-chan string, finish chan struct{}) <-chan string { // HL
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			case <-finish: // HL
				close(c) // HL
				wg.Done()
				return
			}
		}
	}()
	return c
}

// FANIN END OMIT

// MAIN START OMIT
func main() {
	wg.Add(1)
	finish := make(chan struct{})            // HL
	c := fanIn(talk("A"), talk("B"), finish) // HL
	for value := range c {
		fmt.Println(value)
		if len(value) > 4 {
			close(finish)
		}
	}
	wg.Wait()
}

// MAIN END OMIT
