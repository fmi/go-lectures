package main

import "fmt"

var (
	name   string = "Чочко"
	age    uint8  = 27
	pName *string
)

func main() {
	pName = &name
	fmt.Printf("name е на адрес %p и има стойност %s\n", pName, name)
	fmt.Printf("age е на адрес %p и има стойност %d\n", &age, age)
}
