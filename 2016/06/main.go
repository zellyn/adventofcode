package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

func flip(lines []string) []string {
	var result []string

	for _, line := range lines {
		for len(result) < len(line) {
			result = append(result, "")
		}
		for i, ch := range line {
			result[i] = result[i] + string(ch)
		}
	}
	return result
}

func mapStringToString(strings []string, f func(string) string) []string {
	result := make([]string, 0, len(strings))
	for _, s := range strings {
		result = append(result, f(s))
	}
	return result
}

func mostCommon(s string) string {
	counts := map[rune]int{}
	max := 0
	result := ""
	for _, ch := range s {
		counts[ch]++
		if counts[ch] > max {
			max = counts[ch]
			result = string(ch)
		}
	}
	return result
}

func leastCommon(s string) string {
	counts := map[rune]int{}
	for _, ch := range s {
		counts[ch]++
	}
	min := len(s)
	result := ""
	for ch, count := range counts {
		if count < min {
			min = count
			result = string(ch)
		}
	}
	return result
}

func decode(filename string) (string, string, error) {
	lines, err := util.ReadLines(filename)
	if err != nil {
		return "", "", err
	}

	flipped := flip(lines)
	mostCommons := mapStringToString(flipped, mostCommon)
	leastCommons := mapStringToString(flipped, leastCommon)
	return strings.Join(mostCommons, ""), strings.Join(leastCommons, ""), nil
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
