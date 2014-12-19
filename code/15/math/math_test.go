package math

import "testing"

func TestPow(t *testing.T) {
	a := pow(2.0, 4.0)

	if a != 16.0 {
		t.Fatalf("we suck at math!")

	}
}
