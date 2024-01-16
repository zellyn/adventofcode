package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/zellyn/adventofcode/interval"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parseInterval(s string) (interval.I, error) {
	ints, err := util.ParseInts(s, "-")
	if err != nil {
		return interval.I{}, err
	}
	return interval.New(ints[0], ints[1]), nil
}

func parse(inputs []string) ([]interval.I, error) {
	return util.MapE(inputs, parseInterval)
}

func coalesced(intervals []interval.I) []interval.I {
	interval.Sort(intervals)

	pos := 0
	for pos < len(intervals)-1 {
		i1, i2 := intervals[pos], intervals[pos+1]
		if i1.Overlaps(i2) || i1.Adjacent(i2) {
			intervals[pos] = i1.Merge(i2)
			intervals = slices.Delete(intervals, pos+1, pos+2)
		} else {
			pos++
		}
	}

	return intervals
}

func part1(inputs []string) (int, error) {
	intervals, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	intervals = coalesced(intervals)

	return intervals[0][1] + 1, nil
}

func part2(inputs []string, last int) (int, error) {
	intervals, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	intervals = coalesced(intervals)

	ll := len(intervals) - 1

	count := 0
	if intervals[0][0] > 0 {
		count += intervals[0][0]
	}

	for i := 0; i < ll; i++ {
		i1, i2 := intervals[i], intervals[i+1]
		count += i2[0] - i1[1] - 1
	}

	if intervals[ll][1] < last {
		count += last - intervals[ll][1]
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
