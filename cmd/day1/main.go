package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

// TODO: exercise to make this generic and move to pkg
// a sorted []int would have been easier, but we can save some
// performance here since we don't need to reorder the values
// each time, we just need to know which one is the lowest and
// therefore the next number to be evicted
type maxHolder struct {
	values    []int // n highest values, configured on construction
	minValIdx int   // index of current lowest value
}

func newMaxHolder(capacity int) *maxHolder {
	return &maxHolder{
		values:    make([]int, capacity),
		minValIdx: 0,
	}
}

func (m *maxHolder) Add(n int) bool {
	// see if new value is greater than our lowest value and if so, replace it
	if n > m.values[m.minValIdx] {
		m.values[m.minValIdx] = n
		// set the new smallest value
		var min int
		for i, v := range m.values {
			if i == 0 {
				min = v
				m.minValIdx = i
				continue
			}
			if v < min {
				m.minValIdx = i
			}
		}
		return true
	}
	return false
}

func (m *maxHolder) Sum() int {
	sum := 0
	for _, v := range m.values {
		sum += v
	}
	return sum
}

func main() {
	fmt.Printf("DAY 1\n")
	partOneTime := time.Now()
	lines := strings.Split(input, "\n")
	fmt.Printf("Part One:       %s\n", partOne(lines))
	partOneDuration := time.Since(partOneTime)
	fmt.Printf("Execution time: %s\n", partOneDuration.String())
	partTwoTime := time.Now()
	lines = strings.Split(input, "\n")
	fmt.Printf("Part Two:       %s\n", partTwo(lines))
	partTwoDuration := time.Since(partTwoTime)
	fmt.Printf("Execution time: %s\n", partTwoDuration.String())
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
	mh := newMaxHolder(3)
	for _, s := range sums {
		mh.Add(s)
	}
	return strconv.Itoa(mh.Sum())
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
