package main

import "fmt"

var subject string

func init() {
	subject = "world"
}

func main() {
	fmt.Printf("Hello, %s!\n", subject)
}
