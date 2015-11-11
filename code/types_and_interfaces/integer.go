package main

import "fmt"

type integer int

func (i integer) Abs() integer {
	switch {
	case i < 0:
		return -i
	case i == 0:
		return 0
	default:
		return i
	}
}

func (i *integer) Increment() {
	*i++
}

func main() {
	var number integer
	number = -42

	number.Increment()
	fmt.Println(number.Abs())
}

// end OMIT
