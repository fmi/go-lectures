package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MAIN START OMIT
func main() {
	kiro := talk("Kiro") // HL
	ned := talk("Ned")   // HL
	for i := 0; i < 5; i++ {
		fmt.Println(<-kiro)
		fmt.Println(<-ned)
	}
}

// MAIN END OMIT

// TALK START OMIT
func talk(msg string) <-chan string { // HL
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

// TALK END OMIT
