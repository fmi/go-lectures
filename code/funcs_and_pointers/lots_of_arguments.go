package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3, 4, 5))
	sum()           //0
	sum(2, 3)       //5
	sum(2, 3, 4, 5) //14
}

func sum(args ...int) int {
	result := 0
	for _, v := range args {
		result += v
	}

	return result
}
