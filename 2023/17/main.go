package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/graph"
)

var printf = func(string, ...any) {}

// var printf = fmt.Printf

var opposites = map[string]string{
	"N": "S",
	"S": "N",
	"E": "W",
	"W": "E",
}

type node struct {
	m    charmap.M
	end  geom.Vec2
	pos  geom.Vec2
	last string
	min  int
	max  int
}

var _ graph.Node = node{}

func (n node) lastCount() (string, int) {
	if n.last == "" {
		return "X", 0
	}
	count := 1
	for ; count < len(n.last) && n.last[count] == n.last[count-1]; count++ {
	}
	return n.last[:1], count
}

func (n node) Key() string {
	dir, count := n.lastCount()
	return n.pos.String() + strconv.Itoa(count) + dir
}

func (n node) End() bool {
	return n.pos == n.end
}

func (n node) String() string {
	return "[" + n.pos.String() + " " + n.last + "]"
}

func (n node) Neighbors() []graph.CostedNode {
	printf(" Neighbors of node %s:\n", n.Key())
	var result []graph.CostedNode

	lastName, lastCount := n.lastCount()

	for name, dir := range geom.Compass4 {
		printf("  considering %s:\n", name)
		// No more than max in a row
		if lastName == name && lastCount == n.max {
			continue
		}
		// No u-turns
		oppositeName, validDir := opposites[lastName]
		printf("   lastName=%q, oppositeName=%q, name=%q\n", lastName, oppositeName, name)
		if oppositeName == name {
			printf("   would be u-turn\n")
			continue
		}
		if validDir && lastCount < n.min && name != lastName {
			continue
		}
		if name != lastName {
			// Check if we can even go that far in this direction
			_, ok := n.m[n.pos.Add(dir.Mul(n.min))]
			if !ok {
				continue
			}
		}
		newPos := n.pos.Add(dir)
		runeCost, ok := n.m[newPos]
		if !ok {
			continue
		}
		cost := int(runeCost - '0')
		nn := n
		nn.pos = newPos
		nn.last = name + n.last
		result = append(result, graph.CostedNode{
			N:     nn,
			Steps: cost,
		})
	}

	printf(" neighbors = %v\n", result)

	return result
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	min, max := m.MinMax()

	start := node{
		m:   m,
		end: max,
		pos: min,
		min: 1,
		max: 3,
	}

	return graph.Dijkstra(start)
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	min, max := m.MinMax()

	start := node{
		m:   m,
		end: max,
		pos: min,
		min: 4,
		max: 10,
	}

	return graph.Dijkstra(start)
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
