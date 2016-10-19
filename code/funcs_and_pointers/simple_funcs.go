package main

import (
	"fmt"
	"strings"
)

func ping() {
	fmt.Println("pong")
}

func myprint(s string) {
	fmt.Println(s)
}

func shout(s string) string {
	return strings.ToUpper(s) + "!!!1!"
}

func main() {
	ping()
	myprint("woot")
	myprint(shout("woot"))
}
