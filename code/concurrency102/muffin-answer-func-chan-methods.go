package main

import "fmt"

type queue chan string
type callableStriner func() string

func (q queue) add(s string) {
	q <- s
}
func (q queue) poll() string {
	return <-q
}
func (cs callableStriner) String() string {
	return cs()
}

func main() {
	q := make(queue, 10)
	q.add("test")
	fmt.Printf("%s\n", callableStriner(q.poll))
}
