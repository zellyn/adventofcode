package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/zellyn/adventofcode/myslices"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var badCloserScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completingCloserScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

type result int

const (
	resultOk = iota
	resultCloserMismatch
	resultExtraCloser
	resultIncomplete
)

type verdict struct {
	badCloser rune
	expected  string
	result    result
}

func stackToString(stack []rune) string {
	stack = slices.Clone(stack)
	slices.Reverse(stack)
	return string(stack)
}

func eval(s string) verdict {
	var wants []rune

	for _, c := range s {
		pair, ok := pairs[c]
		if ok {
			wants = append(wants, pair)
			continue
		}

		if len(wants) == 0 {
			return verdict{
				badCloser: c,
				result:    resultExtraCloser,
			}
		}

		if c == wants[len(wants)-1] {
			wants = wants[:len(wants)-1]
			continue
		}

		return verdict{
			badCloser: c,
			result:    resultCloserMismatch,
			expected:  stackToString(wants),
		}
	}

	if len(wants) == 0 {
		return verdict{
			result: resultOk,
		}
	}

	return verdict{
		result:   resultIncomplete,
		expected: stackToString(wants),
	}
}

func part1(inputs []string) (int, error) {
	score := 0
	for _, input := range inputs {
		v := eval(input)
		if v.result == resultOk || v.result == resultIncomplete {
			continue
		}
		score += badCloserScores[v.badCloser]
	}
	return score, nil
}

func part2(inputs []string) (int, error) {
	var scores []int

	for _, input := range inputs {
		v := eval(input)
		if v.result != resultIncomplete {
			continue
		}
		// printf("%s can be completed with %s ", input, v.expected)
		score := 0
		for _, c := range v.expected {
			score = score*5 + completingCloserScores[c]
		}
		// printf(" for a score of %d\n", score)
		scores = append(scores, score)
	}
	return myslices.Medianish(scores), nil
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
