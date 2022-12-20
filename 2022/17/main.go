package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/lists"
	"github.com/zellyn/adventofcode/util"
)

const BACKGROUND = '.'

const rockShapes = `
####

.#.
###
.#.

..#
..#
###

#
#
#
#

##
##
`

var LEFT = geom.Vec2{X: -1, Y: 0}
var RIGHT = geom.Vec2{X: 1, Y: 0}
var DOWN = geom.Vec2{X: 0, Y: -1}

type jets struct {
	s string
	i int
}

func (j *jets) next() geom.Vec2 {
	result := LEFT
	if j.s[j.i] == '>' {
		result = RIGHT
	}
	j.i = (j.i + 1) % len(j.s)
	return result
}

type rocks struct {
	maps []charmap.M
	i    int
}

func (r *rocks) next() charmap.M {
	result := r.maps[r.i].Copy()
	r.i = (r.i + 1) % len(r.maps)
	return result
}

func parseRocks() *rocks {
	var result []charmap.M
	lines := util.TrimmedLines(rockShapes)
	chunks := lists.Split(lines, func(s string) bool { return s == "" })
	for _, rock := range chunks {
		lists.Reverse(rock)
		m := charmap.ParseWithBackground(rock, '.')
		min, _ := m.MinMax()
		result = append(result, m.Translated(min.Neg()))
	}
	return &rocks{maps: result}
}

func drawWalls(m charmap.M, startY, endY int) int {
	m.DrawLine(geom.Vec2{X: -1, Y: startY}, geom.Vec2{X: -1, Y: endY}, '|')
	m.DrawLine(geom.Vec2{X: 7, Y: startY}, geom.Vec2{X: 7, Y: endY}, '|')
	return endY
}

func fall(m charmap.M, r *charmap.M) bool {
	r2 := r.Translated(DOWN)
	if r2.Overlaps(m) {
		return true
	}
	*r = r2
	return false
}

func blow(m charmap.M, r *charmap.M, dir geom.Vec2) {
	r2 := r.Translated(dir)
	if !r2.Overlaps(m) {
		*r = r2
	}
}

func printJoined(m, r charmap.M) {
	fmt.Println(m.Copy().Paste(r, geom.Vec2{}).AsStringFlipY(BACKGROUND))
}

