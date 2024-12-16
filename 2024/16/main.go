package main

import (
	"container/heap"
	"fmt"
	"maps"
	"math"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func cloneAndAdd(m map[geom.Vec2]bool, pos geom.Vec2) map[geom.Vec2]bool {
	res := make(map[geom.Vec2]bool, len(m)+1)
	maps.Copy(res, m)
	res[pos] = true
	return res
}

func (n nodeDetails) neighbors(m charmap.M) []nodeDetails {
	res := []nodeDetails{
		{
			id:    n.id.WithVel(n.id.Vel.Clockwise90()),
			steps: 1000 + n.steps,
			from:  n.id,
		},
		{
			id:    n.id.WithVel(n.id.Vel.CounterClockwise90()),
			steps: 1000 + n.steps,
			from:  n.id,
		},
	}

	if m[n.id.Step().Pos] != '#' {
		res = append(res, nodeDetails{
			id:    n.id.Step(),
			steps: 1 + n.steps,
			from:  n.id,
		})
	}

	return res
}

// A nodeHeap is a min-heap of nodeDetails
type nodeHeap []nodeDetails

func (h nodeHeap) Len() int { return len(h) }
func (h nodeHeap) Less(i, j int) bool {
	return h[i].steps < h[j].steps
}
func (h nodeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(nodeDetails))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *nodeHeap) pop() nodeDetails {
	return heap.Pop(h).(nodeDetails)
}

func (h *nodeHeap) push(n nodeDetails) {
	heap.Push(h, n)
}

type nodeDetails struct {
	id    geom.PosVel2
	steps int
	from  geom.PosVel2
}

func solve(m charmap.M) (int, int) {
	startPos, ok := m.Find('S')
	if !ok {
		panic("no start")
	}
	endPos, ok := m.Find('E')
	if !ok {
		panic("no end")
	}

	start := nodeDetails{
		id:    geom.PosVel2{Pos: startPos, Vel: geom.E},
		steps: 0,
		from:  geom.PosVel2{Pos: startPos, Vel: geom.E},
	}

	froms := make(map[geom.PosVel2][]geom.PosVel2)
	minDists := make(map[geom.PosVel2]int)
	maxSteps := math.MaxInt

	todo := nodeHeap([]nodeDetails{start})

	for todo.Len() > 0 {

		next := todo.pop()
		if next.steps > maxSteps {
			break
		}
		if next.id.Pos == endPos && next.steps < maxSteps {
			maxSteps = next.steps
		}

		known, ok := minDists[next.id]
		if ok {
			if known < next.steps {
				continue
			}
			// must have equal cost: merge froms
			mustAdd := true
			for _, from := range froms[next.id] {
				if from == next.from {
					mustAdd = false
					break
				}
			}
			if mustAdd {
				froms[next.id] = append(froms[next.id], next.from)
			}

			// neighbors have already been added
			continue
		} else {
			minDists[next.id] = next.steps
			froms[next.id] = []geom.PosVel2{next.from}

			for _, nn := range next.neighbors(m) {
				todo.push(nn)
			}
		}
	}

	processed := make(map[geom.PosVel2]bool)
	seenPositions := make(map[geom.Vec2]bool)
	seenPositions[endPos] = true
	fromTodo := []geom.PosVel2{
		endPos.WithVel(geom.N),
		endPos.WithVel(geom.S),
		endPos.WithVel(geom.E),
		endPos.WithVel(geom.W),
	}

	for len(fromTodo) > 0 {
		next := fromTodo[0]
		fromTodo = fromTodo[1:]
		if processed[next] {
			continue
		}
		processed[next] = true
		seenPositions[next.Pos] = true
		for _, from := range froms[next] {
			if !processed[from] {
				fromTodo = append(fromTodo, from)
			}
		}
	}

	return maxSteps, len(seenPositions)
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	minSteps, _ := solve(m)
	return minSteps, nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	_, seenCount := solve(m)
	return seenCount, nil
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
