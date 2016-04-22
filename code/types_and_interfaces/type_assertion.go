package main

import "fmt"

func main() {
	// START OMIT
	var value interface{}
	value = 20
	value = "asd"
	str, ok := value.(string)
	if ok {
		fmt.Println("We have a string: ", str)
	} else {
		fmt.Println("Oops, not a string")
	}
	// END OMIT
}
