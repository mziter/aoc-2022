package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTreeGrid(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	grid := CreateTreeGrid([]byte(input))
	assert.Equal(t, 5, len(grid))
	assert.Equal(t, 5, len(grid[0]))
	assert.Equal(t, byte('3'), grid[0][0])
	assert.Equal(t, byte('0'), grid[4][4])
	assert.Equal(t, byte('0'), grid[4][4])
	assert.Equal(t, byte('5'), grid[2][1])
	assert.Equal(t, byte('9'), grid[3][4])
}

func TestCreateResultGrid(t *testing.T) {
	h := 5
	w := 3
	grid := CreateResultGrid(h, w)

	assert.Equal(t, h, len(grid))
	assert.Equal(t, w, len(grid[0]))
}

func TestScanOutUp(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	treeBytes := CreateTreeGrid([]byte(input))
	actual := scanOutUp(treeBytes, 3, 2)
	expected := 2
	assert.Equal(t, expected, actual)

	actual = scanOutUp(treeBytes, 4, 3)
	expected = 4
	assert.Equal(t, expected, actual)

	actual = scanOutUp(treeBytes, 0, 3)
	expected = 0
	assert.Equal(t, expected, actual)
}

func TestScanOutDown(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	treeBytes := CreateTreeGrid([]byte(input))
	actual := scanOutDown(treeBytes, 3, 2)
	expected := 1
	assert.Equal(t, expected, actual)

	actual = scanOutDown(treeBytes, 4, 3)
	expected = 0
	assert.Equal(t, expected, actual)

	actual = scanOutDown(treeBytes, 0, 3)
	expected = 4
	assert.Equal(t, expected, actual)
}

func TestScanOutLeft(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	treeBytes := CreateTreeGrid([]byte(input))
	actual := scanOutLeft(treeBytes, 3, 2)
	expected := 2
	assert.Equal(t, expected, actual)

	actual = scanOutLeft(treeBytes, 4, 3)
	expected = 3
	assert.Equal(t, expected, actual)

	actual = scanOutLeft(treeBytes, 0, 3)
	expected = 3
	assert.Equal(t, expected, actual)

	actual = scanOutLeft(treeBytes, 2, 0)
	expected = 0
	assert.Equal(t, expected, actual)
}

func TestScanOutRight(t *testing.T) {
	input := `30373
25512
65332
33549
35390`
	treeBytes := CreateTreeGrid([]byte(input))
	actual := scanOutRight(treeBytes, 3, 2)
	expected := 2
	assert.Equal(t, expected, actual)

	actual = scanOutRight(treeBytes, 4, 3)
	expected = 1
	assert.Equal(t, expected, actual)

	actual = scanOutRight(treeBytes, 0, 3)
	expected = 1
	assert.Equal(t, expected, actual)

	actual = scanOutRight(treeBytes, 3, 4)
	expected = 0
	assert.Equal(t, expected, actual)
}

func TestScanLeft(t *testing.T) {
	visible := CreateResultGrid(5, 5)
	input := `30373
25512
65332
33549
35390`
	expected := [][]bool{
		{false, false, false, true, true},
		{false, false, true, false, true},
		{true, true, false, true, true},
		{false, false, false, false, true},
		{false, false, false, true, true},
	}
	treeBytes := CreateTreeGrid([]byte(input))
	scanVisibleLeft(treeBytes, visible)
	assert.ElementsMatch(t, expected, visible)
}

func TestScanRight(t *testing.T) {
	visible := CreateResultGrid(5, 5)
	input := `30373
25512
65332
33549
35390`
	expected := [][]bool{
		{true, false, false, true, false},
		{true, true, false, false, false},
		{true, false, false, false, false},
		{true, false, true, false, true},
		{true, true, false, true, false},
	}
	treeBytes := CreateTreeGrid([]byte(input))
	scanVisibleRight(treeBytes, visible)
	assert.ElementsMatch(t, expected, visible)
}

func TestScanUp(t *testing.T) {
	visible := CreateResultGrid(5, 5)
	input := `30373
25512
65332
33549
35390`
	expected := [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{true, false, false, false, false},
		{false, false, true, false, true},
		{true, true, true, true, true},
	}
	treeBytes := CreateTreeGrid([]byte(input))
	scanVisibleUp(treeBytes, visible)
	assert.ElementsMatch(t, expected, visible)
}

func TestScanDown(t *testing.T) {
	visible := CreateResultGrid(5, 5)
	input := `30373
25512
65332
33549
35390`
	expected := [][]bool{
		{true, true, true, true, true},
		{false, true, true, false, false},
		{true, false, false, false, false},
		{false, false, false, false, true},
		{false, false, false, true, false},
	}
	treeBytes := CreateTreeGrid([]byte(input))
	scanVisibleDown(treeBytes, visible)
	assert.ElementsMatch(t, expected, visible)
}

var gridResult [][]byte

func BenchmarkCreateGrid(b *testing.B) {
	var r [][]byte
	input := `30373
25512
65332
33549
35390`
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = CreateTreeGrid([]byte(input))
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	gridResult = r
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
		r = partTwo(input)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
