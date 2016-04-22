package main

import "fmt"

func main() {
	add := func(x, y int) int { return x + y }
	fmt.Println(add(4, 5))

	fmt.Println(func(x, y int) int { return x + y }(3, 8))
}
