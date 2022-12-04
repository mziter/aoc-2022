package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	exampleInput := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
2-8,3-7
`
	got := partOne([]byte(exampleInput))
	want := "3"
	if got != want {
		t.Errorf("wanted %s but got %s", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	exampleInput := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`
	got := partTwo([]byte(exampleInput))
	want := "4"
	if got != want {
		t.Errorf("wanted %s but got %s", want, got)
	}
}

var result string

func BenchmarkPartOne(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partOne(input)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partTwo(input)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
