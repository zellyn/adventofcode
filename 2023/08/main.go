package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// var printf = fmt.Printf
var printf = func(string, ...any) {}

type node struct {
	name  string
	left  string
	right string
}

var nodeRe = regexp.MustCompile(`[A-Z0-9]{3}`)

func parseInput(inputs []string) (string, map[string]node) {
	steps := inputs[0]

	nodeMap := make(map[string]node)
	for _, s := range inputs[2:] {
		found := nodeRe.FindAllString(s, -1)
		nodeMap[found[0]] = node{name: found[0], left: found[1], right: found[2]}
	}

	return steps, nodeMap
}

func part1(inputs []string) (int, error) {
	steps, nodes := parseInput(inputs)
	length := 0
	nodeName := "AAA"
	for i := 0; ; i++ {
		if nodeName == "ZZZ" {
			break
		}
		length++
		if steps[i%len(steps)] == 'L' {
			nodeName = nodes[nodeName].left
		} else {
			nodeName = nodes[nodeName].right
		}
	}
	return length, nil
}

func step(step byte, names []string, nodes map[string]node) {
	for i, name := range names {
		if step == 'L' {
			names[i] = nodes[name].left
		} else {
			names[i] = nodes[name].right
		}
	}
}

func allGood(names []string) bool {
	for _, name := range names {
		if !strings.HasSuffix(name, "Z") {
			return false
		}
	}
	return true
}

func part2(inputs []string) (int, error) {
	steps, nodes := parseInput(inputs)
	stepCount := len(steps)
	var names []string
	for name := range nodes {
		if strings.HasSuffix(name, "A") {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	i := 0
	for ; i < stepCount*100000; i++ {
		if allGood(names) {
			return i, nil
		}
		step(steps[i%stepCount], names, nodes)
	}
	checkpoint := slices.Clone(names)
	zSteps := make([]int, len(checkpoint))
	countCount := 0
	knownZs := make([]int, len(checkpoint))
	zLengths := make([]int, len(checkpoint))
	for countCount < len(checkpoint) {
		step(steps[i%stepCount], names, nodes)
		i++
		for j, name := range names {
			if name[2] == 'Z' {
				if zSteps[j] == 0 {
					zSteps[j] = i
					knownZs[j] = i
				} else {
					if zSteps[j] != -1 {
						count := i - zSteps[j]
						printf("%d looped in %d steps, which is %f step-loops\n", j, count, float32(count)/float32(stepCount))
						countCount++
						zLengths[j] = count
					}
					zSteps[j] = -1
				}
			}
		}
	}

	for j, name := range names {
		printf("%s hits zero at step %d and repeats in %d steps\n", name, knownZs[j], zLengths[j])
	}

	max := 0
	count := 0
	for done := false; !done; {
		if count%1000 == 0 {
			printf("%d: max=%d, knownZs=%v\n", count, max, knownZs)
		}
		count++
		done = true
		for j, knownZ := range knownZs {
			if knownZ > max {
				max = knownZ
			}
			for knownZ < max {
				done = false
				knownZ += zLengths[j]
				knownZs[j] = knownZ
				if knownZ > max {
					max = knownZ
				}
			}
		}
	}

	return knownZs[0], nil
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
