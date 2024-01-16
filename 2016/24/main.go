package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/dgraph"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type item struct {
	pos   geom.Vec2
	steps int
}

type node struct {
	allLengths []map[int]int
	at         int
	seen       string
	target     int
}

func (n node) Key() string {
	return strconv.Itoa(n.at) + "-" + n.seen
}

func (n node) End() bool {
	if len(n.seen) != len(n.allLengths) {
		return false
	}
	if n.target == -1 {
		return true
	}
	return n.target == n.at
}

func (n node) Neighbors() []dgraph.CostedNode {
	stepsTo := n.allLengths[n.at]
	res := make([]dgraph.CostedNode, 0, len(stepsTo))

	for index, steps := range stepsTo {
		res = append(res, dgraph.CostedNode{
			N: node{
				allLengths: n.allLengths,
				at:         index,
				seen:       mergeSeen(n.seen, index),
				target:     n.target,
			},
			Steps: steps})
	}

	return res
}

func mergeSeen(seen string, here int) string {
	s := strconv.Itoa(here)
	for _, c := range seen {
		if c == rune(s[0]) {
			return seen
		}
	}
	return sortString(seen + s)
}

func sortString(s string) string {
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func lengthsFrom(m charmap.M, pos geom.Vec2) map[int]int {
	res := make(map[int]int)
	seen := make(map[geom.Vec2]bool)
	todo := []item{{pos: pos, steps: 0}}

	for len(todo) > 0 {
		this := todo[0]
		todo = todo[1:]
		if seen[this.pos] {
			continue
		}
		seen[this.pos] = true
		if this.pos != pos {
			r := m[this.pos]
			if r >= '0' && r <= '9' {
				res[int(r-'0')] = this.steps
			}
		}

		for _, nPos := range this.pos.Neighbors4() {
			if m[nPos] == '#' || seen[nPos] {
				continue
			}
			todo = append(todo, item{
				pos:   nPos,
				steps: this.steps + 1,
			})
		}
	}

	return res
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	var points []geom.Vec2
	for pos, r := range m {
		if r >= '0' && r <= '9' {
			index := int(r - '0')
			for len(points) < index+1 {
				points = append(points, geom.Vec2{})
			}
			points[index] = pos
		}
	}
	allLengths := util.Map(points, func(pos geom.Vec2) map[int]int {
		return lengthsFrom(m, pos)
	})

	start := node{
		allLengths: allLengths,
		at:         0,
		seen:       "0",
		target:     -1,
	}

	return dgraph.Dijkstra(start)
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	var points []geom.Vec2
	for pos, r := range m {
		if r >= '0' && r <= '9' {
			index := int(r - '0')
			for len(points) < index+1 {
				points = append(points, geom.Vec2{})
			}
			points[index] = pos
		}
	}
	allLengths := util.Map(points, func(pos geom.Vec2) map[int]int {
		return lengthsFrom(m, pos)
	})

	start := node{
		allLengths: allLengths,
		at:         0,
		seen:       "0",
		target:     0,
	}

	return dgraph.Dijkstra(start)
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
