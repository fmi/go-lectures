package main

import "fmt"

func counter(start int) func() int {
	count := start
	return func() int {
		count++
		return count
	}
}

func main() {
	c := counter(3)
	fmt.Println("The first 3 values are", c(), c(), c())
}
