package main

import (
	"fmt"
	"sort"
)

type InspectFunc = func(int) int
type ManageFunc = func(int) int

// queue vs slicing strut vs channel
type Monkey struct {
	ka           *KeepAway
	inspectCount int
	items        chan int
	inspect      InspectFunc
	manage       ManageFunc
	test         int
	whenTrue     int // monkey to throw to when test is true
	whenFalse    int // monkey to throw to when test is false
}

func NewMonkey(inspect InspectFunc, test int, manage ManageFunc, whenTrue int, whenFalse int, ka *KeepAway) *Monkey {
	return &Monkey{
		inspectCount: 0,
		items:        make(chan int, 50),
		inspect:      inspect,
		manage:       manage,
		test:         test,
		whenTrue:     whenTrue,
		whenFalse:    whenFalse,
		ka:           ka,
	}
}

func (m *Monkey) Turn() {
	if len(m.items) == 0 {
		return
	}
	for {
		select {
		case next := <-m.items:
			// pop off son!
			// inspect
			worry := m.inspect(next)
			m.inspectCount++
			// manage worry
			worry = m.manage(worry)
			// test and send to other monkey
			if worry%m.test == 0 {
				m.ka.monkeys[m.whenTrue].items <- worry
			} else {
				m.ka.monkeys[m.whenFalse].items <- worry
			}
		default:
			return
		}
	}
}

func main() {
	fmt.Println("Part I:", partOne())
	fmt.Println("Part II:", partTwo())
}

func partOne() int {
	ka := NewKeepAway()
	return ka.Play(true)
}

func partTwo() int {
	ka := NewKeepAway()
	return ka.Play(false)
}

type KeepAway struct {
	monkeys []*Monkey
}

func NewKeepAway() *KeepAway {
	ka := &KeepAway{
		monkeys: make([]*Monkey, 8),
	}
	ka.monkeys[0] = ka.monkeyZero()
	ka.monkeys[1] = ka.monkeyOne()
	ka.monkeys[2] = ka.monkeyTwo()
	ka.monkeys[3] = ka.monkeyThree()
	ka.monkeys[4] = ka.monkeyFour()
	ka.monkeys[5] = ka.monkeyFive()
	ka.monkeys[6] = ka.monkeySix()
	ka.monkeys[7] = ka.monkeySeven()
	return ka
}

func (ka *KeepAway) Play(isPartOne bool) int {
	if !isPartOne {
		lcm := 1
		for _, m := range ka.monkeys {
			lcm *= m.test
		}
		for _, m := range ka.monkeys {
			m.manage = func(n int) int {
				return n % lcm
			}
		}
	}
	var rounds int
	if isPartOne {
		rounds = 20
	} else {
		rounds = 10000
	}
	for i := 0; i < rounds; i++ {
		for _, m := range ka.monkeys {
			m.Turn()
		}
	}
	sort.Slice(ka.monkeys, func(i, j int) bool {
		return ka.monkeys[i].inspectCount > ka.monkeys[j].inspectCount
	})
	return ka.monkeys[0].inspectCount * ka.monkeys[1].inspectCount
}

// Monkey 0:
//
//	Starting items: 83, 62, 93
//	Operation: new = old * 17
//	Test: divisible by 2
//	  If true: throw to monkey 1
//	  If false: throw to monkey 6
func (ka *KeepAway) monkeyZero() *Monkey {
	monkey := NewMonkey(func(n int) int { return n * 17 },
		2,
		func(n int) int { return n / 3 },
		1,
		6,
		ka,
	)
	monkey.items <- 83
	monkey.items <- 62
	monkey.items <- 93
	return monkey
}

// Monkey 1:
//
//	Starting items: 90, 55
//	Operation: new = old + 1
//	Test: divisible by 17
//	  If true: throw to monkey 6
//	  If false: throw to monkey 3
func (ka *KeepAway) monkeyOne() *Monkey {
	monkey := NewMonkey(func(n int) int { return n + 1 },
		17,
		func(n int) int { return n / 3 },
		6,
		3,
		ka,
	)
	monkey.items <- 90
	monkey.items <- 55
	return monkey
}

