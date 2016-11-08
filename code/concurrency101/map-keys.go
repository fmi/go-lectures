package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("%+v\n", map[*[]int]int{&slice: 5})
}
