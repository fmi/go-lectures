package main

/*
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	cs := C.CString("42")            // alloc on C's heap
	defer C.free(unsafe.Pointer(cs)) // don't leak
	answer := C.atoi(cs)
	fmt.Println(answer)
}
