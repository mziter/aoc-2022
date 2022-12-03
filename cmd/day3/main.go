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
	fmt.Println("DAY 3")
	fmt.Println("Part I:  ", partOne(input))
	fmt.Println("Part II: ", partTwo(input))
	fmt.Println("Part II(async): ", partTwoAsync(input))
}

func priority(b byte) int {
	i := int(b)
	if i >= 97 { // is lowercase, higher ascii, but lower priority
		return i - 96
	}
	return i - 38 // otherwise we know its uppercase, and use different math
}

func lineLengths(content []byte) []int {
	var lengths []int
	l := 0
	for _, b := range content {
		if b == byte('\n') {
			if l > 0 {
				lengths = append(lengths, l)
				l = 0
			}
		} else {
			l++
		}
	}
	return lengths
}

func partTwoAsync(content []byte) string {
	lines := bytes.Split(content, []byte{'\n'})
	nGroups := len(lines) / 3
	resultC := make(chan int, nGroups)
	defer close(resultC)
	for i := 0; i < len(lines)-3; i = i + 3 {
		go partTwoWork(lines[i:i+4], resultC)
	}
	nResults := 0
	sum := 0
	for {
		for r := range resultC {
			nResults++
			sum += r
			if nResults == nGroups {
				return strconv.Itoa(sum)
			}
		}
	}
}

func partTwoWork(lines [][]byte, resultC chan int) {
	line := 0
	lineSeen := make([]int, 52)
	for _, l := range lines {
		for _, b := range l {
			p := priority(b)
			seenLastLine := lineSeen[p-1]
			if line == 2 && seenLastLine == 2 {
				resultC <- p
				return
			}
			if seenLastLine == line {
				lineSeen[p-1] = line + 1
			}
		}
		line++
	}
}

func partTwo(content []byte) string {
	lines := bytes.Split(content, []byte{'\n'})
	line := 0
	sum := 0
	lineSeen := make([]int, 52)
	found := false
	for _, l := range lines {
		for _, b := range l {
			if !found {
				p := priority(b)
				seenLastLine := lineSeen[p-1]
				if line == 2 && seenLastLine == 2 {
					sum += p
					found = true
				}
				if seenLastLine == line {
					lineSeen[p-1] = line + 1
				}
			}
		}
		if line == 2 { // last of group: reset
			line = 0
			found = false
			for i := 0; i < 52; i++ {
				lineSeen[i] = 0
			}
			continue
		}
		line++
	}
	return strconv.Itoa(sum)
}

func partOne(content []byte) string {
	lines := bytes.Split(content, []byte{'\n'})
	sum := 0
	seen := make([]int, 52)
	found := false
	for _, l := range lines {
		halfLen := len(l) / 2
		for linePos, b := range l {
			// first compartment
			if linePos < halfLen {
				p := priority(b)
				seen[p-1] = 1
			}
			// second compartment
			if linePos >= halfLen && !found {
				p := priority(b)
				if seen[p-1] == 1 {
					sum += p
					found = true
				}
			}
			linePos++
		}
		for i := 0; i < 52; i++ {
			seen[i] = 0
		}
		found = false
	}
	return strconv.Itoa(sum)
}
