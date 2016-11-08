package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		fmt.Println("SCV: Reportin' for duty")
		time.Sleep(2 * time.Second)
		fmt.Println("SCV: Job's finished!")
		c <- 1
	}()

	fmt.Println("Main does other time-consuming work...")
	time.Sleep(1 * time.Second)
	fmt.Println("Main is done")
	<-c

	fmt.Println("Everyone is done")
}
