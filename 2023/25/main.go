package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"

	"golang.org/x/exp/maps"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parseGraph(inputs []string) ([]string, map[string][]string, error) {
	neighbors := make(map[string][]string)

	add := func(a, b string) {
		if !slices.Contains(neighbors[a], b) {
			neighbors[a] = append(neighbors[a], b)
		}
		if !slices.Contains(neighbors[b], a) {
			neighbors[b] = append(neighbors[b], a)
		}
	}

	for _, input := range inputs {
		node, neighbors, ok := strings.Cut(input, ": ")
		if !ok {
			return nil, nil, fmt.Errorf("weird input line: %q", input)
		}

		for _, neighbor := range strings.Split(neighbors, " ") {
			add(node, neighbor)
		}
	}

	nodes := maps.Keys(neighbors)
	sort.Strings(nodes)

	return nodes, neighbors, nil
}

func printDot(nodes []string, g map[string][]string, omit ...edge) {

	omitMap := make(map[edge]bool)
	for _, e := range omit {
		omitMap[e] = true
		omitMap[edge{e[1], e[0]}] = true
	}

	printf("graph {\n")

	for _, node := range nodes {
		for _, neighbor := range g[node] {
			if omitMap[edge{node, neighbor}] {
				continue
			}
			if neighbor > node {
				printf("  %s -- %s\n", node, neighbor)
			}
		}
	}

	printf("}\n")
}

type edge [2]string

func (e edge) rev() edge {
	return edge{e[1], e[0]}
}

func (e edge) addToMap(m map[edge]bool) {
	m[e] = true
	m[e.rev()] = true
}

func connected(start string, g map[string][]string, omit map[edge]bool) []string {
	var res []string
	visited := make(map[string]bool)
	todo := []string{start}

	for len(todo) > 0 {
		from := todo[0]
		todo = todo[1:]
		if visited[from] {
			continue
		}
		visited[from] = true
		res = append(res, from)
		for _, to := range g[from] {
			if visited[to] || omit[edge{from, to}] {
				continue
			}
			todo = append(todo, to)
		}
	}

	sort.Strings(res)
	return res
}

func pathBetween(start, end string, g map[string][]string, omit map[edge]bool) []string {
	visited := map[string]bool{start: true}
	todo := [][]string{{start}}

	for len(todo) > 0 {
		soFar := todo[0]
		todo = todo[1:]
		last := soFar[len(soFar)-1]
		if last == end {
			return soFar
		}
		for _, neighbor := range g[last] {
			if visited[neighbor] || omit[edge{last, neighbor}] {
				continue
			}
			visited[neighbor] = true
			todo = append(todo, append(slices.Clip(soFar), neighbor))
		}
	}

	return nil
}

func pathsBetween(start, end string, g map[string][]string) int {
	omit := make(map[edge]bool)

	for i := 0; i < 4; i++ {
		path := pathBetween(start, end, g, omit)
		if path == nil {
			return i
		}
		if i == 3 {
			return 4
		}

		for i, from := range path[:len(path)-1] {
			to := path[i+1]
			omit[edge{from, to}] = true
		}
	}
	return 4
}

func cutResult(g map[string][]string, edges []edge) (int, bool) {
	omit := make(map[edge]bool)
	for _, e := range edges {
		e.addToMap(omit)
	}
	group1 := connected(edges[0][0], g, omit)
	group2 := connected(edges[0][1], g, omit)

	lg := len(g)
	l1 := len(group1)
	l2 := len(group2)

	return l1 * l2, l1+l2 == lg
}

func part1(inputs []string) (int, error) {
	nodes, g, err := parseGraph(inputs)
	if err != nil {
		return 0, err
	}

	seed := nodes[0]
	groupA := map[string]bool{
		seed: true,
	}

	for _, node := range nodes[1:] {
		if pathsBetween(seed, node, g) > 3 {
			groupA[node] = true
		}
	}

	return len(groupA) * (len(g) - len(groupA)), nil
}

func part2(inputs []string) (int, error) {
	return 42, nil
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
