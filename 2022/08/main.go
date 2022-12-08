package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/fun"
	"github.com/zellyn/adventofcode/geom"
)

func visible(pos geom.Vec2, m map[geom.Vec2]int, min, max geom.Vec2) bool {
	height := m[pos]

	for _, dir := range geom.Dirs4 {
		p := pos
		for {
			p = p.Add(dir)
			if !p.Within(min, max) {
				return true
			}
			if m[p] >= height {
				break
			}
		}
	}
	return false
}

func visibleInDir(pos, dir geom.Vec2, m map[geom.Vec2]int, min, max geom.Vec2) int {
	height := m[pos]
	count := 0
	p := pos
	for {
		p = p.Add(dir)
		if !p.Within(min, max) {
			return count
		}
		h := m[p]
		if h >= height {
			return count + 1
		}
		count++
	}
}

func scenicScore(pos geom.Vec2, m map[geom.Vec2]int, min, max geom.Vec2) int {
	score := 1
	for _, dir := range geom.Dirs4 {
		score *= visibleInDir(pos, dir, m, min, max)
		if score == 0 {
			return 0
		}
	}
	return score
}

func totalVisible(m map[geom.Vec2]int, min, max geom.Vec2) int {
	total := 0
	for pos := range m {
		if visible(pos, m, min, max) {
			total++
		}
	}
	return total
}

func parse(inputs []string) (m map[geom.Vec2]int, min, max geom.Vec2) {
	charMap := charmap.Parse(inputs)
	min, max = charMap.MinMax()
	intMap := fun.MapMapValues(charMap, func(_ geom.Vec2, r rune) int {
		return int(r - '0')
	})
	return intMap, min, max
}

func part1(inputs []string) (int, error) {
	m, min, max := parse(inputs)
	return totalVisible(m, min, max), nil
}

func part2(inputs []string) (int, error) {
	mHeights, min, max := parse(inputs)
	mScores := fun.MapMapValues(mHeights, func(pos geom.Vec2, _ int) int {
		return scenicScore(pos, mHeights, min, max)
	})
	return fun.MapMax(mScores), nil
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
