package main

import "fmt"

func main() {
	fmt.Println(sumAndCount(2, 3, 4, 5))
}

func sumAndCount(args ...int) (result int, count int) {
	count = len(args)
	for _, v := range args {
		result += v
	}

	return
}
