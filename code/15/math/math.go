package math

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"

func pow(a, b float64) float64 {
	return float64(C.pow(C.double(a), C.double(b)))

}
