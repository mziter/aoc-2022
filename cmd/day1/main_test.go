package main

import (
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	lines := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	want := "24000"
	have := partOne(lines)
	if want != have {
		t.Errorf("Wanted answer of %s, but had answer %s", want, have)
	}
}

func TestPartTwo(t *testing.T) {
	lines := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	want := "45000"
	have := partTwo(lines)
	if want != have {
		t.Errorf("Wanted answer of %s, but had answer %s", want, have)
	}
}

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
