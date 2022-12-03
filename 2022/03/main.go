package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/stringset"
)

func getSets(input string) (stringset.S, stringset.S) {
	l := len(input) / 2
	return stringset.OfRunes(input[:l]), stringset.OfRunes(input[l:])
}

func score(s string) int {
	r := s[0]
	if r >= 'a' && r <= 'z' {
		return int(r-'a') + 1
	}
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}
	panic(fmt.Sprintf("weird rune to score: '%c'", r))
}

func intersectScore(input string) int {
	s1, s2 := getSets(input)
	i := stringset.Intersect(s1, s2)
	if len(i) != 1 {
		panic(fmt.Sprintf("want len(intersection)==1; got %d, for input %q", len(i), input))
	}
	for k := range i {
		return score(k)
	}
	return 0
}

func intersect3Score(input1, input2, input3 string) int {
	s1, s2, s3 := stringset.OfRunes(input1), stringset.OfRunes(input2), stringset.OfRunes(input3)
	i := stringset.Intersect(s1, s2)
	i = stringset.Intersect(i, s3)
	if len(i) != 1 {
		panic(fmt.Sprintf("want len(intersection)==1; got %d, for input %q", len(i), input))
	}
	for k := range i {
		return score(k)
	}
	return 0
}

func part1(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		sum += intersectScore(input)
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	sum := 0
	for i := 0; i < len(inputs); i += 3 {
		sum += intersect3Score(inputs[i], inputs[i+1], inputs[i+2])
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
