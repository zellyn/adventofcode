package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type random int64

func (r *random) next() int64 {
	a := *r * 64
	*r ^= a
	*r %= 16777216

	a = *r / 32
	*r ^= a
	*r %= 16777216

	a = *r * 2048
	*r ^= a
	*r %= 16777216

	return int64(*r)
}

func part1(inputs []string) (int64, error) {
	ints, err := util.StringsToInts(inputs)
	if err != nil {
		return 0, nil
	}

	var sum int64
	for _, i := range ints {
		r := random(i)
		for range 2000 {
			r.next()
		}
		sum += int64(r)
	}
	return sum, nil
}

func deltasAndPrices(seed int64, iters int) []geom.Vec2 {
	var res []geom.Vec2
	r := random(seed)
	last := seed % 10
	for range iters {
		this := r.next() % 10
		res = append(res, geom.V2(int(this), int(this-last)))
		last = this
	}
	return res
}

func codePrices(seed int64, iters int) map[[4]int]int {
	res := make(map[[4]int]int, iters)
	dps := deltasAndPrices(seed, iters)
	for i := range len(dps) - 3 {
		key := [4]int{
			dps[i].Y,
			dps[i+1].Y,
			dps[i+2].Y,
			dps[i+3].Y,
		}
		if _, ok := res[key]; !ok {
			res[key] = dps[i+3].X
		}
	}
	return res
}

func updateTotals(seed int64, iters int, sums map[[4]int]int) {
	dps := deltasAndPrices(seed, iters)
	seen := make(map[[4]int]bool, iters)
	for i := range len(dps) - 3 {
		key := [4]int{
			dps[i].Y,
			dps[i+1].Y,
			dps[i+2].Y,
			dps[i+3].Y,
		}
		if !seen[key] {
			seen[key] = true
			sums[key] += dps[i+3].X
		}
	}
}

func sumBananas(code [4]int, allCodePrices []map[[4]int]int) int {
	total := 0

	for _, codePrices := range allCodePrices {
		total += codePrices[code]
	}

	return total
}

func part2(inputs []string) (int, error) {
	ints, err := util.StringsToInts(inputs)
	if err != nil {
		return 0, nil
	}

	sums := make(map[[4]int]int)
	for _, i := range ints {
		updateTotals(int64(i), 2000, sums)
	}

	best := 0
	for _, v := range sums {
		best = max(best, v)
	}

	return best, nil
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
