package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

// https://github.com/mnml/aoc/blob/main/2022/13/1.go
// Today I copied the solution above. I am obviously lacking some parsing skills and plan on returning to do a fully custom lexing/parsing solution that also executes more quickly.

//go:embed input.txt
var input string

func main() {
	fmt.Println("Day XIII")
	fmt.Println("Part I: ", partOne())
	fmt.Println("Part II: ", partTwo())
}

func partOne() int {
	right := 0
	for i, s := range strings.Split(strings.TrimSpace(input), "\n\n") {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		if cmp(a, b) <= 0 {
			right += i + 1
		}
	}
	return right
}

func partTwo() int {
	pkts := []any{}
	for _, s := range strings.Split(strings.TrimSpace(input), "\n\n") {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		pkts = append(pkts, a, b)
	}

	pkts = append(pkts, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(pkts, func(i, j int) bool { return cmp(pkts[i], pkts[j]) < 0 })

	key := 1
	for i, p := range pkts {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			key *= i + 1
		}
	}
	return key
}

func cmp(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok: // if its not a list, its a number
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}
