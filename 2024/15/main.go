package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) (charmap.M, []geom.Vec2, error) {
	paras := util.LinesByParagraph(inputs)
	if len(paras) != 2 {
		return nil, nil, fmt.Errorf("weird input: want 2 paragraphs; got %d: \n%s", len(paras), strings.Join(inputs, "\n"))
	}

	m := charmap.Parse(paras[0])
	dirs := util.Map([]rune(strings.Join(paras[1], "")), func(r rune) geom.Vec2 {
		switch r {
		case '^':
			return geom.N
		case 'v':
			return geom.S
		case '>':
			return geom.E
		case '<':
			return geom.W
		}
		panic(fmt.Sprintf("Weird direction char: '%c'", r))
	})

	return m, dirs, nil
}

func walk(dir geom.Vec2, m charmap.M) {
	pos, found := m.Find('@')
	if !found {
		panic("robot not found!")
	}

	var p geom.Vec2
	for p = pos.Add(dir); m[p] == 'O'; p = p.Add(dir) {
	}
	for m[p] == '.' {
		next := p.Sub(dir)
		r := m[next]
		m[p] = r
		m[next] = '.'
		if r == '@' {
			break
		}
		p = next
	}
}

func score(m charmap.M) int {
	total := 0
	for _, pos := range m.FindAll('O') {
		total += 100*pos.Y + pos.X
	}
	for _, pos := range m.FindAll('[') {
		total += 100*pos.Y + pos.X
	}
	return total
}

func double(m charmap.M) charmap.M {
	res := make(charmap.M, len(m)*2)

	for pos, r := range m {
		newPos := pos.WithX(pos.X * 2)
		newPos2 := newPos.E()

		switch r {
		case '#', '.':
			res[newPos] = r
			res[newPos2] = r
		case 'O':
			res[newPos] = '['
			res[newPos2] = ']'
		case '@':
			res[newPos] = '@'
			res[newPos2] = '.'
		}
	}

	return res
}

func boxCanMove(pos geom.Vec2, dir geom.Vec2, m charmap.M) bool {
	if m[pos] != '[' {
		panic("boxCanMove needs the left hand side of a box")
	}

	switch dir {
	case geom.W:
		switch m[pos.W()] {
		case '.':
			return true
		case '#':
			return false
		case ']':
			return boxCanMove(pos.W().W(), dir, m)
		}
	case geom.E:
		switch m[pos.E().E()] {
		case '.':
			return true
		case '#':
			return false
		case '[':
			return boxCanMove(pos.E().E(), dir, m)
		}
	case geom.N, geom.S:
		p1 := pos.Add(dir)
		p2 := p1.E()
		r1 := m[p1]
		r2 := m[p2]
		s := string(r1) + string(r2)
		if r1 == '#' || r2 == '#' {
			return false
		}
		if s == "[]" {
			return boxCanMove(p1, dir, m)
		}
		res := true
		if r1 == ']' {
			res = res && boxCanMove(p1.W(), dir, m)
		}
		if r2 == '[' {
			res = res && boxCanMove(p2, dir, m)
		}
		return res
	}
	panic("weird dir")
}

func canMove(pos geom.Vec2, dir geom.Vec2, m charmap.M) bool {
	if m[pos] != '@' {
		panic("canMove only works for the player")
	}
	newPos := pos.Add(dir)
	r := m[newPos]
	switch r {
	case '.':
		return true
	case '#':
		return false
	case '[':
		return boxCanMove(newPos, dir, m)
	case ']':
		return boxCanMove(newPos.W(), dir, m)
	}
	panic(fmt.Sprintf(`weird char: "%c"`, r))
}

func moveBox(pos geom.Vec2, dir geom.Vec2, m charmap.M) {
	if m[pos] != '[' {
		panic("moveBox needs the left hand side of a box")
	}

	switch dir {
	case geom.W:
		if m[pos.W()] == ']' {
			moveBox(pos.W().W(), dir, m)
		}
		m[pos.W()] = '['
		m[pos] = ']'
		m[pos.E()] = '.'
	case geom.E:
		if m[pos.E().E()] == '[' {
			moveBox(pos.E().E(), dir, m)
		}
		m[pos] = '.'
		m[pos.E()] = '['
		m[pos.E().E()] = ']'
	case geom.N, geom.S:
		p1 := pos.Add(dir)
		p2 := p1.E()
		r1 := m[p1]
		r2 := m[p2]
		s := string(r1) + string(r2)
		if s == "[]" {
			moveBox(p1, dir, m)
		}
		if r1 == ']' {
			moveBox(p1.W(), dir, m)
		}
		if r2 == '[' {
			moveBox(p2, dir, m)
		}
		m[p1] = '['
		m[p2] = ']'
		m[pos] = '.'
		m[pos.E()] = '.'
	default:
		panic("weird dir")
	}
}

func move(pos geom.Vec2, dir geom.Vec2, m charmap.M) {
	if m[pos] != '@' {
		panic("canMove only works for the player")
	}
	newPos := pos.Add(dir)
	r := m[newPos]
	if r == '[' {
		moveBox(newPos, dir, m)
	} else if r == ']' {
		moveBox(newPos.W(), dir, m)
	}
	m[newPos] = '@'
	m[pos] = '.'
}

func maybeMove(dir geom.Vec2, m charmap.M) {
	pos, found := m.Find('@')
	if !found {
		panic("robot not found!")
	}

	if canMove(pos, dir, m) {
		move(pos, dir, m)
	}
}

func part1(inputs []string) (int, error) {
	m, dirs, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	for _, dir := range dirs {
		walk(dir, m)
	}

	fmt.Printf("%s\n\n", m.AsString('_'))

	return score(m), nil
}

func part2(inputs []string) (int, error) {
	m, dirs, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	m = double(m)
	for _, dir := range dirs {
		maybeMove(dir, m)
	}
	fmt.Printf("%s\n\n\n", m.AsString('_'))
	return score(m), nil
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
