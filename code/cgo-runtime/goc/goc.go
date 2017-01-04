package goc

import "C"

//export GreetFromGo
func GreetFromGo(name string) {
	println("Hello from Go, ", name)
}

func main() {
	// Needed by cgo in order to generate a library
}
