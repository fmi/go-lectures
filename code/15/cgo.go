package main

/*
#include <stdlib.h>
*/
import "C" // HL

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
