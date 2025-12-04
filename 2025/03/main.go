package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([]int, error) {
	return nil, nil
}

func maxJoltage(s string) int {
	ints := util.Map([]rune(s), func(r rune) int { return int(r - '0') })
	maxFirst := slices.Max(ints[:len(ints)-1])
	firstPos := slices.Index(ints, maxFirst)
	maxSecond := slices.Max(ints[firstPos+1:])
	return maxFirst*10 + maxSecond
}

func findMax(ints []int, howMany int) int {
	if howMany == 0 {
		return 0
	}
	maxFirst := slices.Max(ints[:len(ints)-howMany+1])
	firstPos := slices.Index(ints, maxFirst)
	rest := findMax(ints[firstPos+1:], howMany-1)

	factor := 1
	for range howMany - 1 {
		factor *= 10
	}
	return factor*maxFirst + rest
}

func part1(banks []string) (int, error) {
	sum := 0

	for _, bank := range banks {
		sum += maxJoltage(bank)
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	sum := 0

	for _, input := range inputs {
		ints := util.Map([]rune(input), func(r rune) int { return int(r - '0') })
		sum += findMax(ints, 12)
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