// Monkey 2:
//
//	Starting items: 91, 78, 80, 97, 79, 88
//	Operation: new = old + 3
//	Test: divisible by 19
//	  If true: throw to monkey 7
//	  If false: throw to monkey 5
func (ka *KeepAway) monkeyTwo() *Monkey {
	monkey := NewMonkey(func(n int) int { return n + 3 },
		19,
		func(n int) int { return n / 3 },
		7,
		5,
		ka,
	)
	monkey.items <- 91
	monkey.items <- 78
	monkey.items <- 80
	monkey.items <- 97
	monkey.items <- 79
	monkey.items <- 88
	return monkey
}

// Monkey 3:
//
//	Starting items: 64, 80, 83, 89, 59
//	Operation: new = old + 5
//	Test: divisible by 3
//	  If true: throw to monkey 7
//	  If false: throw to monkey 2
func (ka *KeepAway) monkeyThree() *Monkey {
	monkey := NewMonkey(func(n int) int { return n + 5 },
		3,
		func(n int) int { return n / 3 },
		7,
		2,
		ka,
	)
	monkey.items <- 64
	monkey.items <- 80
	monkey.items <- 83
	monkey.items <- 89
	monkey.items <- 59
	return monkey
}

// Monkey 4:
//
//	Starting items: 98, 92, 99, 51
//	Operation: new = old * old
//	Test: divisible by 5
//	  If true: throw to monkey 0
//	  If false: throw to monkey 1
func (ka *KeepAway) monkeyFour() *Monkey {
	monkey := NewMonkey(func(n int) int { return n * n },
		5,
		func(n int) int { return n / 3 },
		0,
		1,
		ka,
	)
	monkey.items <- 98
	monkey.items <- 92
	monkey.items <- 99
	monkey.items <- 51
	return monkey
}

// Monkey 5:
//
//	Starting items: 68, 57, 95, 85, 98, 75, 98, 75
//	Operation: new = old + 2
//	Test: divisible by 13
//	  If true: throw to monkey 4
//	  If false: throw to monkey 0
func (ka *KeepAway) monkeyFive() *Monkey {
	monkey := NewMonkey(func(n int) int { return n + 2 },
		13,
		func(n int) int { return n / 3 },
		4,
		0,
		ka,
	)
	monkey.items <- 68
	monkey.items <- 57
	monkey.items <- 95
	monkey.items <- 85
	monkey.items <- 98
	monkey.items <- 75
	monkey.items <- 98
	monkey.items <- 75
	return monkey
}

// Monkey 6:
//
//	Starting items: 74
//	Operation: new = old + 4
//	Test: divisible by 7
//	  If true: throw to monkey 3
//	  If false: throw to monkey 2
func (ka *KeepAway) monkeySix() *Monkey {
	monkey := NewMonkey(func(n int) int { return n + 4 },
		7,
		func(n int) int { return n / 3 },
		3,
		2,
		ka,
	)
	monkey.items <- 74
	return monkey
}

// Monkey 7:
//
//	Starting items: 68, 64, 60, 68, 87, 80, 82
//	Operation: new = old * 19
//	Test: divisible by 11
//	  If true: throw to monkey 4
//	  If false: throw to monkey 5
func (ka *KeepAway) monkeySeven() *Monkey {
	monkey := NewMonkey(func(n int) int { return n * 19 },
		11,
		func(n int) int { return n / 3 },
		4,
		5,
		ka,
	)
	monkey.items <- 68
	monkey.items <- 64
	monkey.items <- 60
	monkey.items <- 68
	monkey.items <- 87
	monkey.items <- 80
	monkey.items <- 82
	return monkey
}
