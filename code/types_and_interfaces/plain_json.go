package main

import (
	"encoding/json"
	"fmt"
)

type Rectangle struct {
	X, Y int
}

func main() {
	// start OMIT
	var empty interface{}
	emptyRect := new(Rectangle)

	rect := &Rectangle{X: 12, Y: 64}

	marshalledRect, _ := json.Marshal(rect)
	fmt.Printf("%s\n", marshalledRect)

	json.Unmarshal(marshalledRect, emptyRect)
	fmt.Printf("%#v\n", emptyRect)

	json.Unmarshal(marshalledRect, &empty)
	fmt.Printf("%#v\n", empty)
	// end OMIT
}
