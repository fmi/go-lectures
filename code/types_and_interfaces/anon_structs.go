package main

import "fmt"

// start OMIT
type Person struct {
	firstName, lastName string
}

func (p Person) Name() string {
	return p.firstName + " " + p.lastName
}

type Student struct {
	Person
	facultyNumber int16
}

func main() {
	s := Student{Person{"Иван", "Иванов"}, 123}
	fmt.Printf("We have a student with name %s and FN %d", s.Name(), s.facultyNumber)
}

// end OMIT
