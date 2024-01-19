package main

import (
	"fmt"
	"os"
	"strings"
)

// var printf = fmt.Printf
var printf = func(string, ...any) {}

type graph map[string][]string

func parse(inputs []string) graph {
	res := make(graph)
	for _, input := range inputs {
		a, b, _ := strings.Cut(input, "-")
		res[a] = append(res[a], b)
		res[b] = append(res[b], a)
	}

	return res
}

func paths(node string, seen string, g graph, path string) int {
	count := 0
	for _, neighbor := range g[node] {
		if neighbor == "start" {
			continue
		}
		if neighbor == "end" {
			printf("%s-end\n", path)
			count++
		} else if neighbor[0] >= 'A' && neighbor[0] <= 'Z' {
			count += paths(neighbor, seen, g, path+"-"+neighbor)
		} else {
			if !strings.Contains(seen, neighbor) {
				count += paths(neighbor, seen+neighbor, g, path+"-"+neighbor)
			}
		}
	}

	return count
}

func pathsWithDouble(node string, seen string, double string, doubleSeen bool, g graph, path string) int {
	count := 0
	for _, neighbor := range g[node] {
		if neighbor == "start" {
			continue
		}
		if neighbor == "end" {
			if strings.Contains(seen, double) {
				printf("%s-end\n", path)
				count++
			}
		} else if neighbor[0] >= 'A' && neighbor[0] <= 'Z' {
			count += pathsWithDouble(neighbor, seen, double, doubleSeen, g, path+"-"+neighbor)
		} else {
			if !strings.Contains(seen, neighbor) {
				if double == neighbor && !doubleSeen {
					count += pathsWithDouble(neighbor, seen, double, true, g, path+"-"+neighbor)
				} else {
					count += pathsWithDouble(neighbor, seen+neighbor, double, doubleSeen, g, path+"-"+neighbor)
				}
			}
		}
	}

	return count
}

func part1(inputs []string) (int, error) {
	g := parse(inputs)
	return paths("start", "", g, "start"), nil
}

func part2(inputs []string) (int, error) {
	g := parse(inputs)
	count := paths("start", "", g, "start")

	for node := range g {
		if node != "start" && node != "end" && node[0] >= 'a' && node[0] <= 'z' {
			count += pathsWithDouble("start", "", node, false, g, "start")
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
