package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Wait for my announcement...")
	Announce("Hello, from Goroutine!", time.Second * 2)
	time.Sleep(time.Second * 3)
}

func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()
}
