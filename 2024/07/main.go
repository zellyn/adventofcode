package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type eqn struct {
	target  int
	factors []int
}

func parseLine(input string) (eqn, error) {
	first, rest, ok := strings.Cut(input, ": ")
	if !ok {
		return eqn{}, fmt.Errorf(`weird line (can't find ": "): %q`, input)
	}
	target, err := strconv.Atoi(first)
	if err != nil {
		return eqn{}, fmt.Errorf("weird non-number %q in input %q", first, input)
	}
	ints, err := util.ParseInts(rest, " ")
	if err != nil {
		return eqn{}, fmt.Errorf("error in input %q: %v", input, err)
	}
	return eqn{
		target:  target,
		factors: ints,
	}, nil
}

func parse(inputs []string) ([]eqn, error) {
	return util.MapE(inputs, parseLine)
}

func canMake(target, acc int, rest []int) bool {
	if acc > target {
		return false
	}
	if len(rest) == 0 {
		return target == acc
	}
	return canMake(target, acc*rest[0], rest[1:]) || canMake(target, acc+rest[0], rest[1:])
}

func join(a, b int) int {
	for n := b; n > 0; n /= 10 {
		a = a * 10
	}
	return a + b
}

func canMake2(target, acc int, rest []int) bool {
	if acc > target {
		return false
	}
	if len(rest) == 0 {
		return target == acc
	}
	return canMake2(target, acc*rest[0], rest[1:]) ||
		canMake2(target, acc+rest[0], rest[1:]) ||
		canMake2(target, join(acc, rest[0]), rest[1:])
}

func part1(inputs []string) (int, error) {
	eqns, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, e := range eqns {
		if canMake(e.target, e.factors[0], e.factors[1:]) {
			sum += e.target
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	eqns, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	sum := 0
	for _, e := range eqns {
		if canMake2(e.target, e.factors[0], e.factors[1:]) {
			sum += e.target
		}
	}
	return sum, nil
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
