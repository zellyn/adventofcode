package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([]geom.Vec3, error) {
	ints, err := util.ParseLinesOfInts(inputs, ",")
	if err != nil {
		return nil, err
	}
	return util.Map(ints, func(ints []int) geom.Vec3 {
		return geom.V3(ints[0], ints[1], ints[2])
	}), nil
}

type entry struct {
	dist int
	a    int
	b    int
}

func part1(inputs []string, count int) (int, error) {
	vecs, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	entries := make([]entry, 0, (len(vecs)*len(vecs))/2)

	for i, a := range vecs {
		for j := i + 1; j < len(vecs); j++ {
			b := vecs[j]
			entries = append(entries, entry{
				dist: a.Sub(b).MagSq(),
				a:    i,
				b:    j,
			})
		}
	}

	slices.SortFunc(entries, func(a, b entry) int {
		return cmp.Compare(a.dist, b.dist)
	})

	sets := make(map[int][]int)
	which := make([]int, len(vecs))
	for i := range vecs {
		which[i] = i
		sets[i] = []int{i}
	}

	for i := range count {
		e := entries[i]

		whichA := which[e.a]
		whichB := which[e.b]
		if whichA == whichB {
			continue
		}
		which[e.b] = whichA

		for _, i := range sets[whichB] {
			sets[whichA] = append(sets[whichA], i)
			which[i] = whichA
		}

		delete(sets, whichB)
	}

	var sizes []int
	for _, set := range sets {
		sizes = append(sizes, len(set))
	}

	slices.Sort(sizes)
	slices.Reverse(sizes)

	return sizes[0] * sizes[1] * sizes[2], nil
}

func part2(inputs []string) (int, error) {
	vecs, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	entries := make([]entry, 0, (len(vecs)*len(vecs))/2)

	for i, a := range vecs {
		for j := i + 1; j < len(vecs); j++ {
			b := vecs[j]
			entries = append(entries, entry{
				dist: a.Sub(b).MagSq(),
				a:    i,
				b:    j,
			})
		}
	}

	slices.SortFunc(entries, func(a, b entry) int {
		return cmp.Compare(a.dist, b.dist)
	})

	sets := make(map[int][]int)
	which := make([]int, len(vecs))
	for i := range vecs {
		which[i] = i
		sets[i] = []int{i}
	}

	for i := range entries {
		e := entries[i]

		whichA := which[e.a]
		whichB := which[e.b]
		if whichA == whichB {
			continue
		}
		which[e.b] = whichA

		for _, i := range sets[whichB] {
			sets[whichA] = append(sets[whichA], i)
			which[i] = whichA
		}

		delete(sets, whichB)

		if len(sets) == 1 {
			vecA, vecB := vecs[e.a], vecs[e.b]
			return vecA.X * vecB.X, nil
		}
	}

	return 0, fmt.Errorf("something went wrong! never reached one set")
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
