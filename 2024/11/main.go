package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func next(i int) []int {
	if i == 0 {
		return []int{1}
	}
	digits := strconv.Itoa(i)
	if len(digits)%2 == 0 {
		i1, _ := strconv.Atoi(digits[:len(digits)/2])
		i2, _ := strconv.Atoi(digits[len(digits)/2:])
		return []int{i1, i2}
	}
	return []int{i * 2024}
}

func makeCounts(ints []int) map[int]int {
	res := make(map[int]int)
	for _, i := range ints {
		res[i]++
	}
	return res
}

func nextCounts(counts map[int]int) map[int]int {
	res := make(map[int]int)
	for i, count := range counts {
		for _, j := range next(i) {
			res[j] += count
		}
	}
	return res
}

func sum(counts map[int]int) int {
	total := 0
	for _, count := range counts {
		total += count
	}
	return total
}

func part1(input string, iters int) (int, error) {
	ints := util.MustParseInts(input, " ")
	counts := makeCounts(ints)
	for range iters {
		counts = nextCounts(counts)
	}
	return sum(counts), nil
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
