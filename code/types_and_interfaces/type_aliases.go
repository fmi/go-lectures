package main

// START OMIT
type integer int

func isPositive(a integer) bool {
	return a > 0
}

type integerAlias = int

func isPositive1(a integerAlias) bool {
	return a > 0
}

func main() {
	a := 5
	isPositive(a) // cannot use a (type int) as type integer in argument to isPositive
	isPositive1(a)
}

// END OMIT
