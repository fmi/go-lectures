package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
func fanIn(input1, input2 <-chan string) <-chan string { // HL
	c := make(chan string)
	go func() { // HL
		for {
			select { // HL
			case s := <-input1:
				c <- s // HL
			case s := <-input2:
				c <- s // HL
			} // HL
		}
	}()
	return c
}

// FANIN END OMIT

// MAIN START OMIT
func main() {
	c := fanIn(talk("Boris"), talk("Kiril")) // HL
	for i := 0; i < 10; i++ {
		fmt.Println(<-c) // HL
	}
}

// MAIN END OMIT
