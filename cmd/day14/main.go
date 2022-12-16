package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"

	"github.com/mziter/aoc-2022/pkg/strutil"
)

//go:embed input.txt
var input string

const (
	offsetY = 0             // subtract minimum value to place smallest value at 0 to conserve space
	offsetX = -250          // subtract minimum value to place smallest value at 0 to conserve space
	sourceY = 0             // sand source coordinates
	sourceX = 500 + offsetX // sand source coordinates
	floorY  = 157 + 2
)

func main() {
	fmt.Println("DAY XIV")
	fmt.Println("Part I: ", partOne())
	fmt.Println("Part II: ", partTwo())
}

func printGrid(grid [][]byte) {
	for i := 0; i < 300; i++ {
		if i == 500+offsetX {
			fmt.Printf("*")
		} else {
			fmt.Printf(" ")
		}
		fmt.Printf("\n")
	}
	for _, r := range grid {
		for _, v := range r {
			if v == 0 {
				fmt.Printf(".")
			}
			fmt.Printf("%s", string(v))
		}
		fmt.Printf("\n")
	}
}

func partOne() int {
	grid := parseInput(input, false)
	grainsSettled := simulateSandUntilFalloff(grid, sourceX, 0, 0)
	return grainsSettled
}

func partTwo() int {
	grid := parseInput(input, true)
	grainsSettled := 0
	lastDownX := sourceX
	lastDownY := 0
	x := lastDownX
	y := lastDownY
	blocked := false
	for !blocked {
		// try to go straight down first
		if grid[y+1][x] == 0 {
			y = y + 1
			continue
		}
		// try to go down+left next
		if grid[y+1][x-1] == 0 {
			x = x - 1
			y = y + 1
			continue
		}
		// try to go down+right next
		if grid[y+1][x+1] == 0 {
			x = x + 1
			y = y + 1
			continue
		}
		// settle
		if y == 0 && x == sourceX {
			grid[y][x] = 'o'
			blocked = true
			grainsSettled++
			break
		}
		grid[y][x] = 'o'
		x = sourceX
		y = sourceY
		grainsSettled++
	}
	return grainsSettled
}

func simulateSandUntilFalloff(grid [][]byte, x, y, count int) int {
	if y >= len(grid)-1 {
		return count
	}
	// try to go straight down first
	if grid[y+1][x] == 0 {
		return simulateSandUntilFalloff(grid, x, y+1, count)
	}
	// try to go down+left next
	if grid[y+1][x-1] == 0 {
		return simulateSandUntilFalloff(grid, x-1, y+1, count)
	}
	// try to go down+right next
	if grid[y+1][x+1] == 0 {
		return simulateSandUntilFalloff(grid, x+1, y+1, count)
	}
	// settle
	grid[y][x] = 'o'
	return simulateSandUntilFalloff(grid, sourceX, sourceY, count+1)
}

func parseInput(input string, recordFloor bool) [][]byte {
	r := strings.NewReader(input)
	s := bufio.NewScanner(r)
	grid := make([][]byte, 160)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, 500)
	}
	var ps *PointScanner
	for s.Scan() {
		if ps == nil {
			ps = NewPointScanner(s.Bytes())
		} else {
			ps.Reload(s.Bytes())
		}
		lastX := 0
		lastY := 0
		first := true
		for ps.Scan() {
			x, y := ps.Point()
			if first {
				lastX = x
				lastY = y
				first = false
				continue
			}
			recordRock(grid, lastX, lastY, x, y)
			lastX = x
			lastY = y
		}
	}
	if recordFloor {
		for x := 0; x < len(grid[0]); x++ {
			grid[floorY][x] = '#'
		}
	}
	return grid
}

func recordRock(grid [][]byte, fromX, fromY, toX, toY int) {
	if fromX == toX {
		if fromY > toY {
			for i := toY; i <= fromY; i++ {
				grid[i+offsetY][fromX+offsetX] = '#'
			}
		} else {
			for i := fromY; i <= toY; i++ {
				grid[i+offsetY][fromX+offsetX] = '#'
			}
		}
		return
	}
	if fromY == toY {
		if fromX > toX {
			for i := toX; i <= fromX; i++ {
				grid[fromY+offsetY][i+offsetX] = '#'
			}
		} else {
			for i := fromX; i <= toX; i++ {
				grid[fromY+offsetY][i+offsetX] = '#'
			}
		}
		return
	}
	fmt.Printf("record fromx=%d, fromy=%d, tox=%d, toy=%d", fromX, fromY, toX, toY)
	panic("one of either the x or y coordinates should have been equal on the ends of a line")
}

type PointScanner struct {
	r   int    // read index
	buf []byte // all bytes to scan
	x   int    // last read token value x
	y   int    // last read token value y
}

func NewPointScanner(bs []byte) *PointScanner {
	return &PointScanner{
		r:   0,
		buf: bs,
		x:   0,
		y:   0,
	}
}

// Reload will load a new []byte into an existing scanner and reset all the state
func (ps *PointScanner) Reload(bs []byte) {
	ps.r = 0
	ps.buf = bs
	ps.x = 0
	ps.y = 0
}

func (ps *PointScanner) Scan() bool {
	for i := ps.r; i < len(ps.buf); i++ {
		if !isDigit(ps.buf[i]) {
			continue
		}
		start := i
		j := i
		for {
			isEOL := func() bool { return j == len(ps.buf) }
			isComma := func() bool { return ps.buf[j] == ',' }
			isSpace := func() bool { return ps.buf[j] == ' ' }
			if isEOL() || isSpace() {
				n, err := strutil.ParseUintBytes(ps.buf[start:j], 10, 32)
				if err != nil {
					panic(err)
				}
				ps.y = int(n)
				ps.r = j + 1
				return true
			}
			if isComma() {
				n, err := strutil.ParseUintBytes(ps.buf[start:j], 10, 32)
				if err != nil {
					panic(err)
				}
				ps.x = int(n)
				start = j + 1
				j++
				continue
			}
			j++
		}
	}
	return false
}

func (ps *PointScanner) Point() (x, y int) {
	if ps.r == 0 {
		panic("can't return a point when nothing has been scanned")
	}
	x = ps.x
	y = ps.y
	return
}

func isDigit(b byte) bool {
	return b >= 48 && b <= 57
}
