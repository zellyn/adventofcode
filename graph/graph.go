package graph

import (
	"container/heap"
	"fmt"
)

type Node interface {
	End() bool
	Neighbors() []CostedNode
	String() string
}

type CostedNode struct {
	N     Node
	Steps int
}

type costedNodeName struct {
	name  string
	steps int
	from  string
}

// A nodeHeap is a min-heap of costedNodeNames
type nodeHeap []costedNodeName

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].steps < h[j].steps }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(costedNodeName))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Dijkstra(start Node) (int, error) {
	// priors := map[string]string{}
	nameToNeighbor := map[string]Node{
		start.String(): start,
	}

	done := map[string]bool{}

	todo := nodeHeap([]costedNodeName{{name: start.String(), steps: 0}})

	for todo.Len() > 0 {
		costedName := heap.Pop(&todo).(costedNodeName)
		if done[costedName.name] {
			continue
		}
		done[costedName.name] = true
		node := nameToNeighbor[costedName.name]
		min := costedName.steps

		// fmt.Printf("Considering node %v: steps=%d\n", node, min)

		if node.End() {
			return min, nil
		}

		for _, info := range node.Neighbors() {
			neighbor := info.N
			neighborName := neighbor.String()
			steps := min + info.Steps

			nameToNeighbor[neighborName] = neighbor
			heap.Push(&todo, costedNodeName{name: neighborName, steps: steps})
		}
	}
	return 0, fmt.Errorf("no winning path found")
}

// PermutationsInt64 returns all permutations of a slice of int64s.
func PermutationsInt64(items []int64) [][]int64 {
	if len(items) <= 1 {
		return [][]int64{items}
	}

	var result [][]int64
	for i, item := range items {
		others := make([]int64, len(items)-1)
		copy(others, items[:i])
		copy(others[i:], items[i+1:])
		ps := PermutationsInt64(others)
		for _, p := range ps {
			val := append([]int64{item}, p...)
			result = append(result, val)
		}
	}
	return result
}

// PermutationsString returns all permutations of a slice of strings.
func PermutationsString(items []string) [][]string {
	if len(items) <= 1 {
		return [][]string{items}
	}

	var result [][]string
	for i, item := range items {
		others := make([]string, len(items)-1)
		copy(others, items[:i])
		copy(others[i:], items[i+1:])
		ps := PermutationsString(others)
		for _, p := range ps {
			val := append([]string{item}, p...)
			result = append(result, val)
		}
	}
	return result
}
