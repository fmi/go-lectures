package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func randomFeed(count, max int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			c <- rand.Intn(max)
		}
		close(c)
	}()

	return c
}

func main() {
	feed := randomFeed(10, 100)
	for v := range feed {
		fmt.Println(v)
	}
}
