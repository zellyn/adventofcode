package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/graph"
)

type node struct {
	pos    geom.Vec2
	m      charmap.M
	height int
}

func (n *node) End() bool {
	return n.m[n.pos] == 'E'
}

func (n *node) Key() string {
	return fmt.Sprintf("(%d,%d): '%c'", n.pos.X, n.pos.Y, n.m[n.pos])
}

func (n *node) Neighbors() []graph.CostedNode {
	var result []graph.CostedNode
	for _, pos := range n.pos.Neighbors4() {
		c, ok := n.m[pos]
		if !ok {
			continue
		}
		h := height(c)
		if h > n.height+1 {
			continue
		}
		pos := pos
		result = append(result, graph.CostedNode{
			N: &node{
				pos:    pos,
				m:      n.m,
				height: h,
			},
			Steps: 1,
		})
	}

	return result
}

func height(c rune) int {
	if c == 'S' {
		return 0
	}
	if c == 'E' {
		return 25
	}
	return int(c - 'a')
}

func part1(input []string) (int, error) {
	m := charmap.Parse(input)
	start, ok := m.Find('S')
	if !ok {
		panic("Cannot find 'S' in map")
	}

	startNode := &node{
		pos:    start,
		m:      m,
		height: 0,
	}

	shortestLength, err := graph.Dijkstra(startNode)
	if err != nil {
		return 0, err
	}

	return shortestLength, nil
}

func part2(input []string) (int, error) {
	m := charmap.Parse(input)
	start, ok := m.Find('S')
	if !ok {
		panic("Cannot find 'S' in map")
	}
	m[start] = 'a'

	min := len(m)
	for _, startPos := range m.FindAll('a') {
		startNode := &node{
			pos:    startPos,
			m:      m,
			height: 0,
		}

		shortestLength, err := graph.Dijkstra(startNode)
		if err != nil {
			continue
		}

		if shortestLength < min {
			min = shortestLength
		}
	}

	return min, nil
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
