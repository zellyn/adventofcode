package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/zellyn/adventofcode/stringset"
	"github.com/zellyn/adventofcode/util"
)

var printf = func(string, ...any) {}

// var printf = fmt.Printf

func parse(inputs []string) (map[string]stringset.S, error) {
	links := make(map[string]stringset.S)

	for _, input := range inputs {
		from, to, ok := strings.Cut(input, "-")
		if !ok {
			return nil, fmt.Errorf("weird input: %q", input)
		}

		if links[from] == nil {
			links[from] = make(stringset.S)
		}
		links[from][to] = true
		if links[to] == nil {
			links[to] = make(stringset.S)
		}
		links[to][from] = true
	}

	return links, nil
}

func part1(inputs []string) (int, error) {
	links, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	triples := make(stringset.S)

	for from, tos := range links {
		if !strings.HasPrefix(from, "t") {
			continue
		}

		for o1 := range tos {
			for o2 := range tos {
				if o1 == o2 || !links[o1][o2] {
					continue
				}

				names := []string{from, o1, o2}
				slices.Sort(names)
				key := strings.Join(names, "-")
				triples[key] = true
			}
		}
	}

	return len(triples), nil
}

func indent(depth int, format string, a ...any) {
	printf(fmt.Sprintf("%3d:", depth)+strings.Repeat(" ", depth)+format, a...)
}

func largestGroup(prefix string, d int, names []string, links map[string]stringset.S, in stringset.S, possible stringset.S) []string {
	// Base case
	if len(names) == 0 {
		return nil
	}

	this := names[0]
	indent(d, "%slargestGroup([%q,...], links, in=%s, out=%s\n", prefix, this, in)

	thisTo := links[this]

	canInclude := thisTo.ContainsAll(in)
	var res []string

	if canInclude {
		newNames := util.Filter(names[1:], func(name string) bool {
			return thisTo[name]
		})

		restWith := largestGroup("include: ", d+1, newNames, links, in.ClonePlus(this), possible)
		res = append(restWith, this)
	}

	restWithout := largestGroup("exclude: ", d+1, names[1:], links, in, possible)
	if len(restWithout) > len(res) {
		res = restWithout
	}

	return res
}

func part2(inputs []string) (string, error) {
	links, err := parse(inputs)
	if err != nil {
		return "", nil
	}

	var names []string
	for key := range links {
		names = append(names, key)
	}

	slices.Sort(names)
	names = slices.Compact(names)

	slices.SortFunc(names, func(a, b string) int {
		return cmp.Compare(len(links[a]), len(links[b]))
	})

	best := largestGroup("", 0, names, links, nil, stringset.New(names...))

	slices.Sort(best)

	return strings.Join(best, ","), nil
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
