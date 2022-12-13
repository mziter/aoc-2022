package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

const (
	inputWidth  = 173
	inputHeight = 41
)

// this will represent our combined int using x and y
// we shift the x value and combine it with y to make
// a unique value that is fast to calculate
type Point = uint32
type Grid [][]byte

func pack(x, y uint16) Point {
	return (uint32(x)<<16 | uint32(y))
}

func unpack(point Point) (x, y uint16) {
	x = uint16(point >> 16)
	y = uint16(point)
	return
}

func main() {
	fmt.Println("DAY XII")
	fmt.Println("Part I: ", partOne(input))
	fmt.Println("Part II: ", partTwo(input))
}

func partTwo(input []byte) int {
	grid := make([][]byte, inputHeight)
	for i := range grid {
		grid[i] = make([]byte, inputWidth)
	}
	end, as := parseInputAndAs(input, grid)
	len := shortestPath(grid, as, end)
	return len
}

func partOne(input []byte) int {
	grid := make([][]byte, inputHeight)
	for i := range grid {
		grid[i] = make([]byte, inputWidth)
	}
	start, end := parseInput(input, grid)
	len := shortestPath(grid, []Point{start}, end)
	return len
}

func shortestPath(grid Grid, starts []Point, end Point) int {
	queue := make([]Point, 0, 200)
	pathLen := 0
	seen := make([]bool, inputHeight*inputWidth)
	queue = append(queue, starts...)
	for {
		len := len(queue)
		if len == 0 {
			panic("queue is empty man")
		}
		for i := 0; i < len; i++ {
			var point Point
			point, queue = queue[0], queue[1:]
			if point == end {
				return pathLen
			}
			addNeighbors(point, grid, seen, &queue)
		}
		pathLen++
	}
}

func addNeighbors(p Point, g Grid, seen []bool, q *[]Point) {
	px, py := unpack(p)
	x, y := int(px), int(py)
	if x-1 >= 0 {
		nx, ny := x-1, y
		neighbor := pack(uint16(nx), uint16(ny))
		pv := g[py][px]
		nv := g[ny][nx]
		idx := (ny * 173) + nx
		if nv <= pv+1 && !seen[idx] {
			seen[idx] = true
			*q = append(*q, neighbor)
		}
	}
	if x+1 < inputWidth {
		nx, ny := x+1, y
		pv := g[py][px]
		nv := g[ny][nx]
		idx := (ny * 173) + nx
		if nv <= pv+1 && !seen[idx] {
			seen[idx] = true
			neighbor := pack(uint16(nx), uint16(ny))
			*q = append(*q, neighbor)
		}
	}
	if y-1 >= 0 {
		nx, ny := x, y-1
		neighbor := pack(uint16(nx), uint16(ny))
		pv := g[py][px]
		nv := g[ny][nx]
		idx := (ny * 173) + nx
		if nv <= pv+1 && !seen[idx] {
			seen[idx] = true
			*q = append(*q, neighbor)
		}
	}
	if y+1 < inputHeight {
		nx, ny := x, y+1
		neighbor := pack(uint16(nx), uint16(ny))
		pv := g[py][px]
		nv := g[ny][nx]
		idx := (ny * 173) + nx
		if nv <= pv+1 && !seen[idx] {
			seen[idx] = true
			*q = append(*q, neighbor)
		}
	}
}

func parseInput(input []byte, grid Grid) (start, end Point) {
	reader := bytes.NewReader(input)
	s := bufio.NewScanner(reader)
	r := 0
	for s.Scan() {
		for i, b := range s.Bytes() {
			grid[r][i] = b
			if b == 'S' {
				start = pack(uint16(i), uint16(r))
				grid[r][i] = 'a'
			}
			if b == 'E' {
				end = pack(uint16(i), uint16(r))
				grid[r][i] = 'z'
			}
		}
		r++
	}
	return
}

func parseInputAndAs(input []byte, grid Grid) (Point, []Point) {
	reader := bytes.NewReader(input)
	as := make([]Point, 0, 2300)
	var end Point
	s := bufio.NewScanner(reader)
	r := 0
	for s.Scan() {
		for i, b := range s.Bytes() {
			grid[r][i] = b
			if b == 'a' {
				a := pack(uint16(i), uint16(r))
				grid[r][i] = 'a'
				as = append(as, a)
			}
			if b == 'E' {
				end = pack(uint16(i), uint16(r))
				grid[r][i] = 'z'
			}
		}
		r++
	}
	return end, as
}
