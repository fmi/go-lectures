package main

import "fmt"

func main() {
	var value interface{}
	value = 20
	value = "asd"

	if str, ok := value.(string); ok {
		fmt.Printf("We have a %T: %v\n", str, str)
	} else {
		fmt.Println("Oops, not a string")
	}
}
