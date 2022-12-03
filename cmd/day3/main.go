package main

import (
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

func partTwo(content []byte) string {
	line := 0
	sum := 0
	lineSeen := make([]int, 52)
	found := false
	for _, b := range content {
		if b == byte('\n') {
			if line == 2 { // last of group: reset
				line = 0
				found = false
				for i := 0; i < 52; i++ {
					lineSeen[i] = 0
				}
				continue
			} else {
				line++
				continue
			}
		}
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
	return strconv.Itoa(sum)
}

func partOne(content []byte) string {
	lengths := lineLengths(content)
	line := 0
	linePos := 0
	sum := 0
	seen := make([]int, 52)
	found := false
	for _, b := range content {
		halfLen := lengths[line] / 2
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
		if b == byte('\n') {
			// reset array for next line
			for i := 0; i < 52; i++ {
				seen[i] = 0
			}
			line++
			linePos = 0
			found = false
			continue
		} else {
			linePos++
		}
	}
	return strconv.Itoa(sum)
}
