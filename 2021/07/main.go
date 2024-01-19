package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/myslices"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func linearCost(ints []int, target int) int {
	sum := 0
	for _, i := range ints {
		sum += abs(i - target)
	}
	return sum
}

func squaredCost(ints []int, target int) int {
	sum := 0

	for _, i := range ints {
		steps := abs(i - target)
		sum += steps * (steps + 1) / 2
	}

	return sum
}

func part1(inputs []string) (int, error) {
	ints, err := util.ParseInts(inputs[0], ",")
	if err != nil {
		return 0, err
	}

	target := myslices.Medianish(ints)

	return linearCost(ints, target), nil
}

func part2(inputs []string) (int, error) {
	ints, err := util.ParseInts(inputs[0], ",")
	if err != nil {
		return 0, err
	}

	least, most := myslices.MinMax(ints)
	res := most * most * len(ints) * len(ints)

	for target := least; target <= most; target++ {
		if cost := squaredCost(ints, target); cost < res {
			res = cost
		}
	}
	return res, nil
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
