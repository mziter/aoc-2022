package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/mziter/aoc-2022/pkg/collections/ringbuf"
)

//go:embed input.txt
var input []byte

func main() {
	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partTwo(content []byte) string {
	rb := ringbuf.New[byte](14)
	// we know domain of input is lowercase alphabet only
	// we will keep index zero base by subtracting 'a' ascii value of 97
	var freq [26]byte
	for i, b := range content {
		old := rb.Add(b)
		freq[b-97]++
		if i >= 14 {
			freq[old-97]--
			dup := false
			for i := 0; i < 26; i++ {
				if freq[i] >= 2 {
					dup = true
					break
				}
			}
			if !dup {
				return strconv.Itoa(i + 1)
			}
		}
	}
	return "nothing found, something went wrong"
}

func partOne(content []byte) string {
	rb := ringbuf.New[byte](4)
LOOP:
	for i, b := range content {
		rb.Add(b)
		if i >= 3 {
			for i := 0; i < 4; i++ {
				for j := 1; j < 4; j++ {
					if i != j {
						if rb.Data[i] == rb.Data[j] {
							// break out of nested loop
							continue LOOP
						}
					}
				}
			}
			return strconv.Itoa(i + 1)
		}
	}
	return "nothing found, something went wrong"
}
