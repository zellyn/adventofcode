package main

import (
	"fmt"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/graph"
	"github.com/zellyn/adventofcode/util"
)

type vec2 = geom.Vec2

const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	wall  = '#'
	space = '.'
	empty = ' '
)

type cell struct {
	pos              vec2
	exit             bool
	normalNeighbors  []*cell
	inwardNeighbors  []*cell
	outwardNeighbors []*cell
}

func (c *cell) neighborString() string {
	var normals []string
	for _, nn := range c.normalNeighbors {
		normals = append(normals, "("+nn.Key()+")")
	}
	var inwards []string
	for _, nn := range c.inwardNeighbors {
		inwards = append(inwards, "("+nn.Key()+")")
	}
	var outwards []string
	for _, nn := range c.outwardNeighbors {
		outwards = append(outwards, "("+nn.Key()+")")
	}
	return "[" + strings.Join(normals, ",") + " | " + strings.Join(inwards, ",") + " | " + strings.Join(inwards, ",") + "]"
}

func (c *cell) Key() string {
	return fmt.Sprintf("%d,%d", c.pos.X, c.pos.Y)
}

func (c *cell) End() bool {
	return c.exit
}

func (c *cell) Neighbors() []graph.CostedNode {
	var result []graph.CostedNode
	for _, nn := range c.normalNeighbors {
		result = append(result, graph.CostedNode{N: nn, Steps: 1})
	}
	for _, nn := range c.inwardNeighbors {
		result = append(result, graph.CostedNode{N: nn, Steps: 1})
	}
	for _, nn := range c.outwardNeighbors {
		result = append(result, graph.CostedNode{N: nn, Steps: 1})
	}
	return result
}

var _ = graph.Node(&cell{})

type leveledCell struct {
	level int
	cell  *cell
}

func (lc leveledCell) Key() string {
	return fmt.Sprintf("%d:%d,%d", lc.level, lc.cell.pos.X, lc.cell.pos.Y)
}

func (lc leveledCell) End() bool {
	return lc.cell.exit && lc.level == 0
}

func (lc leveledCell) Neighbors() []graph.CostedNode {
	var result []graph.CostedNode
	for _, nn := range lc.cell.normalNeighbors {
		newCell := leveledCell{
			level: lc.level,
			cell:  nn,
		}
		result = append(result, graph.CostedNode{N: newCell, Steps: 1})
	}
	for _, nn := range lc.cell.inwardNeighbors {
		newCell := leveledCell{
			level: lc.level + 1,
			cell:  nn,
		}
		result = append(result, graph.CostedNode{N: newCell, Steps: 1})
	}
	if lc.level > 0 {
		for _, nn := range lc.cell.outwardNeighbors {
			newCell := leveledCell{
				level: lc.level - 1,
				cell:  nn,
			}
			result = append(result, graph.CostedNode{N: newCell, Steps: 1})
		}
	}
	return result
}

var _ = graph.Node(leveledCell{})

type state struct {
	rm    map[vec2]rune
	m     map[vec2]*cell
	start vec2
	end   vec2
}

func newState(filename string) (*state, error) {
	lines, err := util.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	s := &state{
		rm: map[vec2]rune{},
		m:  map[vec2]*cell{},
	}

	// Read rune map.
	for y, line := range lines {
		for x, r := range line {
			pos := vec2{X: x, Y: y}
			s.rm[pos] = r
			if r == space {
				s.m[pos] = &cell{pos: pos}
			}
		}
	}

	// Compute initial cell info.
	for pos, r := range s.rm {
		if r != space {
			continue
		}
		c := s.m[pos]
		for _, nn := range geom.Neighbors4(pos) {
			if s.rm[nn] == space {
				c.normalNeighbors = append(c.normalNeighbors, s.m[nn])
			}
		}
		s.m[pos] = c
	}

	// Compute label info.
	firstLabels := map[string][]vec2{}
	for pos, r := range s.rm {
		if isUpper(r) {
			mapPos, inner, label := s.label(pos)
			if label == "" {
				panic(fmt.Sprintf("weird label at position %v", pos))
			}
			if label == "AA" {
				s.start = mapPos
				continue
			}
			if label == "ZZ" {
				s.end = mapPos
				s.m[mapPos].exit = true
				continue
			}

			if positions, ok := firstLabels[label]; ok {
				if len(positions) == 1 && positions[0] != mapPos {
					firstLabels[label] = append(firstLabels[label], mapPos)
					otherPos := positions[0]
					if inner {
						s.m[mapPos].inwardNeighbors = append(s.m[mapPos].inwardNeighbors, s.m[otherPos])
						s.m[otherPos].outwardNeighbors = append(s.m[otherPos].outwardNeighbors, s.m[mapPos])
					} else {
						s.m[mapPos].outwardNeighbors = append(s.m[mapPos].outwardNeighbors, s.m[otherPos])
						s.m[otherPos].inwardNeighbors = append(s.m[otherPos].inwardNeighbors, s.m[mapPos])
					}
				}
			} else {
				firstLabels[label] = append(firstLabels[label], mapPos)
			}
		}
	}

	if s.start == (vec2{}) {
		return nil, fmt.Errorf("Cannot find start 'AA'")
	}
	if s.end == (vec2{}) {
		return nil, fmt.Errorf("Cannot find end 'ZZ'")
	}
	return s, nil
}

func (s *state) draw() {
	charmap.Draw(s.rm, ' ')
}

func isUpper(r rune) bool {
	return strings.ContainsRune(uppercase, r)
}

func (s *state) label(pos vec2) (vec2, bool, string) {
	r0 := s.rm[pos]
	if !isUpper(r0) {
		return vec2{}, false, ""
	}

	var r1 rune
	var spacePos, otherPos vec2
	var inner bool
	for i, dir := range geom.Dirs4 {
		pos1 := pos.Add(dir)
		n1 := s.rm[pos1]
		if n1 == space {
			pos2 := pos.Sub(dir)
			n2 := s.rm[pos2]
			if isUpper(n2) {
				spacePos = pos1
				otherPos = pos2
				r1 = n2
				inner = s.rm[pos2.Sub(dir)] == empty
				break
			}
			return vec2{}, false, ""
		}
		if isUpper(n1) {
			pos2 := pos1.Add(dir)
			if s.rm[pos2] == space {
				spacePos = pos2
				otherPos = pos1
				r1 = n1
				inner = s.rm[pos.Sub(dir)] == empty
				break
			}
		}
		if i == 3 {
			fmt.Println("bad end")
			return vec2{}, false, ""
		}
	}

	if pos.X < otherPos.X || pos.Y < otherPos.Y {
		return spacePos, inner, string(r0) + string(r1)
	}
	return spacePos, inner, string(r1) + string(r0)
}

func (s *state) minSteps() (int, error) {
	c := s.m[s.start]
	steps, err := graph.Dijkstra(c)
	if err != nil {
		return 0, err
	}
	return steps, nil
}

func (s *state) minRecursiveSteps() (int, error) {
	lc := leveledCell{
		cell: s.m[s.start],
	}
	steps, err := graph.Dijkstra(lc)
	if err != nil {
		return 0, err
	}
	return steps, nil
}

func main() {
	fmt.Println("everything in the tests!")
}
