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
