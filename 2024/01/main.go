package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([]int, []int, error) {
	ints, err := util.MapE(inputs, util.ParseFieldInts)
	if err != nil {
		return nil, nil, err
	}
	var firsts, seconds []int
	for _, ii := range ints {
		firsts = append(firsts, ii[0])
		seconds = append(seconds, ii[1])
	}
	return firsts, seconds, nil
}

func part1(inputs []string) (int, error) {
	firsts, seconds, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	sort.Ints(firsts)
	sort.Ints(seconds)
	sum := 0
	for i, i1 := range firsts {
		i2 := seconds[i]
		diff := i1 - i2
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	firsts, seconds, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	counts := make(map[int]int)
	for _, num := range seconds {
		counts[num]++
	}
	sum := 0
	for _, num := range firsts {
		sum += num * counts[num]
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
