package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/mziter/aoc-2022/pkg/ioutil"
)

//go:embed input.txt
var input string

// TODO: exercise to make this generic and move to pkg
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
	start := time.Now()
	lines := ioutil.GetLines(input)
	fmt.Printf("DAY 1\n")
	fmt.Printf("Part One:       %s\n", partOne(lines))
	fmt.Printf("Part Two:       %s\n", partTwo(lines))
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration.String())
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
