package main

import "fmt"

type greeter struct {
	Name string
}

func (g *greeter) SayHi(name string) {
	fmt.Printf("Hi, %s! I am %s (%p)\n", name, g.Name, g)
}

func main() {
	var p = &greeter{Name: "p"}
	pGreeter := p.SayHi

	fmt.Printf("Calling p.SayHi:\n")
	p.SayHi("Pesho")

	fmt.Printf("Calling pGreeter:\n")
	pGreeter("Pesho") // The same as p.SayHi("Pesho")

	unboundGreeter := (*greeter).SayHi
	fmt.Printf("Calling unboundGreeter:\n")
	unboundGreeter(p, "Gosho")

	unboundGreeter(&greeter{Name: "other"}, "Maria")
}
