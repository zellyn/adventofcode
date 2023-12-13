package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func verticalMirror(m charmap.M) int {
	return horizontalMirror(m.Transpose())
}

func countDiffs(s1, s2 string) int {
	diffs := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffs++
		}
	}
	return diffs
}

func isSmudgedHorizontalMirror(lines []string, pos int) bool {
	diffs := 0
	iters := len(lines) - pos
	if pos < iters {
		iters = pos
	}

	for i := 0; i < iters; i++ {
		if lines[pos-i-1] == lines[pos+i] {
			continue
		}
		diffs += countDiffs(lines[pos-i-1], lines[pos+i])
		if diffs > 1 {
			return false
		}
	}

	return diffs == 1
}

func isHorizontalMirror(lines []string, pos int) bool {
	iters := len(lines) - pos
	if pos < iters {
		iters = pos
	}

	for i := 0; i < iters; i++ {
		if lines[pos-i-1] != lines[pos+i] {
			return false
		}
	}
	return true
}

func horizontalMirror(m charmap.M) int {
	lines := m.AsStrings('.')
	for i := 1; i < len(lines); i++ {
		if isHorizontalMirror(lines, i) {
			return i
		}
	}
	return -1
}

func smudgedVerticalMirror(m charmap.M) int {
	return smudgedHorizontalMirror(m.Transpose())
}

func smudgedHorizontalMirror(m charmap.M) int {
	lines := m.AsStrings('.')
	for i := 1; i < len(lines); i++ {
		if isSmudgedHorizontalMirror(lines, i) {
			return i
		}
	}
	return -1

}

func part1(inputs []string) (int, error) {
	sum := 0
	for _, input := range util.LinesByParagraph(inputs) {
		m := charmap.Parse(input)

		if v := verticalMirror(m); v != -1 {
			sum += v
		}
		if h := horizontalMirror(m); h != -1 {
			sum += 100 * h
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	sum := 0
	for _, input := range util.LinesByParagraph(inputs) {
		m := charmap.Parse(input)

		if v := smudgedVerticalMirror(m); v != -1 {
			sum += v
		}
		if h := smudgedHorizontalMirror(m); h != -1 {
			sum += 100 * h
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
