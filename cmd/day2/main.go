package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Shape = byte

type Outcome = byte

const (
	bA = byte('A')
	bB = byte('B')
	bC = byte('C')
	bX = byte('X')
	bY = byte('Y')
	bZ = byte('Z')
)

//go:embed input.txt
var input string

func shapeScore(shape Shape) int {
	switch shape {
	case bX:
		return 1
	case bY:
		return 2
	case bZ:
		return 3
	default:
		panic("Something went wrong")
	}
}

func desiredShape(oppShape, outcome Outcome) Shape {
	if oppShape == bA && outcome == bX {
		return bZ
	}
	if oppShape == bA && outcome == bY {
		return bX
	}
	if oppShape == bA && outcome == bZ {
		return bY
	}
	if oppShape == bB && outcome == bX {
		return bX
	}
	if oppShape == bB && outcome == bY {
		return bY
	}
	if oppShape == bB && outcome == bZ {
		return bZ
	}
	if oppShape == bC && outcome == bX {
		return bY
	}
	if oppShape == bC && outcome == bY {
		return bZ
	}
	if oppShape == bC && outcome == bZ {
		return bX
	}
	panic("Something went wrong")
}

func outcome(oppShape, ourShape Shape) int {
	if oppShape == bA && ourShape == bX {
		return 3
	}
	if oppShape == bA && ourShape == bY {
		return 6
	}
	if oppShape == bA && ourShape == bZ {
		return 0
	}
	if oppShape == bB && ourShape == bX {
		return 0
	}
	if oppShape == bB && ourShape == bY {
		return 3
	}
	if oppShape == bB && ourShape == bZ {
		return 6
	}
	if oppShape == bC && ourShape == bX {
		return 6
	}
	if oppShape == bC && ourShape == bY {
		return 0
	}
	if oppShape == bC && ourShape == bZ {
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
