package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(output chan string) {
		for i := 0; i < 5; i++ {
			output <- fmt.Sprintf("sending N=%d", i)
		}
		close(output)
	}(ch)

	for i := 0; i < 7; i++ {
		val, ok := <-ch
		fmt.Printf("Recieved: %#v, %#v\n", val, ok)
	}

	ch <- fmt.Sprintf("where is your towel?")
}
