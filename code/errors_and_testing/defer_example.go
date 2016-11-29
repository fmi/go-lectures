package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("first")

	if false { // ex. try to open a file
		fmt.Println("error")
		return
	}
	defer fmt.Println("second")

	fmt.Println("done")
}
