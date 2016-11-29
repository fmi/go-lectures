package main

import "fmt"
import "time"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	select {
	case v1 := <-c:
		fmt.Printf("received %d in v1\n", v1)
	case v2 := <-c:
		fmt.Printf("received %d in v2\n", v2)
	case <-time.After(1 * time.Nanosecond):
		fmt.Printf("timeout\n")
	default:
		fmt.Printf("nothing!\n")
	}
}
