package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

const (
	top = iota
	right
	bottom
	left
)

type piece struct {
	id    int
	sides [4]string
}

func (p piece) fliplr() piece {
	return piece{
		id:    p.id,
		sides: [4]string{util.Reverse(p.sides[top]), p.sides[left], util.Reverse(p.sides[bottom]), p.sides[right]},
	}
}

func (p piece) clockwise() piece {
	return piece{
		id:    p.id,
		sides: [4]string{util.Reverse(p.sides[left]), p.sides[top], util.Reverse(p.sides[right]), p.sides[bottom]},
	}
}

func (p piece) ways() []piece {
	result := make([]piece, 0, 8)
	seen := make(map[piece]bool)

	for i := 0; i < 4; i++ {
		if !seen[p] {
			result = append(result, p)
			seen[p] = true
		}
		p = p.clockwise()
	}

	p = p.fliplr()

	for i := 0; i < 4; i++ {
		if !seen[p] {
			result = append(result, p)
			seen[p] = true
		}
		p = p.clockwise()
	}

	return result
}

type board [][]piece

func (b board) N() int {
	return len(b)
}

func NewBoard(n int) board {
	b := make([][]piece, n)
	for i := range b {
		b[i] = make([]piece, n)
	}
	return b
}

func (b board) Copy() board {
	c := make([][]piece, b.N())

	for i, row := range b {
		for _, col := range row {
			c[i] = append(c[i], col)
		}
	}

	return c
}

func (b board) Next(pos geom.Vec2) geom.Vec2 {
	n := b.N()
	pos.X = (pos.X + 1) % n
	if pos.X == 0 {
		pos.Y = (pos.Y + 1) % n
	}
	return pos
}

func (b board) Fit(pos geom.Vec2, p piece) bool {
	want := b[pos.Y][pos.X]
	for i, w := range want.sides {
		if w != "" && w != p.sides[i] {
			return false
		}
	}
	return true
}

func (b board) Place(pos geom.Vec2, p piece) board {
	bb := b.Copy()
	n := b.N()
	x, y := pos.X, pos.Y
	bb[y][x] = p
	if pos.X > 0 {
		bb[y][x-1].sides[right] = p.sides[left]
	}
	if pos.X < n-1 {
		bb[y][x+1].sides[left] = p.sides[right]
	}
	if pos.Y > 0 {
		bb[y-1][x].sides[bottom] = p.sides[top]
	}
	if pos.Y < n-1 {
		bb[y+1][x].sides[top] = p.sides[bottom]
	}

	return bb
}

type bag map[int]piece

func (bg bag) bg(p piece) {
	bg[p.id] = p
}

// Take returns a copy of the bag without the given piece.
func (bg bag) Take(p piece) bag {
	if _, ok := bg[p.id]; !ok {
		panic(fmt.Sprintf("can't remove piece %d: not in bag: %v", p.id, bg))
	}
	bb := make(map[int]piece, len(bg)-1)
	for k, v := range bg {
		if k != p.id {
			bb[k] = v
		}
	}
	return bb
}

func pieceFromCharmap(id int, m charmap.M) piece {
	tl := geom.Vec2{X: 0, Y: 0}
	tr := geom.Vec2{X: 9, Y: 0}
	br := geom.Vec2{X: 9, Y: 9}
	bl := geom.Vec2{X: 0, Y: 9}

	return piece{
		id: id,
		sides: [4]string{
			m.SliceAsString(tl, tr, '.'),
			m.SliceAsString(tr, br, '.'),
			m.SliceAsString(bl, br, '.'),
			m.SliceAsString(tl, bl, '.'),
		},
	}
}

func parse(inputs []string) (bag, map[int]charmap.M, error) {
	paras := util.LinesByParagraph(inputs)
	bg := make(bag)
	charmaps := make(map[int]charmap.M)

	for _, para := range paras {
		first := para[0]
		if !strings.HasPrefix(first, "Tile ") || !strings.HasSuffix(first, ":") {
			return nil, nil, fmt.Errorf("Weird start of tile: %q", first)
		}

		id, err := strconv.Atoi(first[5 : len(first)-1])
		if err != nil {
			return nil, nil, fmt.Errorf("weird start of tile: %q: %w", first, err)
		}

		m := charmap.Parse(para[1:])
		p := pieceFromCharmap(id, m)

		charmaps[id] = m
		bg.bg(p)
	}

	return bg, charmaps, nil
}

func findOrientation(m charmap.M, p piece) charmap.M {
	for i := 0; i < 4; i++ {
		if pieceFromCharmap(p.id, m) == p {
			return m
		}
		flipped := m.FlipLR()
		if pieceFromCharmap(p.id, flipped) == p {
			return flipped
		}
		m = m.Clockwise()
	}
	return nil
}

func solve(bg bag, b board, pos geom.Vec2, prefix string) board {
	// fmt.Printf("%s%v\n", prefix, pos)
	if len(bg) == 0 {
		return b
	}
	pos2 := b.Next(pos)
	prefix2 := prefix + " "
	for _, p := range bg {
		bg2 := bg.Take(p)
		for _, pp := range p.ways() {
			if !b.Fit(pos, pp) {
				continue
			}
			b2 := solve(bg2, b.Place(pos, pp), pos2, prefix2)
			if b2 != nil {
				return b2
			}
		}
	}
	return nil
}

func parseAndSolve(inputs []string) (board, map[int]charmap.M, error) {
	bg, charmaps, err := parse(inputs)
	if err != nil {
		return nil, nil, err
	}
	var n int
	switch len(bg) {
	case 9:
		n = 3
	case 144:
		n = 12
	default:
		return nil, nil, fmt.Errorf("weird number of pieces: %d", len(bg))
	}

	b := solve(bg, NewBoard(n), geom.Vec2{}, "")
	if b == nil {
		return nil, nil, fmt.Errorf("no solution found")
	}

	return b, charmaps, nil
}

func part1(inputs []string) (int, error) {
	b, _, err := parseAndSolve(inputs)
	if err != nil {
		return 0, err
	}

	n := b.N()
	prod := 1
	prod *= b[0][0].id
	prod *= b[0][n-1].id
	prod *= b[n-1][n-1].id
	prod *= b[n-1][0].id

	return prod, nil
}

func part2(inputs []string) (int, error) {
	seamonster := charmap.Parse(util.TrimmedLines(`
	..................#.
	#....##....##....###
	.#..#..#..#..#..#...
	`)).Without('.')
	omonster := seamonster.Replacing('#', 'O')
	b, charmaps, err := parseAndSolve(inputs)
	if err != nil {
		return 0, err
	}

	n := b.N()
	image := charmap.New(n*8, n*8, 'X')

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			p := b[y][x]
			cm := findOrientation(charmaps[p.id], p)
			if cm == nil {
				return 0, fmt.Errorf("couldn't find charmap orientation at (%d,%d)", x, y)
			}
			cm = cm.Subset(geom.Vec2{X: 1, Y: 1}, geom.Vec2{X: 8, Y: 8})
			image.Paste(cm, geom.Vec2{X: x * 8, Y: y * 8})
		}
	}

	for i := 0; i < 8; i++ {
		poses := image.AllInstances(seamonster)
		if len(poses) > 0 {
			for _, pos := range poses {
				image.Paste(omonster, pos)
			}
			return image.Count('#'), nil
		}
		image = image.Clockwise()
		if i == 4 {
			image = image.FlipLR()
		}
	}
	return 0, nil
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
