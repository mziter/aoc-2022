package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/mziter/aoc-2022/pkg/strutil"
)

// this will represent our combined int using x and y
// we shift the x value and combine it with y to make
// a unique value that is fast to calculate
type Point = uint32

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day X")
	fmt.Println("Part I: ", partOne(input))
	fmt.Println("Part II: ", partTwo(input, 9))
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func pack(x, y uint16) Point {
	return (uint32(x)<<16 | uint32(y))
}

func unpack(point Point) (x, y uint16) {
	x = uint16(point >> 16)
	y = uint16(point)
	return
}

func updateTail(headX, headY, tailX, tailY int) (int, int) {
	distX := abs(headX - tailX)
	distY := abs(headY - tailY)
	if distX > 2 || distY > 2 {
		panic("distance should never be greater than two")
	}
	if distX == 2 || distY == 2 {
		if headX > tailX {
			tailX++
		}
		if headX < tailX {
			tailX--
		}
		if headY > tailY {
			tailY++
		}
		if headY < tailY {
			tailY--
		}
	}
	return tailX, tailY
}

func parseInstructionBytes(s *bufio.Scanner) (dir byte, amount int) {
	b := s.Bytes()
	dir = b[0]
	n, err := strutil.ParseUintBytes(b[2:], 10, strconv.IntSize)
	if err != nil {
		panic(err)
	}
	amount = int(n)
	return

}

func partOne(content string) int {
	headX := 0
	headY := 0
	tailX := 0
	tailY := 0
	point := pack(uint16(tailX), uint16(tailY))
	visited := make(map[uint32]bool, 7000)
	visited[point] = true

	r := strings.NewReader(content)
	s := bufio.NewScanner(r)
	for s.Scan() {
		dir, amount := parseInstructionBytes(s)
		switch dir {
		case 'U':
			for i := 0; i < amount; i++ {
				headY++
				tailX, tailY = updateTail(headX, headY, tailX, tailY)
				visited[pack(uint16(tailX), uint16(tailY))] = true
			}
		case 'D':
			for i := 0; i < amount; i++ {
				headY--
				tailX, tailY = updateTail(headX, headY, tailX, tailY)
				visited[pack(uint16(tailX), uint16(tailY))] = true
			}
		case 'L':
			for i := 0; i < amount; i++ {
				headX--
				tailX, tailY = updateTail(headX, headY, tailX, tailY)
				visited[pack(uint16(tailX), uint16(tailY))] = true
			}
		case 'R':
			for i := 0; i < amount; i++ {
				headX++
				tailX, tailY = updateTail(headX, headY, tailX, tailY)
				visited[pack(uint16(tailX), uint16(tailY))] = true
			}
		}

	}
	return len(visited)
}

func partTwo(content string, numTailKnots int) int {
	knots := make([]int, 2*(numTailKnots+1))
	visited := make(map[uint32]bool, 7000)
	startPoint := pack(uint16(0), uint16(0))
	visited[startPoint] = true

	r := strings.NewReader(content)
	s := bufio.NewScanner(r)
	for s.Scan() {
		dir, amount := parseInstructionBytes(s)
		switch dir {
		case 'U':
			for i := 0; i < amount; i++ {
				knots[1]++
				hx := knots[0]
				hy := knots[1]
				for k := 2; k < len(knots)-1; k += 2 {
					knots[k], knots[k+1] = updateTail(hx, hy, knots[k], knots[k+1])
					hx = knots[k]
					hy = knots[k+1]
				}
				visited[pack(uint16(hx), uint16(hy))] = true
			}
		case 'D':
			for i := 0; i < amount; i++ {
				knots[1]--
				hx := knots[0]
				hy := knots[1]
				for k := 2; k < len(knots)-1; k += 2 {
					knots[k], knots[k+1] = updateTail(hx, hy, knots[k], knots[k+1])
					hx = knots[k]
					hy = knots[k+1]
				}
				visited[pack(uint16(hx), uint16(hy))] = true
			}
		case 'L':
			for i := 0; i < amount; i++ {
				knots[0]--
				hx := knots[0]
				hy := knots[1]
				for k := 2; k < len(knots)-1; k += 2 {
					knots[k], knots[k+1] = updateTail(hx, hy, knots[k], knots[k+1])
					hx = knots[k]
					hy = knots[k+1]
				}
				visited[pack(uint16(hx), uint16(hy))] = true
			}
		case 'R':
			for i := 0; i < amount; i++ {
				knots[0]++
				hx := knots[0]
				hy := knots[1]
				for k := 2; k < len(knots)-1; k += 2 {
					knots[k], knots[k+1] = updateTail(hx, hy, knots[k], knots[k+1])
					hx = knots[k]
					hy = knots[k+1]
				}
				visited[pack(uint16(hx), uint16(hy))] = true
			}
		}
	}
	return len(visited)
}
