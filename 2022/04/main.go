package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/interval"
	"github.com/zellyn/adventofcode/util"
)

func parseRanges(input string) (interval.I, interval.I) {
	input2 := strings.ReplaceAll(input, "-", ",")
	ints := util.MustParseInts(input2, ",")
	if len(ints) != 4 {
		panic(fmt.Sprintf("weird input parses into %d ints: %q", len(ints), input))
	}
	return interval.New(ints[0], ints[1]), interval.New(ints[2], ints[3])
}

func redundant(input string) bool {
	i1, i2 := parseRanges(input)
	return i1.Contains(i2) || i2.Contains(i1)
}

func part1(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		if redundant(input) {
			sum++
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		i1, i2 := parseRanges(input)
		if i1.Overlaps(i2) {
			sum++
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
