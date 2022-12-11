package main

import (
	"testing"
)

var result int

func BenchmarkPartOne(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partOne(input, false)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partOne(input, true)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
