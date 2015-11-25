package main

import (
	"fmt"
)

func deferExample() {
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf(" %v", i)
		}()
	}
}

func main() {
	deferExample()
}
