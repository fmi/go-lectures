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

	for val := range ch {
		fmt.Printf("Recieved: %#v\n", val)
	}

	ch <- fmt.Sprintf("where is your towel?")
}
