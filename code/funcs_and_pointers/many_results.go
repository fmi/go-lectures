package main

import "fmt"

func main() {
	sum, count := sumAndCount(2, 3, 4, 5)
	fmt.Println("Резултатът от sumAndCount() е", sum, count)
}

func sumAndCount(args ...int) (int, int) {
	result := 0
	for _, v := range args {
		result += v
	}

	return result, len(args)
}
