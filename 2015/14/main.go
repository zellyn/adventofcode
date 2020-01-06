package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

func distInSeconds(seconds, speed, fly, rest int) int {
	wholes := seconds / (fly + rest)
	part := seconds % (fly + rest)
	if part > fly {
		part = fly
	}
	return wholes*speed*fly + speed*part
}

func distances(seconds int, deer []util.StringsAndInts) map[string]int {
	result := map[string]int{}
	for _, d := range deer {
		result[d.Strings[0]] = distInSeconds(seconds, d.Ints[0], d.Ints[1], d.Ints[2])
	}
	return result
}

func dist(input []string, seconds int) (int, error) {
	parsed, err := util.ParseStringsAndInts(input, 15, []int{0}, []int{3, 6, 13})
	if err != nil {
		return 0, err
	}

	max := 0
	for _, d := range distances(seconds, parsed) {
		if d > max {
			max = d
		}
	}
	return max, nil
}

func points(input []string, seconds int) (int, error) {
	parsed, err := util.ParseStringsAndInts(input, 15, []int{0}, []int{3, 6, 13})
	if err != nil {
		return 0, err
	}
	points := map[string]int{}

	for sec := 1; sec <= seconds; sec++ {
		dists := distances(sec, parsed)
		max := 0
		for _, d := range dists {
			if d > max {
				max = d
			}
		}
		for name, d := range dists {
			if d == max {
				points[name]++
			}
		}
	}

	max := 0

	for _, p := range points {
		if p > max {
			max = p
		}
	}

	return max, nil
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
