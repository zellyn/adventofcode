package main

import (
	"errors"
	"fmt"
	"math"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type state struct {
	pos  geom.Vec2
	warp int
	seen map[geom.Vec2]bool
	dir  geom.Vec2
}

func countPaths(m charmap.M, maxSteps int) (int, error) {
	maxSeen := maxSteps + 1
	if maxSteps == math.MaxInt {
		maxSeen = maxSteps
	}
	startPos, ok := m.Find('S')
	if !ok {
		return 0, errors.New("Can't find start ('S')")
	}
	endPos, ok := m.Find('E')
	if !ok {
		return 0, errors.New("Can't find end ('E')")
	}
	start := state{
		pos:  startPos,
		seen: map[geom.Vec2]bool{startPos: true},
		warp: 0,
	}

	todo := []state{start}

	count := 0

	for len(todo) > 0 {
		next := todo[0]
		todo = todo[1:]

		// fmt.Printf(" consideringi %s\n", next.pos)

		// This and everything after are too long.
		if len(next.seen) >= maxSeen {
			break
		}

		if next.warp == 1 {
			todo = append(todo, state{
				pos:  next.pos.Add(next.dir),
				warp: -1,
				seen: util.SetPlus(next.seen, next.pos.Add(next.dir)),
			})
			continue
		}

		for _, nn := range next.pos.Neighbors4() {
			if next.seen[nn] || m[nn] == '#' {
				continue
			}
			if nn == endPos {
				count++
				continue
			}
			todo = append(todo, state{
				pos:  nn,
				seen: util.SetPlus(next.seen, nn),
				warp: next.warp,
			})
		}

		if len(next.seen)+1 < maxSeen && next.warp == 0 {
			for _, dir := range geom.Dirs4 {
				p1 := next.pos.Add(dir)
				p2 := p1.Add(dir)
				if next.seen[p1] || next.seen[p2] {
					continue
				}
				if m[p1] != '#' {
					continue
				}
				if m[p2] == '#' {
					continue
				}
				if p2 == endPos {
					count++
					continue
				}
				todo = append(todo, state{
					pos:  p1,
					seen: util.SetPlus(next.seen, p1),
					warp: 1,
					dir:  dir,
				})
			}
		}
	}

	return count, nil
}

func part1(inputs []string, minSave int) (int, error) {
	m := charmap.Parse(inputs)
	path, err := m.MinPathRunes('S', 'E', ".E")
	if err != nil {
		return 0, err
	}

	pathSteps := make(map[geom.Vec2]int, len(path))
	for i, pos := range path {
		pathSteps[pos] = i + 1
	}

	count := 0

	for _, pos := range path {
		steps := pathSteps[pos]
		for _, dir := range geom.Dirs4 {
			p1 := pos.Add(dir)
			p2 := p1.Add(dir)
			if m[p1] != '#' || m[p2] == '#' {
				continue
			}
			otherSteps := pathSteps[p2]
			if otherSteps == 0 {
				continue
			}
			if otherSteps-steps-2 >= minSave {
				count++
			}
		}
	}

	return count, nil
}

func part2(inputs []string, minSave int) (int, error) {
	m := charmap.Parse(inputs)
	path, err := m.MinPathRunes('S', 'E', ".E")
	if err != nil {
		return 0, err
	}

	pathSteps := make(map[geom.Vec2]int, len(path))
	for i, pos := range path {
		pathSteps[pos] = i + 1
	}

	count := 0

	for i, pos1 := range path {
		for j := i + minSave; j < len(path); j++ {
			pos2 := path[j]
			taxiDistance := pos1.Taxi(pos2)
			if taxiDistance > 20 {
				continue
			}
			normalDiff := j - i
			saved := normalDiff - taxiDistance
			if saved >= minSave {
				count++
			}
		}
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
