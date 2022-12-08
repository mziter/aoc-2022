package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func simpleFS() *FSNode {
	a := &FSNode{
		name:     "a",
		size:     10,
		isDir:    false,
		children: make([]*FSNode, 0),
	}

	c := &FSNode{
		name:     "c",
		size:     1,
		isDir:    false,
		children: make([]*FSNode, 0),
	}

	b := &FSNode{
		name:  "b",
		size:  0,
		isDir: true,
		children: []*FSNode{
			c,
		},
	}

	return &FSNode{
		name:  "/",
		size:  0,
		isDir: true,
		children: []*FSNode{
			a, b,
		},
	}
}

func TestFilesystem(t *testing.T) {
	root := simpleFS()
	root.CalculateSize()
	want := 11
	got := root.size
	assert.Equal(t, want, got)
	root.PrettyPrint()
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
