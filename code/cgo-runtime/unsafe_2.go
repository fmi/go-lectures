package main

import "unsafe"
import "fmt"

type slice struct {
	array      unsafe.Pointer
	size, _cap int
}

func main() {
	var p = []string{"Hello", " "}
	p = append(p, "World!")
	var s = (*slice)(unsafe.Pointer(&p))
	var sizeOfString = unsafe.Sizeof("")
	fmt.Printf("size=%d, cap=%d\n", s.size, s._cap)
	for i := 0; s.size > i; i++ {
		fmt.Printf("[%d]: `%s`\n", i,
			*(*string)(unsafe.Pointer(uintptr(s.array) + uintptr(i)*sizeOfString)))
	}
}
