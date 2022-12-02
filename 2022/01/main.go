package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/zellyn/adventofcode/util"
)

func elfints(inputs []string) [][]int {
	paras := util.LinesByParagraph(inputs)
	result := make([][]int, 0, len(paras))

	for _, lines := range paras {
		result = append(result, util.MustStringsToInts(lines))
	}
	return result
}

func maxElfSum(elfcals [][]int) int {
	max := 0
	for _, cals := range elfcals {
		sum := util.Sum(cals)
		if sum > max {
			max = sum
		}
	}
	return max
}

func part1(inputs []string) (int, error) {
	cals := elfints(inputs)
	return maxElfSum(cals), nil
}

func part2(inputs []string) (int, error) {
	cals := elfints(inputs)
	totals := util.MapSum(cals)
	sort.Ints(totals)
	l := len(totals)

	return totals[l-1] + totals[l-2] + totals[l-3], nil
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
