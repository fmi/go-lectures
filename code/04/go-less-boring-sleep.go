package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boring("boring!") // HL
	fmt.Println("Listening.")
	time.Sleep(2 * time.Second) // HL
	fmt.Println("You are way too boring. I am leaving.")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
