package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	lines := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	want := "24000"
	have := partOne(lines)
	if want != have {
		t.Errorf("Wanted answer of %s, but had answer %s", want, have)
	}
}

func TestPartTwo(t *testing.T) {
	lines := []string{
		"1000",
		"2000",
		"3000",
		"",
		"4000",
		"",
		"5000",
		"6000",
		"",
		"7000",
		"8000",
		"9000",
		"",
		"10000",
	}
	want := "45000"
	have := partTwo(lines)
	if want != have {
		t.Errorf("Wanted answer of %s, but had answer %s", want, have)
	}
}
