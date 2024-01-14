package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type disc struct {
	position int
	size     int
	start    int
}

func (d disc) waitAfter(t int) int {
	pos := (d.start + t) % d.size
	return (d.size - pos) % d.size
}

func parse(inputs []string) ([]disc, error) {
	res := make([]disc, 0, len(inputs))
	for i, ints := range util.ParseRegexInts(inputs, false) {
		if len(ints) != 4 {
			return nil, fmt.Errorf("weird input with %d inputs, not 4: %q", len(ints), inputs[i])
		}
		res = append(res, disc{position: i + 1, size: ints[1], start: ints[3]})
	}

	return res, nil
}

func cmp(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func sortDiscs(discs []disc) {
	slices.SortFunc(discs, func(a, b disc) int {
		return -cmp(a.size, b.size)
	})
}

func findFirstDrop(discs []disc) int {
	d0 := discs[0]
	rest := discs[1:]
	t0 := d0.waitAfter(d0.position)

OUTER:
	for t := t0; ; t += d0.size {
		for _, d := range rest {
			if d.waitAfter(d.position+t) != 0 {
				continue OUTER
			}
		}
		return t
	}
}

func part1(inputs []string) (int, error) {
	discs, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	sortDiscs(discs)

	return findFirstDrop(discs), nil
}

func part2(inputs []string) (int, error) {
	discs, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	discs = append(discs, disc{position: len(discs) + 1, size: 11})
	sortDiscs(discs)

	return findFirstDrop(discs), nil
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
