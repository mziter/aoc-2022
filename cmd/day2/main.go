package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Shape = byte

type Outcome = byte

//go:embed input.txt
var input string

func shapeScore(shape Shape) int {
	switch shape {
	case byte('X'):
		return 1
	case byte('Y'):
		return 2
	case byte('Z'):
		return 3
	default:
		panic("Something went wrong")
	}
}

func desiredShape(oppShape, outcome Outcome) Shape {
	if oppShape == byte('A') && outcome == byte('X') {
		return byte('Z')
	}
	if oppShape == byte('A') && outcome == byte('Y') {
		return byte('X')
	}
	if oppShape == byte('A') && outcome == byte('Z') {
		return byte('Y')
	}
	if oppShape == byte('B') && outcome == byte('X') {
		return byte('X')
	}
	if oppShape == byte('B') && outcome == byte('Y') {
		return byte('Y')
	}
	if oppShape == byte('B') && outcome == byte('Z') {
		return byte('Z')
	}
	if oppShape == byte('C') && outcome == byte('X') {
		return byte('Y')
	}
	if oppShape == byte('C') && outcome == byte('Y') {
		return byte('Z')
	}
	if oppShape == byte('C') && outcome == byte('Z') {
		return byte('X')
	}
	panic("Something went wrong")
}

func outcome(oppShape, ourShape Shape) int {
	if oppShape == byte('A') && ourShape == byte('X') {
		return 3
	}
	if oppShape == byte('A') && ourShape == byte('Y') {
		return 6
	}
	if oppShape == byte('A') && ourShape == byte('Z') {
		return 0
	}
	if oppShape == byte('B') && ourShape == byte('X') {
		return 0
	}
	if oppShape == byte('B') && ourShape == byte('Y') {
		return 3
	}
	if oppShape == byte('B') && ourShape == byte('Z') {
		return 6
	}
	if oppShape == byte('C') && ourShape == byte('X') {
		return 6
	}
	if oppShape == byte('C') && ourShape == byte('Y') {
		return 0
	}
	if oppShape == byte('C') && ourShape == byte('Z') {
		return 3
	}
	panic("Something went wrong")
}

func roundScore(oppShape, ourShape Shape) int {
	return shapeScore(ourShape) + outcome(oppShape, ourShape)
}

func main() {
	fmt.Printf("DAY 2\n")
	lines := strings.Split(input, "\n")
	fmt.Printf("partOne: %s\n", partOne(lines))
	fmt.Printf("partTwo: %s\n", partTwo(lines))
}

func partOne(lines []string) string {
	score := 0
	for _, line := range lines {
		signs := strings.Split(line, " ")
		score += roundScore(signs[0][0], signs[1][0])
	}
	return strconv.Itoa(score)
}

func partTwo(lines []string) string {
	score := 0
	for _, line := range lines {
		signs := strings.Split(line, " ")
		oppShape := signs[0][0]
		desiredOutcome := signs[1][0]
		signToPlay := desiredShape(oppShape, desiredOutcome)
		score += roundScore(oppShape, signToPlay)
	}
	return strconv.Itoa(score)
}
