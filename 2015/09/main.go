package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/dgraph"
)

type node struct {
	name  string
	path  string
	graph map[string]map[string]int
}

var _ dgraph.Node = node{}

func (n node) End() bool {
	return strings.Count(n.path, ".") == len(n.graph)-1
}

func (n node) Neighbors() []dgraph.CostedNode {
	var result []dgraph.CostedNode
	if n.name == "" {
		for nn := range n.graph {
			result = append(result, dgraph.CostedNode{
				Steps: 0,
				N: node{
					name:  nn,
					path:  nn,
					graph: n.graph,
				},
			})
		}
		return result
	}
	for nn, dist := range n.graph[n.name] {
		if strings.Contains(n.path, nn) {
			continue
		}
		result = append(result, dgraph.CostedNode{
			Steps: dist,
			N: node{
				name:  nn,
				path:  n.path + "." + nn,
				graph: n.graph,
			},
		})
	}
	return result
}

func (n node) Key() string {
	return n.path
}

func parseInput(input []string) (map[string]map[string]int, error) {
	result := map[string]map[string]int{}
	for _, line := range input {
		parts := strings.Split(line, " ")
		if len(parts) != 5 {
			return nil, fmt.Errorf("stange input line: %q", line)
		}
		from := parts[0]
		to := parts[2]
		dist, err := strconv.Atoi(parts[4])
		if err != nil {
			return nil, fmt.Errorf("error parsing input line: %q: %v", line, err)
		}
		if result[from] == nil {
			result[from] = map[string]int{}
		}
		if result[to] == nil {
			result[to] = map[string]int{}
		}
		result[from][to] = dist
		result[to][from] = dist
	}
	return result, nil
}

func shortestDistance(input []string) (int, error) {
	g, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	start := node{
		name:  "",
		path:  "",
		graph: g,
	}
	return dgraph.Dijkstra(start)
}

func longestDistance(input []string) (int, error) {
	g, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	var cities []string
	for c := range g {
		cities = append(cities, c)
	}
	perms := dgraph.PermutationsString(cities)
	longest := 0

OUTER:
	for _, perm := range perms {
		length := 0

		for i := 1; i < len(perm); i++ {
			from := perm[i-1]
			to := perm[i]
			if g[from] == nil || g[from][to] == 0 {
				continue OUTER
			}
			length += g[from][to]
		}

		if length > longest {
			longest = length
		}
	}
	return longest, nil
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
