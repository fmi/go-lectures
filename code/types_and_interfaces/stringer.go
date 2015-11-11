package main

import (
	"fmt"
	"strconv"
)

// start OMIT
type myint uint64

func (i myint) String() string {
	return "myint in binary is " + strconv.FormatUint(uint64(i), 2)
}

func main() {
	var i myint
	i = 5
	fmt.Printf("Value: %s\n", i)
}

// end OMIT
