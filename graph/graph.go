package graph

import (
	"bytes"
	"fmt"
)

type Node[T any] struct {
	Name  string
	Props T
	Start bool
	End   bool
}

func (n *Node[T]) String() string {
	var startString, endString string
	if n.Start {
		startString = " Start=true"
	}
	if n.End {
		startString = " End=true"
	}
	return fmt.Sprintf("Node:[%s%s%s]", n.Name, startString, endString)
}

type Edge[T any] struct {
	Cost int
	From *Node[T]
	To   *Node[T]
}

func (e *Edge[T]) String() string {
	return fmt.Sprintf("Edge:[%s -> %s (%d)]", e.From.Name, e.To.Name, e.Cost)
}

func (e *Edge[T]) ToEnd() bool {
	return e.To.End
}

type Graph[T any] struct {
	Nodes []*Node[T]
	Edges []*Edge[T]

	nameToNode    map[string]*Node[T]
	nameToEdges   map[string][]*Edge[T]
	nameToIsStart map[string]bool
	nameToIsEnd   map[string]bool
}

func (g *Graph[T]) String() string {
	b := new(bytes.Buffer)

	b.WriteString("Graph:[\n  Nodes:[\n")
	for _, node := range g.Nodes {
		b.WriteString("    " + node.String() + "\n")
	}
	b.WriteString("  ]\n  Edges:[\n")
	for _, edge := range g.Edges {
		b.WriteString("    " + edge.String() + "\n")
	}
	b.WriteString("  ]\n]")

	return b.String()
}

func NewGraph[T any](nodes []*Node[T], edges []*Edge[T]) *Graph[T] {
	g := &Graph[T]{
		Nodes:         nodes,
		Edges:         edges,
		nameToNode:    make(map[string]*Node[T], len(nodes)),
		nameToEdges:   make(map[string][]*Edge[T], len(nodes)),
		nameToIsStart: make(map[string]bool),
		nameToIsEnd:   make(map[string]bool),
	}

	for _, node := range nodes {
		g.nameToNode[node.Name] = node
		if node.Start {
			g.nameToIsStart[node.Name] = true
		}
		if node.End {
			g.nameToIsEnd[node.Name] = true
		}
	}

	for _, edge := range edges {
		g.nameToEdges[edge.From.Name] = append(g.nameToEdges[edge.From.Name], edge)
	}

	return g
}

func (g *Graph[T]) Starts() []*Node[T] {
	var res []*Node[T]
	for name := range g.nameToIsStart {
		res = append(res, g.nameToNode[name])
	}
	return res
}

func (g *Graph[T]) End(name string) bool {
	return g.nameToIsEnd[name]
}

func (g *Graph[T]) EdgesFrom(name string) []*Edge[T] {
	return g.nameToEdges[name]
}
