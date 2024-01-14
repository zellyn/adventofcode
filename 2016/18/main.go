package main

import (
	"fmt"
	"os"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var toTrap = map[string]byte{
	"...": '.',
	"..^": '^',
	".^.": '.',
	".^^": '^',
	"^..": '^',
	"^.^": '.',
	"^^.": '^',
	"^^^": '.',
}

func next(s string) string {
	last := len(s) - 1
	res := make([]byte, last+1)
	res[0] = toTrap["."+s[:2]]
	for i := 1; i < last; i++ {
		res[i] = toTrap[s[i-1:i+2]]
	}
	res[last] = toTrap[s[last-1:]+"."]
	return string(res)
}

func safes(s string) int {
	count := 0
	for _, c := range s {
		if c == '.' {
			count++
		}
	}
	return count
}

func part1(row string, rows int) (int, error) {
	count := 0
	for i := 0; i < rows; i++ {
		count += safes(row)
		row = next(row)
	}
	return count, nil
}

func part2(inputs []string) (int, error) {
	return 42, nil
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
