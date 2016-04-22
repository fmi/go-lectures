package main

import "fmt"

func main() {
	var l int
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{7, 6, 5}
	arr := [3]int{1, 1, 1}

	// Копираме трите елемента от `slice2` в `slice1`
	l = copy(slice1, slice2) // slice1 = [7 6 5 4], l = 3
	fmt.Printf("slice1 = %v, l = %v\n", slice1, l)

	// Копираме края на slice1 в началото му
	l = copy(slice1, slice1[2:]) // slice1 = [5 4 5 4], l = 2
	fmt.Printf("slice1 = %v, l = %v\n", slice1, l)

	// Копираме slice1 в slice2
	l = copy(slice2, slice1) // slice2 = [5 4 5], l = 3
	// Копират се само първите 3 елемента, защото len(slice2) = 3
	fmt.Printf("slice2 = %v, l = %v\n", slice2, l)

}
