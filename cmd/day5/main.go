package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/mziter/aoc-2022/pkg/collections/stack"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(content []byte) string {
	piles := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		piles[i] = make([]byte, 0, 100)
	}

	line := 0
	i := 0
	open := false
	lastSpace := 0
	found := 0
	var stacks []*stack.ArrayStack[byte]
	var instructions [3]int
	for idx, b := range content {
		// read blocks for first 8 lines
		if line < 8 {
			if i == 1 || i == 5 || i == 9 || i == 13 || i == 17 || i == 21 || i == 25 || i == 29 || i == 33 {
				if b != ' ' {
					n := (i - 1) / 4
					piles[n] = append(piles[n], b)
				}
			}
		}
		if line == 10 && i == 0 {
			// creates stacks now
			stacks = make([]*stack.ArrayStack[byte], 9)
			for i := 0; i < 9; i++ {
				stacks[i] = stack.NewArrayStack[byte](100)
			}
			for i := 0; i < 9; i++ {
				len := len(piles[i])
				for j := len - 1; j >= 0; j-- {
					stacks[i].Push(piles[i][j])
				}
			}
		}
		if line >= 10 {
			// loop until we find a space or newline
			if b == ' ' || b == '\n' {
				if !open {
					open = true
					lastSpace = idx
				} else {
					open = false
					i, err := strconv.Atoi(string(content[lastSpace+1 : idx]))
					if err != nil {
						panic(fmt.Sprintf("couldn't parse number '%s'", string(content[lastSpace:idx])))
					}
					instructions[found] = i
					found++
					if b == '\n' {
						line++
						found = 0
						i = 0
						// execute instruction
						// for instruction
						for i := 0; i < instructions[0]; i++ {
							v, err := stacks[instructions[1]-1].Pop()
							if err != nil {
								panic(err)
							}
							stacks[instructions[2]-1].Push(v)
						}
					}
				}
			}
		}
		if b == '\n' {
			line++
			i = 0
		} else {
			i++
		}
	}
	var sb strings.Builder
	for i := 0; i < 9; i++ {
		v, err := stacks[i].Pop()
		if err != nil {
			panic(err)
		}
		sb.WriteByte(v)
	}
	return sb.String()
}

func partTwo(content []byte) string {
	piles := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		piles[i] = make([]byte, 0, 100)
	}

	line := 0
	i := 0
	open := false
	lastSpace := 0
	found := 0
	var stacks []*stack.ArrayStack[byte]
	var instructions [3]int
	for idx, b := range content {
		// read blocks for first 8 lines
		if line < 8 {
			if i == 1 || i == 5 || i == 9 || i == 13 || i == 17 || i == 21 || i == 25 || i == 29 || i == 33 {
				if b != ' ' {
					n := (i - 1) / 4
					piles[n] = append(piles[n], b)
				}
			}
		}
		if line == 10 && i == 0 {
			// creates stacks now
			stacks = make([]*stack.ArrayStack[byte], 9)
			for i := 0; i < 9; i++ {
				stacks[i] = stack.NewArrayStack[byte](150)
			}
			for i := 0; i < 9; i++ {
				len := len(piles[i])
				for j := len - 1; j >= 0; j-- {
					stacks[i].Push(piles[i][j])
				}
			}
		}
		if line >= 10 {
			// loop until we find a space or newline
			if b == ' ' || b == '\n' {
				if !open {
					open = true
					lastSpace = idx
				} else {
					open = false
					i, err := strconv.Atoi(string(content[lastSpace+1 : idx]))
					if err != nil {
						panic(fmt.Sprintf("couldn't parse number '%s'", string(content[lastSpace:idx])))
					}
					instructions[found] = i
					found++
					if b == '\n' {
						line++
						found = 0
						i = 0
						// execute instruction
						// for instruction
						amount := instructions[0]
						from := instructions[1] - 1
						to := instructions[2] - 1
						vs := make([]byte, amount)
						for i := 0; i < amount; i++ {
							v, err := stacks[from].Pop()
							if err != nil {
								panic(err)
							}
							vs[i] = v
						}
						stacks[to].PushRev(vs)
					}
				}
			}
		}
		if b == '\n' {
			line++
			i = 0
		} else {
			i++
		}
	}
	var sb strings.Builder
	for i := 0; i < 9; i++ {
		v, err := stacks[i].Pop()
		if err != nil {
			panic(err)
		}
		sb.WriteByte(v)
	}
	return sb.String()
}
