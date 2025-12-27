package main

import (
	"fmt"
	"os"
	"strings"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type node struct {
	name     string
	distance int
}

func parse(inputs []string) (map[string][]string, error) {
	res := make(map[string][]string)
	for _, s := range inputs {
		name, others, ok := strings.Cut(s, ": ")
		if !ok {
			return nil, fmt.Errorf("weird input: %q", s)
		}
		res[name] = strings.Split(others, " ")
	}
	return res, nil
}

func count(from, to string, forward map[string][]string) int {
	cache := make(map[string]int)
	cache[to] = 1
	return countHelper(from, forward, cache)
}

func countHelper(name string, forward map[string][]string, cache map[string]int) int {
	if val, ok := cache[name]; ok {
		return val
	}

	total := 0

	for _, child := range forward[name] {
		total += countHelper(child, forward, cache)
	}
	cache[name] = total
	return total
}

func part1(inputs []string) (int, error) {
	forward, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	return count("you", "out", forward), nil
}

func part2(inputs []string) (int, error) {
	forward, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	svrToFft := count("svr", "fft", forward)
	fftToDac := count("fft", "dac", forward)
	dacToOut := count("dac", "out", forward)
	return svrToFft * fftToDac * dacToOut, nil
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
