package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointScanner(t *testing.T) {
	input := "534,28 -> 452,394 -> 484,36"
	exs := []int{534, 452, 484}
	eys := []int{28, 394, 36}

	xs := make([]int, 0)
	ys := make([]int, 0)
	ps := NewPointScanner([]byte(input))
	i := 0
	for ps.Scan() {
		x, y := ps.Point()
		xs = append(xs, x)
		ys = append(ys, y)
		i++
	}
	assert.Equal(t, exs, xs)
	assert.Equal(t, eys, ys)
}

var result int

func BenchmarkPartOne(b *testing.B) {
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partOne()
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
		r = partTwo()
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
