package main

import "fmt"

type integer int

func (i *integer) Increment() {
	fmt.Println("Incrementing...")
	*i++
}

func main() {
	var number *integer
	number.Increment()
	fmt.Println(number)
}

// end OMIT
