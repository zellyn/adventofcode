package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/zellyn/adventofcode/dgraph"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type node struct {
	pos     geom.Vec2
	size    int
	used    int
	index   int
	narrows bool
}

func (n node) avail() int {
	return n.size - n.used
}

var nodeRe = regexp.MustCompile(
	`^/dev/grid/node-x([0-9]+)-y([0-9]+)` +
		`\s+([0-9]+)T` +
		`\s+([0-9]+)T` +
		`\s+([0-9]+)T` +
		`\s+([0-9]+)%$`)

func parse(inputs []string) ([]node, error) {
	i := 0
	for !strings.HasPrefix(inputs[i], "/dev/") {
		i++
	}
	sais, err := util.ParseByRegexps(inputs[i:], []*regexp.Regexp{nodeRe})
	if err != nil {
		return nil, err
	}

	// Filesystem              Size  Used  Avail  Use%
	// /dev/grid/node-x0-y0     85T   67T    18T   78%

	return util.Map(sais, func(sai util.StringsAndInts) node {
		return node{
			pos:  geom.Vec2{X: sai.Ints[0], Y: sai.Ints[1]},
			size: sai.Ints[2],
			used: sai.Ints[3],
		}
	}), nil
}

func countConnectionsAndMarkUseful(nodes []node, useful map[geom.Vec2]bool) int {
	count := 0
	for i, nodeA := range nodes[:len(nodes)-1] {
		for _, nodeB := range nodes[i+1:] {
			if nodeA.used != 0 && nodeB.avail() >= nodeA.used {
				count++
				if useful != nil {
					useful[nodeA.pos] = true
					useful[nodeB.pos] = true
				}
			}
			if nodeB.used != 0 && nodeA.avail() >= nodeB.used {
				count++
				if useful != nil {
					useful[nodeA.pos] = true
					useful[nodeB.pos] = true
				}
			}
		}
	}
	return count
}

func part1(inputs []string) (int, error) {
	nodes, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	return countConnectionsAndMarkUseful(nodes, nil), nil
}

type state struct {
	neighbors [][]int
	data      int
	hole      int
	goal      int
}

func (s state) clone() state {
	s2 := s
	return s2
}

func (s state) Key() string {
	return fmt.Sprintf("%d-%d", s.data, s.hole)
}

func (s state) End() bool {
	if s.goal != 0 {
		return s.hole == s.goal
	}
	return s.data == 0
}

func (s state) Neighbors() []dgraph.CostedNode {
	res := make([]dgraph.CostedNode, 0, 4)

	holeIndex := s.hole

	for _, neighborIndex := range s.neighbors[holeIndex] {
		s2 := s
		s2.hole = neighborIndex
		if s.data == neighborIndex {
			s2.data = holeIndex
		}
		res = append(res, dgraph.CostedNode{N: s2, Steps: 1})
	}

	return res
}

func part2(inputs []string) (int, error) {
	nodes, err := parse(inputs)
	if err != nil {
		return 0, err
	}

	useful := make(map[geom.Vec2]bool)
	_ = countConnectionsAndMarkUseful(nodes, useful)

	tl, br := nodes[0].pos, nodes[0].pos

	nodeMap := make(map[geom.Vec2]node, len(nodes))
	// m := make(charmap.M)
	for _, n := range nodes {
		if useful[n.pos] {
			nodeMap[n.pos] = n
			// m[n.pos] = '#'
		}
		tl = geom.Min2(tl, n.pos)
		br = geom.Max2(br, n.pos)
	}
	// printf("%s\n", m.AsString('.'))
	tr := geom.Vec2{X: br.X, Y: tl.Y}
	_ = tr

	extent := geom.MakeRect(tl, br)

	s := state{
		neighbors: make([][]int, 0, len(nodes)),
	}

	index := 0
	for _, pos := range extent.Positions() {
		n, ok := nodeMap[pos]
		if !ok {
			continue
		}
		n.index = index
		nodeMap[pos] = n
		if pos == tr {
			s.data = index
		}
		if n.used == 0 && n.size > 0 {
			s.hole = index
		}

		index++
	}

	for _, pos := range extent.Positions() {
		_, ok := nodeMap[pos]
		if !ok {
			continue
		}
		neighbors := make([]int, 0, 4)
		up := false
		down := false
		for _, nPos := range pos.Neighbors4() {
			nn, ok := nodeMap[nPos]
			if !ok || nn.size == 0 {
				continue
			}
			if nPos == pos.N() {
				up = true
			} else if nPos == pos.S() {
				down = true
			}
			neighbors = append(neighbors, nn.index)
		}
		if len(neighbors) == 2 && up && down {
			printf("narrows at %s\n", pos)
			// cut off downward link from narrows
			s.neighbors = append(s.neighbors, []int{min(neighbors[0], neighbors[1])})
		} else {
			s.neighbors = append(s.neighbors, neighbors)
		}

	}

	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			n := nodeMap[geom.Vec2{X: x, Y: y}]
			printf("%d/%d ", n.used, n.size)
		}
		printf("\n")
	}

	return dgraph.Dijkstra(s)
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
