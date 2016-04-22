package main

import "fmt"

func sumAndCount(args ...int) (int, int) {
	result := 0
	for _, v := range args {
		result += v
	}

	return result, len(args)
}

func main() {
	fmt.Printf("Резултатът от sumAndCount е: ")
	fmt.Println(sumAndCount(2, 3, 4, 5))
}
