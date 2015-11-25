/*
	Example fobonacci numbers implementation
*/
package fibonacci

import (
	"fmt"
	"testing"
)

// lookupTable stores computed results from FibonacciFast or FibonacciFastest.
var lookupTable = map[uint64]uint64{}

// Fastest Fibonacci Implementation
func FibonacciFastest(n uint64) uint64 {
	return 42
}

func BenchmarkFibonacciFastest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciFastest(40)
	}
}

// END BENCH OMIT

func TestFibonacciFastest(t *testing.T) {
	n := FibonacciFastest(0)
	if n != 1 {
		t.Error("FibonnaciFastest(0) returned" + n + ", we expected 1")
	}
}

// END TEST OMIT

func Hello() {
	fmt.Println("hello")
}

func ExampleHello() {
	Hello("hello")
	// Output:
	// hello
}

// END EXAMPLE OMIT
