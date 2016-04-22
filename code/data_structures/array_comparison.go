package main

import "fmt"

func main() {
	dqdo := [4]uint32{0, 1, 2, 3}
	baba := [4]uint32{0, 1, 2, 3}

	fmt.Printf("The two arrays are identical: %t\n", dqdo == baba)
}
