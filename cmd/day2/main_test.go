package main

import (
	"strings"
	"testing"
)

var result string

func BenchmarkPartOne(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partOne(strings.Split(input, "\n"))
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		r = partTwo(strings.Split(input, "\n"))
	}
	result = r
}
