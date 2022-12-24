package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/util"
)

type node struct {
	value int
	first bool
	prev  *node
	next  *node
}

func (n *node) walk(steps int) {
	// fmt.Printf(" walk(): node %d %d steps\n", n.value, steps)
	// fmt.Printf("  starting list: ")
	// n.printList()
	for i := 0; i < steps; i++ {
		// fmt.Printf("  step %d: \n", i+1)
		next := n.next
		// fmt.Printf("   next:%d\n", next.value)
		prev := n.prev
		// fmt.Printf("   prev:%d\n", prev.value)
		if n.first {
			n.first = false
			next.first = true
		}

		n.next = next.next
		next.next = n
		prev.next = next
		// fmt.Printf("   prev.next:%d, next.next:%d, n.next:%d\n", prev.next.value, next.next.value, n.next.value)

		next.prev = prev
		n.prev = next
		n.next.prev = n

		// fmt.Printf("   prev.prev:%d, next.prev:%d, n.prev:%d\n", prev.prev.value, next.prev.value, n.prev.value)

		// fmt.Printf("  result: ")
		// n.printList()
	}
}

func (n *node) printList() {
	for ; !n.first; n = n.next {
	}

	for {
		if !n.first {
			fmt.Printf(", ")
		}
		fmt.Print(n.value)
		n = n.next
		if n.first {
			break
		}
	}
	fmt.Println()
}

func setupList(inputs []string) []*node {
	ints := util.MustStringsToInts(inputs)
	mod := len(ints)

	nodes := make([]*node, len(ints))
	for i := range nodes {
		nodes[i] = &node{}
	}
	for i, value := range ints {
		nodes[i].value = value
		nodes[i].next = nodes[(i+1)%mod]
		nodes[i].prev = nodes[(i+mod-1)%mod]
		nodes[0].first = true
	}

	return nodes
}

func mix(nodes []*node) {
	moveMod := len(nodes) - 1
	for _, node := range nodes {
		steps := ((node.value % moveMod) + moveMod) % moveMod
		node.walk(steps)
	}

}

func findKey(nodes []*node) int {
	mod := len(nodes)

	n := nodes[0]
	for ; n.value != 0; n = n.next {
	}

	sum := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 1000%mod; j++ {
			n = n.next
		}
		sum += n.value
	}

	return sum
}

func part1(inputs []string) (int, error) {
	nodes := setupList(inputs)
	mix(nodes)
	return findKey(nodes), nil
}

func part2(inputs []string) (int, error) {
	key := 811589153
	nodes := setupList(inputs)
	for _, n := range nodes {
		n.value *= key
	}
	for i := 0; i < 10; i++ {
		mix(nodes)
	}
	return findKey(nodes), nil
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
