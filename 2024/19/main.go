package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([]string, []string, error) {
	paras := util.LinesByParagraph(inputs)
	if len(paras) != 2 {
		return nil, nil, fmt.Errorf("want 2 paragraphs; got %d", len(paras))
	}
	towels := strings.Split(paras[0][0], ", ")
	slices.SortFunc(towels, func(a, b string) int {
		return -cmp.Compare(len(a), len(b))
	})

	return towels, paras[1], nil
}

func possible(target string, towels []string) bool {
	if target == "" {
		return true
	}

	for _, towel := range towels {
		if strings.HasPrefix(target, towel) && possible(target[len(towel):], towels) {
			return true
		}
	}

	return false
}

var cache map[string]bool = make(map[string]bool)

func cachedPossible(target string, towels []string) bool {
	if target == "" {
		return true
	}
	if cached, ok := cache[target]; ok {
		return cached
	}

	for _, towel := range towels {
		if strings.HasPrefix(target, towel) && cachedPossible(target[len(towel):], towels) {
			cache[target] = true
			return true
		}
	}

	cache[target] = false
	return false
}

var countCache map[string]int = make(map[string]int)

func cachedCount(target string, towels []string) int {
	if target == "" {
		return 1
	}
	if cached, ok := countCache[target]; ok {
		return cached
	}

	count := 0
	for _, towel := range towels {
		if strings.HasPrefix(target, towel) {
			count += cachedCount(target[len(towel):], towels)
		}
	}

	countCache[target] = count
	return count
}

func part1(inputs []string) (int, error) {
	towels, patterns, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	clear(cache)

	count := 0
	for _, pattern := range patterns {
		if cachedPossible(pattern, towels) {
			count++
		}
	}
	return count, nil
}

func part2(inputs []string) (int, error) {
	towels, patterns, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	clear(countCache)

	count := 0
	for _, pattern := range patterns {
		count += cachedCount(pattern, towels)
	}
	return count, nil
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
