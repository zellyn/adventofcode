package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type state struct {
	pos geom.Vec2
	dir geom.Vec2
}

func walk(pos, dir geom.Vec2, m charmap.M) (map[geom.Vec2]bool, bool) {
	seen := map[geom.Vec2]bool{pos: true}
	seenState := make(map[state]bool)
	s := state{
		pos: pos,
		dir: dir,
	}

	for {
		if seenState[s] {
			return seen, true
		}
		seenState[s] = true
		switch m[s.pos.Add(s.dir)] {
		case '#':
			s.dir = s.dir.Clockwise90()
		case '.':
			s.pos = s.pos.Add(s.dir)
			seen[s.pos] = true
		case 0:
			return seen, false
		}
	}
}

func fastWalk(pos, dir, extra geom.Vec2, m charmap.M) bool {
	seenState := make(map[state]bool)
	s := state{
		pos: pos,
		dir: dir,
	}

	for {
		nextPos := s.pos.Add(s.dir)
		nextRune := m[nextPos]
		if nextPos == extra {
			nextRune = '#'
		}
		switch nextRune {
		case '#':
			s.dir = s.dir.Clockwise90()
			if seenState[s] {
				return true
			}
			seenState[s] = true
		case '.':
			s.pos = nextPos
		case 0:
			return false
		}
	}
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	startPos, found := m.Find('^')
	if !found {
		return 0, fmt.Errorf("Can't find '^' starting position")
	}
	m[startPos] = '.'

	seen, _ := walk(startPos, geom.N, m)
	return len(seen), nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	startPos, found := m.Find('^')
	if !found {
		return 0, fmt.Errorf("Can't find '^' starting position")
	}
	m[startPos] = '.'

	seen, _ := walk(startPos, geom.N, m)

	count := 0
	for pos := range seen {
		if pos == startPos {
			continue
		}
		loop := fastWalk(startPos, geom.N, pos, m)
		if loop {
			count++
		}
	}

	return count, nil
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
