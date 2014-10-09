package main

import (
	"fmt"
)

func foo(bar func(int, float64) float64) float64 {
	return bar(5, 3.2)
}

func createRandomGenerator() func() int {
	return func() int {
		return 4
	}
}

func main() {
	fmt.Println(createRandomGenerator()())
}
