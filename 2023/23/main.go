package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/graph"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var dirMap = map[rune]geom.Vec2{
	'^': geom.N,
	'v': geom.S,
	'>': geom.E,
	'<': geom.W,
}

func getNeighbors(m charmap.M, pos geom.Vec2) []geom.Vec2 {
	at := m[pos]
	if dir, ok := dirMap[at]; ok {
		return []geom.Vec2{pos.Add(dir)}
	}
	res := make([]geom.Vec2, 0, 4)
	for _, nPos := range pos.Neighbors4() {
		nAt := m[nPos]
		if nAt == '#' {
			continue
		}
		res = append(res, nPos)
	}
	return res
}

func getEdgeEnd(m charmap.M, pos, lastPos geom.Vec2, targets map[geom.Vec2]bool) (*geom.Vec2, int) {
	nodePos := lastPos
	steps := 0

	for {
		steps++
		if targets[pos] {
			return &pos, steps
		}
		neighbors := util.Filter(getNeighbors(m, pos), func(pos geom.Vec2) bool { return pos != lastPos })
		if len(neighbors) == 0 {
			return nil, 0
		}
		if len(neighbors) > 1 {
			panic(fmt.Sprintf("Got %d>1 neighbors %v at position %s (last=%s), starting from node at %s\n",
				len(neighbors), neighbors, pos, lastPos, nodePos))
		}
		lastPos, pos = pos, neighbors[0]
	}
}

func toGraph(m charmap.M) *graph.Graph[geom.Vec2] {
	posToNode := make(map[geom.Vec2]*graph.Node[geom.Vec2])
	posToIsNode := make(map[geom.Vec2]bool)
	min, max := m.MinMax()

	var nodes []*graph.Node[geom.Vec2]
	var edges []*graph.Edge[geom.Vec2]

	for pos, char := range m {
		if char == '#' || pos.Y == min.Y || pos.Y == max.Y {
			continue
		}
		nn := getNeighbors(m, pos)
		if char == 'S' || char == 'E' || len(nn) > 2 {
			node := &graph.Node[geom.Vec2]{
				Name:  pos.String(),
				Props: pos,
				Start: char == 'S',
				End:   char == 'E',
			}
			nodes = append(nodes, node)
			posToNode[pos] = node
			posToIsNode[pos] = true
		}
	}
	mm := m.Copy()
	for pos := range posToNode {
		mm[pos] = 'N'
	}

	for nodePos, node := range posToNode {
		for _, edgeStart := range getNeighbors(m, nodePos) {
			edgeEnd, steps := getEdgeEnd(m, edgeStart, nodePos, posToIsNode)
			if edgeEnd != nil {
				edges = append(edges, &graph.Edge[geom.Vec2]{
					From: node,
					To:   posToNode[*edgeEnd],
					Cost: steps,
				})
			}
		}
	}
	g := graph.NewGraph(nodes, edges)
	return g
}

type pathRep struct {
	nodes uint64
	cost  int
}

func (p pathRep) add(edge *graph.Edge[geom.Vec2]) pathRep {
	return pathRep{
		nodes: p.nodes | (1 << edge.To.Index()),
		cost:  p.cost + edge.Cost,
	}
}

func (p pathRep) freeOf(edge *graph.Edge[geom.Vec2]) bool {
	return p.nodes&(1<<edge.To.Index()) == 0
}

func allPaths(g *graph.Graph[geom.Vec2], current string, path pathRep) []int {

	for {
		if g.End(current) {
			return []int{path.cost}
		}
		edges := util.Filter(g.EdgesFrom(current), path.freeOf)

		if len(edges) == 0 {
			return nil
		}

		if len(edges) == 1 {
			path = path.add(edges[0])
			current = edges[0].To.Name
			continue
		}

		var res []int
		for _, edge := range edges {
			res = append(res, allPaths(g, edge.To.Name, path.add(edge))...)
		}
		return res
	}

}

func longestPathFromCharmap(m charmap.M) int {
	_, max := m.MinMax()
	start := geom.Vec2{X: 1, Y: 0}
	end := geom.Vec2{X: max.X - 1, Y: max.Y}
	m[start] = 'S'
	m[start.N()] = '#'
	m[end.S()] = '#'
	m[end] = 'E'

	g := toGraph(m)
	printf("%s\n", g)
	paths := allPaths(g, g.Starts()[0].Name, pathRep{})

	return slices.Max(paths)
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	return longestPathFromCharmap(m), nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	for pos, char := range m {
		if _, ok := dirMap[char]; ok {
			m[pos] = '.'
		}
	}
	return longestPathFromCharmap(m), nil
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
