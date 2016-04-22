package main

import "fmt"

func main() {
	x := []int{2, 3, 5, 7, 11}
	y := x[1:3]

	fmt.Println("len(x) =", len(x), ", cap(x) =", cap(x))
	fmt.Println("len(y) =", len(y), ", cap(y) =", cap(y))
}
