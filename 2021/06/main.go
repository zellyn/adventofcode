package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([9]int, error) {
	var res [9]int
	ints, err := util.ParseInts(inputs[0], ",")
	if err != nil {
		return res, err
	}

	for _, i := range ints {
		res[i]++
	}
	return res, nil
}

func part1(inputs []string, days int) (int, error) {
	counts, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	for day := 0; day < days; day++ {
		zeros := counts[0]
		for i := 0; i < 8; i++ {
			counts[i] = counts[i+1]
		}
		counts[6] += zeros
		counts[8] = zeros
	}

	return util.Sum(counts[:]), nil
}

func part2(inputs []string) (int, error) {
	return 42, nil
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