type state struct {
	m       charmap.M
	rocks   *rocks
	jets    *jets
	r       charmap.M
	max     int
	wallMax int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (s *state) dropRock() {
	s.r = s.rocks.next()
	s.r = s.r.Translated(geom.Vec2{X: 2, Y: s.max + 4})
	s.wallMax = drawWalls(s.m, s.wallMax+1, s.r.MaxY())

	for {
		blow(s.m, &s.r, s.jets.next())
		if stuck := fall(s.m, &s.r); stuck {
			break
		}
	}
	s.m.Paste(s.r, geom.Vec2{})
	if rockMax := s.r.MaxY(); rockMax > s.max {
		s.max = rockMax
	}
}

func (s *state) clearBelow(minY int) {
	for k := range s.m {
		if k.Y < minY {
			delete(s.m, k)
		}
	}
}

func newState(input string) *state {
	rocks := parseRocks()
	jets := &jets{s: input}

	_, _ = rocks, jets

	m := charmap.Empty()
	m.DrawLine(geom.Vec2{X: 0, Y: 0}, geom.Vec2{X: 6, Y: 0}, '-')
	m.Put(-1, 0, '+')
	m.Put(7, 0, '+')

	return &state{
		m:     m,
		rocks: rocks,
		jets:  jets,
	}
}

func part1(input string) (int, error) {
	s := newState(input)
	for i := 0; i < 2022; i++ {
		s.dropRock()
	}

	return s.max, nil
}

var collapse_re = regexp.MustCompile(`R+|F+|L+`)

func (s *state) trace() (minY int, path string) {
	pos := geom.Vec2{X: 0, Y: s.max + 2}
	target := pos.WithX(6)
	// fmt.Printf("tracing: %s to %s\n", pos, target)
	dir := 0
	longPath := ""
	minY = pos.Y
	// pic = charmap.Empty()

	for pos != target {
		// fmt.Printf(" pos: %s\n", pos)
		// pic[pos] = '.'
		if pos.Y < minY {
			minY = pos.Y
		}
		// fmt.Printf(" dir=%d\n", dir)
		// fmt.Printf(" geom.Dirs4[dir]=%s\n", geom.Dirs4[dir])
		forward := pos.Add(geom.Dirs4[dir])
		right := pos.Add(geom.Dirs4[(dir+3)%4])
		// fmt.Printf(" forward=%s\n", forward)
		// fmt.Printf(" right=%s\n", right)
		// fmt.Printf(" right dir=%d\n", (dir+3)%4)
		// fmt.Printf(" geom.Dirs4[right dir]=%d\n", geom.Dirs4[(dir+3)%4])
		if !s.m.Has(right) {
			// fmt.Printf("  right to %s\n", right)
			// Go right if we can
			pos = right
			dir = (dir + 3) % 4
			longPath += "R"
		} else if !s.m.Has(forward) {
			// fmt.Printf("  forward to %s\n", forward)
			// Else go forward
			pos = forward
			longPath += "F"
		} else {
			// fmt.Printf("  turn left\n")
			// Reluctantly turn left
			dir = (dir + 1) % 4
			longPath += "L"
		}
		// pic[pos] = '.'
	}

	path = collapse_re.ReplaceAllStringFunc(longPath, func(s string) string { return fmt.Sprintf("%c%d", s[0], len(s)) })
	return minY, path //, pic
}

func part2(input string) (int, error) {
	trillion := 1000000000000
	s := newState(input)

	seen := make(map[string][2]int)

	var seenInfo [2]int
	var count int

	for j := 1; ; j++ {
		s.dropRock()
		var minY int
		minY, path := s.trace()
		if j%10000 == 0 {
			s.clearBelow(minY - 1)
		}

		key := fmt.Sprintf("%d-%d-%s", s.rocks.i, s.jets.i, path)
		var ok bool
		if seenInfo, ok = seen[key]; ok {
			fmt.Printf("path: %s\n", path)
			count = j
			break
		} else {
			seen[key] = [2]int{j, s.max}
		}
	}
	oldCount, oldMax := seenInfo[0], seenInfo[1]
	max := s.max
	maxGap := max - oldMax
	gap := count - oldCount
	piecesLeft := trillion - count
	piecesLeftToFinalMultiple := piecesLeft - piecesLeft%gap
	loopsLeft := piecesLeft - piecesLeftToFinalMultiple
	gapsLeftToFinalMultiple := piecesLeftToFinalMultiple / gap

	fmt.Printf("Dropping %d more pieces\n", loopsLeft)
	for i := 0; i < loopsLeft; i++ {
		s.dropRock()
	}
	extraHeight := s.max - max

	finalHeight := max + gapsLeftToFinalMultiple*maxGap + extraHeight

	fmt.Printf("After %d iterations, the max was %d\n", oldCount, oldMax)
	fmt.Printf("After %d iterations, the max was %d\n", count, max)
	fmt.Printf("That's %d height every %d iterations\n", maxGap, gap)
	fmt.Printf("%d pieces + %d × %d pieces + %d pieces = %d pieces\n", count, gapsLeftToFinalMultiple, gap, loopsLeft,
		count+gapsLeftToFinalMultiple*gap+loopsLeft)
	fmt.Printf("%d + %d × %d + %d = %d height\n", max, gapsLeftToFinalMultiple, maxGap, extraHeight,
		max+gapsLeftToFinalMultiple*maxGap+extraHeight)
	fmt.Printf("finalHeight=%d\n", finalHeight)
	return finalHeight, nil
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
