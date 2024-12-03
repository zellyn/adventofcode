package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func sgn(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func safe(ints []int) bool {
	var diffs []int
	for i := range len(ints) - 1 {
		diffs = append(diffs, ints[i+1]-ints[i])
	}

	sign := sgn(diffs[0])
	if sign == 0 {
		return false
	}

	for _, diff := range diffs {
		if sgn(diff) != sign {
			return false
		}
		if diff < -3 || diff > 3 {
			return false
		}
	}
	return true
}

func anySafe(iints [][]int) bool {
	for _, ints := range iints {
		if safe(ints) {
			return true
		}
	}
	return false
}

func vary(ints []int) [][]int {
	var result [][]int
	result = append(result, ints)
	for i := range ints {
		result = append(result, slices.Delete(slices.Clone(ints), i, i+1))
	}
	return result
}

func part1(inputs []string) (int, error) {
	ints, err := util.MapE(inputs, util.ParseFieldInts)
	if err != nil {
		return 0, err
	}
	safes := util.Map(ints, safe)
	count := 0
	for _, s := range safes {
		if s {
			count++
		}
	}

	return count, nil
}

func part2(inputs []string) (int, error) {
	ints, err := util.MapE(inputs, util.ParseFieldInts)
	if err != nil {
		return 0, err
	}
	variations := util.Map(ints, vary)
	safes := util.Map(variations, anySafe)
	count := 0
	for _, s := range safes {
		if s {
			count++
		}
	}

	return count, nil
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
