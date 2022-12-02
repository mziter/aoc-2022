package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Could probably improve performance like in Day 2 by iterating through all the input instead
// of allocating new strings. Don't feel like it's worth it at this point.
func main() {
	fmt.Printf("DAY 1\n")
	lines := strings.Split(input, "\n")
	fmt.Printf("Part One:       %s\n", partOne(lines))
	lines = strings.Split(input, "\n")
	fmt.Printf("Part Two:       %s\n", partTwo(lines))
}

func partOne(lines []string) string {
	sums := sumByElf(lines)
	max := 0
	for _, s := range sums {
		if s > max {
			max = s
		}
	}
	return strconv.Itoa(max)
}

func partTwo(lines []string) string {
	sums := sumByElf(lines)
	top := make([]int, 3)
	minIdx := 0
	for _, s := range sums {
		if s > top[minIdx] {
			top[minIdx] = s
			// find new min
			min := top[0]
			for i := 0; i < 3; i++ {
				if top[i] <= min {
					minIdx = i
					min = top[i]
				}
			}
		}
	}
	return strconv.Itoa(top[0] + top[1] + top[2])
}

func sumByElf(lines []string) []int {
	var sums []int
	sum := 0
	for i, line := range lines {
		isLast := i == len(lines)-1
		if len(line) > 0 {
			n, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			sum += n
		}
		if len(line) == 0 || isLast {
			sums = append(sums, sum)
			sum = 0
		}
	}
	return sums
}
