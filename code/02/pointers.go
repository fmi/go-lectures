package main

import "fmt"

var (
	name   string = "Чочко"
	age    uint8  = 27
	p_name *string
)

func main() {
	p_name = &name
	fmt.Printf("name е на адрес %p и има стойност %s\n", p_name, name)
	fmt.Printf("age е на адрес %p и има стойност %d\n", &age, age)
}
