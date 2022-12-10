package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateTail(t *testing.T) {
	tcs := []struct {
		headX     int
		headY     int
		tailX     int
		tailY     int
		expectedX int
		expectedY int
		name      string
	}{
		{0, 1, 0, 0, 0, 0, "head being one up should not move tail"},
		{0, 2, 0, 0, 0, 1, "head being two up should move tail up one"},

		{1, 0, 0, 0, 0, 0, "head being one right should not move tail"},
		{2, 0, 0, 0, 1, 0, "head being two right should move tail right one"},

		{1, 1, 0, 0, 0, 0, "head being one away diagonally should not move tail"},
		{2, 2, 0, 0, 1, 1, "head being two away diagonally should move tail diagonally"},
		{2, 2, 0, 0, 1, 1, "head being two away diagonally should move tail diagonally"},
		{1, 2, 0, 0, 1, 1, "head being two up and one away should move tail diagonally"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			actualX, actualY := updateTail(tc.headX, tc.headY, tc.tailX, tc.tailY)
			assert.Equal(t, tc.expectedY, actualY)
			assert.Equal(t, tc.expectedX, actualX)
		})
	}
}

func TestPackedPointsShoudlBeUnique(t *testing.T) {
	x := uint16(3)
	y := uint16(6)
	pointOne := pack(x, y)
	x = uint16(6)
	y = uint16(3)
	pointTwo := pack(x, y)

	assert.NotEqual(t, pointOne, pointTwo)
}
func TestPackUnpack(t *testing.T) {
	x := uint16(4)       // 00000000 00000100
	y := uint16(3)       // 00000000 00000011
	actual := pack(x, y) // 00000000 00000100 00000000 00000011
	assert.Equal(t, uint32(262147), actual)

	x, y = unpack(actual)
	assert.Equal(t, uint16(4), x)
	assert.Equal(t, uint16(3), y)
}

var result int

func BenchmarkPartOne(b *testing.B) {
	var r int
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
	var r int
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partTwo(input, 9)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
