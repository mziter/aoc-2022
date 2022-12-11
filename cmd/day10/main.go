package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"strings"

	"github.com/mziter/aoc-2022/pkg/strutil"
)

//go:embed input.txt
var input []byte

type Inst struct{}
type NoOp int

type VideoSystem struct {
	// cpu
	x           int             // register
	cycle       int             // current cycle
	busy        bool            // is the cpu busy executing an instruction
	currentInst int             // current instruction being executed
	execDone    int             // cycle when execution will finish
	iScan       bufio.Scanner   // instruction scanner for next insruction when ready
	writePixels bool            // whether we should write pixels or not
	sb          strings.Builder // builder for storing state of CRT output
}

func New(iReader io.Reader, writePixels bool) *VideoSystem {
	return &VideoSystem{
		x:           1,
		cycle:       1,
		iScan:       *bufio.NewScanner(iReader),
		writePixels: writePixels,
	}
}

func (vs *VideoSystem) PrintScreen() {
	fmt.Println(vs.sb.String())
}

func (vs *VideoSystem) writePixel() {
	position := (vs.cycle - 1) % 40
	if position >= vs.x-1 && position <= vs.x+1 {
		vs.sb.WriteByte('#')
	} else {
		vs.sb.WriteByte('.')
	}
	if position == 39 {
		vs.sb.WriteByte('\n')
	}
}

func (vs *VideoSystem) Tick() bool {
	// if busy see if execution is finished
	// if finished execute and continue, else
	// increment cycle and return

	if !vs.busy {
		ok := vs.executeNext()
		if !ok {
			return false // no more instructions to execute
		}
	}
	if vs.writePixels {
		vs.writePixel()
	}
	if vs.execDone == vs.cycle {
		vs.x += vs.currentInst
		vs.busy = false
	}
	vs.cycle++
	return true
}

func (vs *VideoSystem) executeNext() bool {
	if gotNext := vs.iScan.Scan(); !gotNext {
		return false
	}
	b := vs.iScan.Bytes()
	switch b[0] {
	case 'a': // we have an addx inst
		var num int
		if b[5] == '-' {
			numB, err := strutil.ParseUintBytes(b[6:], 10, 32)
			if err != nil {
				panic(err)
			}
			num = -int(numB)
		} else {
			numB, err := strutil.ParseUintBytes(b[5:], 10, 32)
			if err != nil {
				panic(err)
			}
			num = int(numB)
		}
		vs.busy = true
		vs.currentInst = num
		vs.execDone = vs.cycle + 1
	case 'n': //we have a noop
	default:
		return false
	}
	return true
}

func main() {
	fmt.Println("Day X")
	fmt.Println("Part I: ", partOne(input, false))
	fmt.Println("Part II: ", partOne(input, true))
}

func partOne(input []byte, shouldDraw bool) int {
	r := bytes.NewReader(input)
	vs := New(r, shouldDraw)
	sum := 0
	for vs.Tick() {
		if vs.cycle == 20 || vs.cycle == 60 || vs.cycle == 100 || vs.cycle == 140 || vs.cycle == 180 || vs.cycle == 220 {
			signal := vs.cycle * vs.x
			sum += signal
		}
	}
	if shouldDraw {
		vs.PrintScreen()
	}
	return sum
}
