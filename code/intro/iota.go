package main

import "fmt"

const (
	B  uint64 = iota + 1
	KB        = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Printf("Result is %d bytes", 2*KB)
}
