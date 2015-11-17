package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boring := func(msg string) {
		for i := 0; ; i++ {
			fmt.Println(msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}

	go boring("boring!")
	fmt.Println("Listening.")
	time.Sleep(2 * time.Second) // HL
	fmt.Println("You are way too boring. I am leaving.")
}
