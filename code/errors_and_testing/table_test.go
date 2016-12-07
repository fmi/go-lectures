package main_test

import (
	"fmt"
	"testing"
	"time"
)

func Fibonacci(i int) int {
	return 42 / 2
}

// TABLE TEST START
func TestFibonacci(t *testing.T) {
	var tests = []struct {
		input, result int
	}{
		{input: 2, result: 1},
		{input: 8, result: 21},
		{input: 10, result: 55},
	}
	for _, test := range tests {
		testingFunc(t, test.input, test.result)
	}
}

// TABLE TEST OMIT

// SUB TEST START
func TestSubFibonacci(t *testing.T) {
	var tests = []struct {
		input, result int
	}{
		{input: 2, result: 1},
		{input: 8, result: 21},
		{input: 10, result: 55},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("Fibonacci(%d)", test.input), func(t *testing.T) {
			testingFunc(t, test.input, test.result)
		})
	}
}

// SUB TEST OMIT

// SUB BENCHMARK START
func BenchmarkSubFibonacci(b *testing.B) {
	var tests = []struct {
		input, result int
	}{
		{input: 2, result: 1},
		{input: 8, result: 21},
		{input: 10, result: 55},
	}
	for _, test := range tests {
		b.Run(fmt.Sprintf("BFibonacci(%d)", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				testingFunc(b, test.input, test.result)
			}
		})
	}
}

// SUB BENCHMARK OMIT

func testingFunc(tb testing.TB, input, expectedResult int) {
	if result := Fibonacci(input); result != expectedResult {
		tb.Fatalf("Expected %d for Fiboncci(%d) but got %d", expectedResult, input, result)
	}
}

// END testingFunc OMIT

var tests = []struct {
	input, result int
}{
	{input: 2, result: 1},
	{input: 8, result: 21},
	{input: 10, result: 55},
}

// SUB GROUP START
func TestGroupSubFibonacci(t *testing.T) {
	t.Run("group1", func(t *testing.T) {
		for _, test := range tests {
			t.Run(fmt.Sprintf("Fibonacci(%d)", test.input), func(t *testing.T) {
				var input, result = test.input, test.result
				t.Parallel()
				time.Sleep(time.Second)
				testingFunc(t, input, result)
			})

		}
		t.Run("NonParallel", func(t *testing.T) {
			t.Fatal("Just cause")
		})
		t.Fatal("Oops")
		t.Run("NonParallel2", func(t *testing.T) {
			t.Fatal("Just Cause II")
		})
	})
}

// SUB GROUP OMIT
