package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

type node struct {
	name   string
	depth  int
	parent string
}

func depth(name string, nodes map[string]node) int {
	v := nodes[name]
	if v.depth != -1 {
		return v.depth
	}
	if v.parent == "" || v.parent == "COM" {
		v.depth = 1
	} else {
		v.depth = depth(v.parent, nodes) + 1
	}
	nodes[name] = v
	return v.depth
}

func parents(name string, nodes map[string]node) []node {
	var result []node
	for name != "" {
		fmt.Println("name", name)
		n := nodes[name]
		result = append([]node{n}, result...)
		name = n.parent
	}
	return result
}

func run() error {
	nodes := map[string]node{
		"COM": node{name: "COM", depth: 0},
	}

	lines, err := util.ReadLines("input")
	if err != nil {
		return err
	}

	for _, line := range lines {
		parts := strings.Split(line, ")")
		if len(parts) != 2 {
			return fmt.Errorf("weird input line: %q", line)
		}
		nodes[parts[1]] = node{
			name:   parts[1],
			depth:  -1,
			parent: parts[0],
		}
	}

	sum := 0
	for k := range nodes {
		sum += depth(k, nodes)
	}

	fmt.Println(sum)

	youParents := parents("YOU", nodes)
	sanParents := parents("SAN", nodes)
	i := 0
	for ; youParents[i] == sanParents[i]; i++ {
	}
	fmt.Println(youParents)
	fmt.Println(sanParents)
	fmt.Println(len(youParents) + len(sanParents) - i - i - 2)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
