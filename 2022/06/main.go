package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/stringset"
)

func different(s string, count int, pos int) bool {
	if pos < count {
		return false
	}
	if pos > len(s) {
		return false
	}
	set := stringset.OfRunes(s[pos-count : pos])
	return len(set) == count
}

func firstDifferent(s string, count int) int {
	for pos := count; pos <= len(s); pos++ {
		if different(s, count, pos) {
			return pos
		}
	}
	return -1
}

func part1(input string) (int, error) {
	return firstDifferent(input, 4), nil
}

func part2(input string) (int, error) {
	return firstDifferent(input, 14), nil
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
