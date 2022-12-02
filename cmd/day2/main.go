package main

import (
	_ "embed"
	"fmt"
	"strconv"
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
		panic(fmt.Sprintf("Something went wrong, received shape=%s", string(shape)))
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
	fmt.Printf("partOne: %s\n", partOne(input))
	fmt.Printf("partTwo: %s\n", partTwo(input))
}

func partOne(content string) string {
	score := 0
	for i, c := range content {
		if c == '\n' {
			score += roundScore(content[i-3], content[i-1])
		}
		if i == len(content)-1 {
			score += roundScore(content[i-2], content[i])
		}
	}
	return strconv.Itoa(score)
}

func partTwo(content string) string {
	score := 0
	for i, c := range content {
		if c == '\n' {
			oppShape := content[i-3]
			desiredOutcome := content[i-1]
			signToPlay := desiredShape(oppShape, desiredOutcome)
			score += roundScore(oppShape, signToPlay)
		}
		if i == len(content)-1 {
			oppShape := content[i-2]
			desiredOutcome := content[i]
			signToPlay := desiredShape(oppShape, desiredOutcome)
			score += roundScore(oppShape, signToPlay)
		}
	}
	return strconv.Itoa(score)
}
