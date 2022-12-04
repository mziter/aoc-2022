package main

import (
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println("DAY 4")
	fmt.Println("Part I:", partOne(input))
	fmt.Println("Part II:", partTwo(input))
}

func isDelim(ch byte) bool {
	return ch == '-' || ch == ',' || ch == '\n'
}
func isEOL(ch byte) bool {
	return ch == '\n'
}

func partOne(content []byte) string {
	count := 0
	valStart := 0
	bounds := make([]int, 4)
	found := 0
	for i, b := range content {
		if isDelim(b) {
			n, err := strconv.Atoi(string(content[valStart:i]))
			if err != nil {
				panic("encountered some bad data")
			}
			bounds[found] = n
			found++
			if isEOL(b) {
				if bounds[0] <= bounds[2] && bounds[1] >= bounds[3] {
					count++
				} else if bounds[2] <= bounds[0] && bounds[3] >= bounds[1] {
					count++
				}
				found = 0
			}
			valStart = i + 1
		}
	}
	return strconv.Itoa(count)
}

func partTwo(content []byte) string {
	count := 0
	valStart := 0
	bounds := make([]int, 4)
	found := 0
	for i, b := range content {
		if isDelim(b) {
			n, err := strconv.Atoi(string(content[valStart:i]))
			if err != nil {
				panic("encountered some bad data")
			}
			bounds[found] = n
			found++
			if isEOL(b) {
				if bounds[2] > bounds[1] || bounds[0] > bounds[3] {
					// don't overlap
				} else {
					count++
				}
				found = 0
			}
			valStart = i + 1
		}
	}
	return strconv.Itoa(count)
}
