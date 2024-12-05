package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type rule struct {
	before int
	after  int
}

func parse(inputs []string) ([]rule, [][]int, error) {
	ps := util.LinesByParagraph(inputs)
	rules, err := util.MapE(ps[0], func(s string) (rule, error) {
		s1, s2, ok := strings.Cut(s, "|")
		if !ok {
			return rule{}, fmt.Errorf("weird input: %q", s)
		}
		i1, err := strconv.Atoi(s1)
		if err != nil {
			return rule{}, fmt.Errorf("weird input: %q: %v", s, err)
		}
		i2, err := strconv.Atoi(s2)
		if err != nil {
			return rule{}, fmt.Errorf("weird input: %q: %v", s, err)
		}
		return rule{
			before: i1,
			after:  i2,
		}, nil
	})
	if err != nil {
		return nil, nil, err
	}

	updates, err := util.MapE(ps[1], func(s string) ([]int, error) {
		return util.ParseInts(s, ",")
	})
	if err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}

func goodUpdate(update []int, rules []rule) bool {
	indices := make(map[int]int)
	for i, page := range update {
		indices[page] = i + 1
	}

	for _, rule := range rules {
		posBefore := indices[rule.before]
		posAfter := indices[rule.after]
		if posBefore*posAfter == 0 {
			continue
		}
		if posBefore > posAfter {
			return false
		}
	}
	return true
}

func part1(inputs []string) (int, error) {
	rules, updates, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	sum := 0
	for _, update := range updates {
		if goodUpdate(update, rules) {
			sum += update[len(update)/2]
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	rules, updates, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	sum := 0
	for _, update := range updates {
		if goodUpdate(update, rules) {
			continue
		}
		slices.SortFunc(update, func(a, b int) int {
			for _, rule := range rules {
				if rule.before == a && rule.after == b {
					return -1
				}
				if rule.before == b && rule.after == a {
					return 1
				}
			}
			return 0
		})
		sum += update[len(update)/2]
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
