package main

import "fmt"

type example struct {
	attrs map[string]int
}

func main() {
	e := new(example)
	e.attrs = make(map[string]int)
	e.attrs["h"] = 42
	fmt.Println(e)
}
