package main

import (
	"encoding/json"
	"fmt"
)

type Triangle struct {
	X, Y, Z int
}

// start OMIT
func (t *Triangle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		X, Y, Z int
		Shape   string
	}{
		X:     t.X,
		Y:     t.Y,
		Z:     t.Z,
		Shape: "Триъгълник",
	})
}

func main() {
	tr := &Triangle{X: 12, Y: 64, Z: 50}
	marshalledTr, _ := json.Marshal(tr)
	fmt.Printf("%s\n", marshalledTr)
}

// end OMIT
