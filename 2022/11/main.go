package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/lists"
	"github.com/zellyn/adventofcode/util"
)

type monkey struct {
	items   []int
	op      [3]string
	divisor int
	targets map[bool]int
	tests   int
}

func parseMonkey(g []string) (*monkey, error) {
	monkey := &monkey{}

	// Starting items
	if !strings.HasPrefix(g[1], "  Starting items: ") {
		return nil, fmt.Errorf("weird input for starting items: %q", g[1])
	}
	ints, err := util.ParseInts(g[1][len("  Starting items: "):], ", ")
	if err != nil {
		return nil, err
	}
	monkey.items = ints

	// Operation
	op := strings.Split(strings.TrimSpace(g[2]), " ")
	if len(op) != 6 {
		return nil, fmt.Errorf("weird operation: %q parses into %#v", g[2], op)
	}
	monkey.op[0], monkey.op[1], monkey.op[2] = op[3], op[4], op[5]

	// Divisor, true target, false target
	var lastInts []int
	for i := 3; i <= 5; i++ {
		parts := strings.Split(g[i], " ")
		if num, err := strconv.Atoi(parts[len(parts)-1]); err != nil {
			return nil, fmt.Errorf("cannot get int from end of line %q: %w", g[i], err)
		} else {
			lastInts = append(lastInts, num)
		}
	}
	monkey.divisor = lastInts[0]
	monkey.targets = map[bool]int{
		true:  lastInts[1],
		false: lastInts[2],
	}
	return monkey, nil
}

func parseMonkeys(input []string) ([]*monkey, error) {
	var result []*monkey
	groups := util.SplitBefore(input, func(s string) bool {
		return strings.HasPrefix(s, "Monkey ")
	})
	for i, g := range groups {
		monkey, err := parseMonkey(g)
		if err != nil {
			return nil, fmt.Errorf("error parsing monkey %d: %w", i, err)
		}
		result = append(result, monkey)
	}
	return result, nil
}

func parseTerm(s string, old int) int {
	if s == "old" {
		return old
	}
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("Weird number in operation: %q", s))
	}
	return num
}

func doMonkey(monkeys []*monkey, i int, mod *int) {
	m := monkeys[i]
	m.tests += len(m.items)
	for _, item := range m.items {
		// perform operation
		var newValue int
		first, op, second := m.op[0], m.op[1], m.op[2]
		left := parseTerm(first, item)
		right := parseTerm(second, item)
		switch op {
		case "+":
			newValue = left + right
		case "-":
			newValue = left - right
		case "*":
			newValue = left * right
		case "/":
			newValue = left / right
		default:
			panic(fmt.Sprintf("Weird operation %q", op))
		}
		// get bored
		if mod == nil {
			newValue /= 3
		} else {
			newValue = newValue % (*mod)
		}
		// test divisibility
		testVal := (newValue % m.divisor) == 0
		target := m.targets[testVal]
		monkeys[target].items = append(monkeys[target].items, newValue)
	}
	m.items = nil
}

func doMonkeys(monkeys []*monkey, mod *int) {
	for i := range monkeys {
		doMonkey(monkeys, i, mod)
	}
}

func part1(input []string) (int, error) {
	monkeys, err := parseMonkeys(input)
	if err != nil {
		return 0, err
	}
	for i := 0; i < 20; i++ {
		doMonkeys(monkeys, nil)
	}

	var tests []int = lists.Map(monkeys, func(m *monkey) int {
		return m.tests
	})
	sort.Ints(tests)
	return tests[len(tests)-1] * tests[len(tests)-2], nil
}

func part2(input []string) (int, error) {
	monkeys, err := parseMonkeys(input)
	if err != nil {
		return 0, err
	}

	mod := 1
	for _, m := range monkeys {
		mod = mod * m.divisor
	}

	for i := 0; i < 10000; i++ {
		doMonkeys(monkeys, &mod)
	}

	var tests []int = lists.Map(monkeys, func(m *monkey) int {
		return m.tests
	})
	sort.Ints(tests)
	return tests[len(tests)-1] * tests[len(tests)-2], nil
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
