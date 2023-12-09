package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

// var printf = fmt.Printf
var printf = func(string, ...any) {}

func findNext(line []int) int {
	last := len(line) - 1
	for last >= 0 {
		printf("%v (%d)\n", line, last)
		for i := 0; i < last; i++ {
			line[i] = line[i+1] - line[i]
		}
		if line[last-1] == 0 {
			break
		}
		last--
	}
	printf("%v -- last=%d\n", line, last)
	for ; last < len(line); last++ {
		printf("%v (%d)\n", line, last)
		line[last] = line[last] + line[last-1]
	}
	printf("%v -- last=%d\n", line, last)
	return line[last-1]
}

func findPrev(line []int) int {
	last := len(line) - 1
	first := 0
	for ; first < last; first++ {
		printf("%v (%d)\n", line, first)
		for i := last; i > first; i-- {
			line[i] = line[i] - line[i-1]
		}
		if line[last] == 0 {
			break
		}
	}
	printf("%v -- first=%d\n", line, first)

	for ; first >= 0; first-- {
		printf("%v (%d)\n", line, first)
		line[first] = line[first] - line[first+1]
	}
	printf("%v -- first=%d\n", line, first)

	return line[0]
}

func part1(inputs []string) (int, error) {
	printf("here1\n")
	lines, err := util.ParseGrid(inputs)
	if err != nil {
		return 0, err
	}
	printf("here2\n")
	return util.MappedSum(lines, findNext), nil
}

func part2(inputs []string) (int, error) {
	printf("here1\n")
	lines, err := util.ParseGrid(inputs)
	if err != nil {
		return 0, err
	}
	printf("here2\n")
	return util.MappedSum(lines, findPrev), nil
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
