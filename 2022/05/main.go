package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

func parseOneLineOfStacks(input string) []string {
	var result []string
	for i := 0; i < len(input); i += 4 {
		switch input[i] {
		case '[':
			result = append(result, input[i+1:i+2])
		case ' ':
			result = append(result, "")
		default:
			panic(fmt.Sprintf("weird input for stacks: %q", input))
		}
	}
	return result
}

func parseStacks(input string) [][]string {
	lines := strings.Split(input, "\n")
	ll := len(lines)
	count := (len(lines[ll-1]) + 1) / 4
	result := make([][]string, count)
	for i := ll - 2; i >= 0; i-- {
		stacks := parseOneLineOfStacks(lines[i])
		for i, stack := range stacks {
			if stack != "" {
				result[i] = append(result[i], stack)
			}
		}
	}
	return result
}

func parseMoves(input string) [][]int {
	var result [][]int
	lines := strings.Split(input, "\n")
	sis, err := util.ParseStringsAndInts(lines, 6, []int{0, 2, 4}, []int{1, 3, 5})
	if err != nil {
		panic(err)
	}

	for _, si := range sis {
		result = append(result, si.Ints)
	}
	return result
}

func parseInput(input string) (stacks [][]string, moves [][]int) {
	parts := strings.Split(input, "\n\n")
	if len(parts) != 2 {
		panic("Weird input: cannot split on double newline into two parts")
	}
	stacks = parseStacks(parts[0])
	moves = parseMoves(parts[1])
	return stacks, moves
}

func performMove(stacks [][]string, move []int) [][]string {
	count, from, to := move[0], move[1]-1, move[2]-1
	lfrom := len(stacks[from])

	for i := 0; i < count; i++ {
		lfrom--
		stacks[to] = append(stacks[to], stacks[from][lfrom])
		stacks[from] = stacks[from][:lfrom]
	}
	return stacks
}

func perform9001Move(stacks [][]string, move []int) [][]string {
	count, from, to := move[0], move[1]-1, move[2]-1
	lfrom := len(stacks[from])

	stacks[to] = append(stacks[to], stacks[from][lfrom-count:]...)
	stacks[from] = stacks[from][:lfrom-count]
	return stacks
}

func part1(input string) (string, error) {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		stacks = performMove(stacks, move)
	}
	result := ""
	for _, s := range stacks {
		if len(s) > 0 {
			result += s[len(s)-1]
		}
	}
	return result, nil
}

func part2(input string) (string, error) {
	stacks, moves := parseInput(input)
	for _, move := range moves {
		stacks = perform9001Move(stacks, move)
	}
	result := ""
	for _, s := range stacks {
		if len(s) > 0 {
			result += s[len(s)-1]
		}
	}
	return result, nil
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
