package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3, 4, 5))
}

func sum(args ...int) int {
	result := 0
	for _, v := range args {
		result += v
	}

	return result
}
