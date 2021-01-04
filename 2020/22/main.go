package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

func parse(inputs []string) ([]int, []int, error) {
	paras := util.LinesByParagraph(inputs)
	if len(paras) != 2 {
		return nil, nil, fmt.Errorf("want 2 paras; got %d", len(paras))
	}
	if paras[0][0] != "Player 1:" {
		return nil, nil, fmt.Errorf("want first para to begin with %q; got %q", "Player 1:", paras[0][0])
	}
	if paras[1][0] != "Player 2:" {
		return nil, nil, fmt.Errorf("want first para to begin with %q; got %q", "Player 2:", paras[0][0])
	}

	as, err := util.StringsToInts(paras[0][1:])
	if err != nil {
		return nil, nil, fmt.Errorf("weird player 1 input: %w", err)
	}
	bs, err := util.StringsToInts(paras[1][1:])
	if err != nil {
		return nil, nil, fmt.Errorf("weird player 2 input: %w", err)
	}
	return as, bs, nil
}

func step(as, bs []int) ([]int, []int) {
	a, b := as[0], bs[0]
	if a > b {
		as = append(as[1:], a, b)
		bs = bs[1:]
	} else {
		as = as[1:]
		bs = append(bs[1:], b, a)
	}

	return as, bs
}

func score(ints []int) int {
	mul := 0
	l := len(ints)
	for i, val := range ints {
		mul += (l - i) * val
	}
	return mul
}

func part1(inputs []string) (int, error) {
	as, bs, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	for {
		if len(as) == 0 {
			return score(bs), nil
		} else if len(bs) == 0 {
			return score(as), nil
		}
		as, bs = step(as, bs)
	}
}

func recursive(as, bs []int, seen map[string]bool) (bool, []int) {
	for {
		if len(as) == 0 {
			return false, bs
		}
		if len(bs) == 0 {
			return true, as
		}
		key := fmt.Sprintf("%v%v", as, bs)
		if seen[key] {
			return true, as
		}
		seen[key] = true

		a, b := as[0], bs[0]
		aWins := false
		if len(as) > a && len(bs) > b {
			// recursive
			as2 := make([]int, a)
			bs2 := make([]int, b)
			copy(as2, as[1:])
			copy(bs2, bs[1:])
			aWins, _ = recursive(as2, bs2, make(map[string]bool))

		} else {
			aWins = a > b
		}
		if aWins {
			as = append(as[1:], a, b)
			bs = bs[1:]
		} else {
			as = as[1:]
			bs = append(bs[1:], b, a)
		}
	}
}

func part2(inputs []string) (int, error) {
	as, bs, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	_, ints := recursive(as, bs, make(map[string]bool))

	return score(ints), nil
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
