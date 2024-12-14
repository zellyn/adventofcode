package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func fill(pos geom.Vec2, group int, m charmap.M, groups map[geom.Vec2]int) (int, int) {
	area := 1
	perimeter := 0
	groups[pos] = group
	for _, nn := range pos.Neighbors4() {
		if m[pos] == m[nn] {
			if groups[nn] == 0 {
				a, p := fill(nn, group, m, groups)
				area += a
				perimeter += p
			}
		} else {
			perimeter++
		}
	}
	return area, perimeter
}

type side struct {
	pos geom.Vec2
	dir geom.Vec2
}

// neighboringSides returns the two sides that would be continuations of the given side
func neighboringSides(s side, m charmap.M) []side {
	result := make([]side, 0, 2)
	for _, neighbor := range []geom.Vec2{s.pos.Add(s.dir.Clockwise90()), s.pos.Add(s.dir.CounterClockwise90())} {
		// Not in same group
		if m[neighbor] != m[s.pos] {
			continue
		}
		// No fence there
		if m[neighbor] == m[neighbor.Add(s.dir)] {
			continue
		}
		result = append(result, side{pos: neighbor, dir: s.dir})
	}
	return result
}

func mergeInto(targetGroup int, sourceGroup int, sideGroups map[side]int) {
	for pos, group := range sideGroups {
		if group == sourceGroup {
			sideGroups[pos] = targetGroup
		}
	}
}

func countSides(group int, m charmap.M, groups map[geom.Vec2]int) int {
	nextGroup := 1
	sideGroups := make(map[side]int)
	sideCount := 0

	for pos, thisGroup := range groups {
		if group != thisGroup {
			continue
		}

		for _, dir := range geom.Dirs4 {
			s := side{pos: pos, dir: dir}
			if groups[pos] == groups[pos.Add(dir)] {
				continue
			}
			sideGroup := sideGroups[s]
			if sideGroup == 0 {
				sideGroup = nextGroup
				nextGroup++
				sideGroups[s] = sideGroup
				sideCount++
			}
			for _, ns := range neighboringSides(s, m) {
				og := sideGroups[ns]
				if og == 0 {
					sideGroups[ns] = sideGroup
					continue
				}
				if og == sideGroup {
					continue
				}
				mergeInto(sideGroup, og, sideGroups)
				sideCount--
			}
		}
	}

	return sideCount
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	groups := make(map[geom.Vec2]int, len(m))
	total := 0
	group := 0
	for pos, _ := range m {
		if groups[pos] > 0 {
			continue
		}
		group++
		area, perimeter := fill(pos, group, m, groups)
		total += area * perimeter
	}

	return total, nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	groups := make(map[geom.Vec2]int, len(m))
	total := 0
	group := 0
	for pos := range m {
		if groups[pos] > 0 {
			continue
		}
		group++
		area, _ := fill(pos, group, m, groups)
		sides := countSides(group, m, groups)
		total += area * sides
	}

	return total, nil
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
