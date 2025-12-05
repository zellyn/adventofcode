package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type interval [2]int

func parse(inputs []string) ([]interval, []int, error) {

	paras := util.LinesByParagraph(inputs)

	intervals, err := util.MapE(paras[0], func(s string) (interval, error) {
		parts := strings.Split(s, "-")
		ii, err := util.MapE(parts, strconv.Atoi)
		if err != nil {
			return interval{}, err
		}
		return interval{ii[0], ii[1]}, nil
	})
	if err != nil {
		return nil, nil, err
	}

	ids, err := util.MapE(paras[1], strconv.Atoi)
	if err != nil {
		return nil, nil, err
	}

	return intervals, ids, nil
}

func overlap(a, b interval) bool {
	if b[1] < a[0] {
		return false
	}
	if a[1] < b[0] {
		return false
	}
	return true
}

func combine(a, b interval) interval {
	return interval{min(a[0], b[0]), max(a[1], b[1])}
}

func combineAll(intervals []interval) []interval {
	slices.SortFunc(intervals, func(a, b interval) int {
		c := cmp.Compare(a[0], b[0])
		if c == 0 {
			c = cmp.Compare(a[1], b[1])
		}
		return c
	})

	result := intervals[:1]
	last := 0

	for _, i := range intervals[1:] {
		if overlap(i, result[last]) {
			result[last] = combine(i, result[last])
		} else {
			result = append(result, i)
			last++
		}
	}

	return result
}

func part1(inputs []string) (int, error) {
	intervals, ids, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	intervals = combineAll(intervals)
	slices.Sort(ids)

	count := 0

	for _, id := range ids {
		for len(intervals) > 0 && id > intervals[0][1] {
			intervals = intervals[1:]
		}
		if len(intervals) == 0 {
			break
		}
		if id < intervals[0][0] {
			continue
		}
		count++
	}

	return count, nil
}

func part2(inputs []string) (int, error) {
	intervals, _, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	intervals = combineAll(intervals)

	count := 0
	for _, interval := range intervals {
		count += interval[1] - interval[0] + 1
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
