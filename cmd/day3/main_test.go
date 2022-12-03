package main

import (
	"reflect"
	"testing"
)

func TestLineLengths(t *testing.T) {
	testString := `vjdieus
cid
c
codjeeeee
`
	contents := []byte(testString)
	got := lineLengths(contents)
	want := []int{7, 3, 1, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("result should have been %v but was %v", want, got)
	}
}

func TestPriority(t *testing.T) {
	tcs := []struct {
		b byte
		p int
	}{
		{'a', 1},
		{'z', 26},
		{'A', 27},
		{'Z', 52},
	}
	for _, tc := range tcs {
		want := tc.p
		got := priority(tc.b)
		if want != got {
			t.Errorf("expected priority to be %d, but was %d", want, got)
		}
	}
}

func TestPartOne(t *testing.T) {
	tcs := []struct {
		testString string
		want       string
	}{
		{
			testString: `PmmdzqPrVvPwwTWBwg
`,
			want: "42",
		},
		{
			testString: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
`,
			want: "54",
		},
	}

	for _, tc := range tcs {
		got := partOne([]byte(tc.testString))
		if tc.want != got {
			t.Errorf("expected sum to be %s, but was %s", tc.want, got)
		}
	}
}

func TestPartTwo(t *testing.T) {
	tcs := []struct {
		testString string
		want       string
	}{
		{
			testString: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
`,
			want: "18",
		},
		{
			testString: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`,
			want: "70",
		},
	}

	for _, tc := range tcs {
		got := partTwo([]byte(tc.testString))
		if tc.want != got {
			t.Errorf("expected sum to be %s, but was %s", tc.want, got)
		}
	}
}

var result string

/*
	func BenchmarkPartOne(b *testing.B) {
		var r string
		for n := 0; n < b.N; n++ {
			// always record the result to prevent
			// the compiler eliminating the function call.
			r = partOne(input)
		}
		// always store the result to a package level variable
		// so the compiler cannot eliminate the Benchmark itself.
		result = r
	}
*/

func BenchmarkPartTwo(b *testing.B) {
	var r string
	for n := 0; n < b.N; n++ {
		// always record the result to prevent
		// the compiler eliminating the function call.
		r = partTwo(input)
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	result = r
}
