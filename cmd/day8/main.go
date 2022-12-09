package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println("DAY IIX")
	fmt.Println("PART I: ", strconv.Itoa(partOne(input)))
	fmt.Println("PART II: ", strconv.Itoa(partTwo(input)))
}

// notes: could be optimized by stopping at first visible tree when scanning in opposite direction (since values
// are already proven to be same size or smaller). Would require a separate visible grid for left/right, up/down, but could save lots
// of checks
func partOne(treeBytes []byte) int {
	treeHeights := CreateTreeGrid(treeBytes)
	dim := len(treeHeights[0])
	visible := CreateResultGrid(dim, dim)
	scanVisibleUp(treeHeights, visible)
	scanVisibleDown(treeHeights, visible)
	scanVisibleLeft(treeHeights, visible)
	scanVisibleRight(treeHeights, visible)
	count := 0
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			if visible[r][c] {
				count++
			}
		}
	}
	return count
}

func partTwo(treeBytes []byte) int {
	treeHeights := CreateTreeGrid(treeBytes)
	dim := len(treeHeights[0])
	visible := CreateResultGrid(dim, dim)
	scanVisibleUp(treeHeights, visible)
	scanVisibleDown(treeHeights, visible)
	scanVisibleLeft(treeHeights, visible)
	scanVisibleRight(treeHeights, visible)
	max := 0
	for r := 0; r < dim; r++ {
		for c := 0; c < dim; c++ {
			if visible[r][c] {
				if r == 0 || c == 0 || r == len(treeHeights)-1 || c == len(treeHeights[0])-1 {
					continue
				}
				up := scanOutUp(treeHeights, r, c)       // scan all directions?
				down := scanOutDown(treeHeights, r, c)   // scan all directions?
				left := scanOutLeft(treeHeights, r, c)   // scan all directions?
				right := scanOutRight(treeHeights, r, c) // scan all directions?
				score := up * down * left * right
				if score > max {
					max = score
				}
			}
		}
	}
	return max

}

func scanOutDown(treeBytes [][]byte, r int, c int) int {
	bottomEdge := len(treeBytes) - 1
	if r == bottomEdge {
		return 0
	}
	count := 1
	for y := r + 1; y < bottomEdge; y++ {
		height := treeBytes[r][c]
		if height > treeBytes[y][c] {
			count++
		} else {
			return count
		}
	}
	return count
}

func scanOutLeft(treeBytes [][]byte, r int, c int) int {
	leftEdge := 0
	if c == leftEdge {
		return 0
	}
	count := 1
	for x := c - 1; x > 0; x-- {
		height := treeBytes[r][c]
		if height > treeBytes[r][x] {
			count++
		} else {
			return count
		}
	}
	return count
}

func scanOutRight(treeBytes [][]byte, r int, c int) int {
	rightEdge := len(treeBytes[0]) - 1
	if c == rightEdge {
		return 0
	}
	count := 1
	for x := c + 1; x < rightEdge; x++ {
		height := treeBytes[r][c]
		if height > treeBytes[r][x] {
			count++
		} else {
			return count
		}
	}
	return count
}

func scanOutUp(treeBytes [][]byte, r int, c int) int {
	topEdge := 0
	if r == topEdge {
		return 0
	}
	count := 1
	for y := r - 1; y > 0; y-- {
		height := treeBytes[r][c]
		if height > treeBytes[y][c] {
			count++
		} else {
			return count
		}
	}
	return count
}

func scanVisibleUp(treeBytes [][]byte, visible [][]bool) {
	for c := 0; c < len(treeBytes[0]); c++ {
		max := byte('0') - 1
		for r := len(treeBytes) - 1; r >= 0; r-- {
			height := treeBytes[r][c]
			if height > max {
				visible[r][c] = true
				max = height
			}
		}
	}
}

func scanVisibleDown(treeBytes [][]byte, visible [][]bool) {
	for c := 0; c < len(treeBytes[0]); c++ {
		max := byte('0') - 1
		for r := 0; r < len(treeBytes); r++ {
			height := treeBytes[r][c]
			if height > max {
				visible[r][c] = true
				max = height
			}
		}
	}
}

func scanVisibleRight(treeBytes [][]byte, visible [][]bool) {
	for r := 0; r < len(treeBytes); r++ {
		max := byte('0') - 1
		for c := 0; c < len(treeBytes[0]); c++ {
			height := treeBytes[r][c]
			if height > max {
				visible[r][c] = true
				max = height
			}
		}
	}
}

func scanVisibleLeft(treeBytes [][]byte, visible [][]bool) {
	for r := 0; r < len(treeBytes); r++ {
		max := byte('0') - 1
		for c := len(treeBytes[0]) - 1; c >= 0; c-- {
			height := treeBytes[r][c]
			if height > max {
				visible[r][c] = true
				max = height
			}
		}
	}
}

func CreateTreeGrid(treeBytes []byte) [][]byte {
	lines := bytes.Split(treeBytes, []byte{'\n'})
	return lines
}

func CreateResultGrid(h int, w int) [][]bool {
	grid := make([][]bool, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]bool, w)
	}
	return grid
}
