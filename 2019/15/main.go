package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/2019/intcode"
	"github.com/zellyn/adventofcode/geom"
)

type vec2 = geom.Vec2

func runSequence(program []int64, phases []int64, debug bool) (signal int64, err error) {
	signal = 0

	for _, phase := range phases {
		_, writes, err := intcode.RunProgram(program, []int64{phase, signal}, debug)
		if err != nil {
			return 0, err
		}
		if len(writes) != 1 {
			return 0, fmt.Errorf("want 1 write at phase %d; got %d (%v)", phase, len(writes), writes)
		}
		signal = writes[0]
	}
	return signal, nil
}

const (
	wall   = '#'
	space  = '.'
	oxygen = 'O'

	dirNorth = 1
	dirSouth = 2
	dirWest  = 3
	dirEast  = 4
)

var dirToVec = map[int]vec2{
	dirNorth: {0, -1},
	dirSouth: {0, 1},
	dirWest:  {-1, 0},
	dirEast:  {1, 0},
}

var vecToDir = map[vec2]int{
	{0, -1}: dirNorth,
	{0, 1}:  dirSouth,
	{-1, 0}: dirWest,
	{1, 0}:  dirEast,
}

type state struct {
	readChan   chan int64
	writeChan  chan int64
	errChan    chan error
	m          map[vec2]rune
	pos        vec2
	progCancel context.CancelFunc
}

func (s *state) shutdown() {
	if s.progCancel != nil {
		s.progCancel()
	}
}

func newState(filename string) (*state, error) {
	s := &state{
		readChan:  make(chan int64),
		writeChan: make(chan int64),
		errChan:   make(chan error),
		m: map[vec2]rune{
			vec2{0, 0}: space,
		},
	}
	program, err := intcode.ReadProgram("input")
	if err != nil {
		return s, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.progCancel = cancel
	go intcode.RunProgramChans(ctx, program, s.readChan, s.writeChan, s.errChan, false, "")

	return s, nil
}

func (s *state) move(to vec2) (bool, error) {
	diff := to.Sub(s.pos)
	dir := vecToDir[diff]
	if dir == 0 {
		return false, fmt.Errorf("Cannot move from %v to %v in one step", s.pos, to)
	}

	select {
	case s.readChan <- int64(dir):
		// cool
	case err := <-s.errChan:
		return false, err
	}

	status := 0

	select {
	case status64 := <-s.writeChan:
		status = int(status64)
	case err := <-s.errChan:
		return false, err
	}

	moved := true

	switch status {
	case 0:
		if s.m[to] != 0 && s.m[to] != wall {
			return false, fmt.Errorf("Got wall for pos %v; but was '%c' last time", to, s.m[to])
		}
		s.m[to] = wall
		moved = false
	case 1:
		if s.m[to] != 0 && s.m[to] != space {
			return false, fmt.Errorf("Got space for pos %v; but was '%c' last time", to, s.m[to])
		}
		s.m[to] = space
		s.pos = to
	case 2:
		if s.m[to] != 0 && s.m[to] != oxygen {
			return false, fmt.Errorf("Got oxygen for pos %v; but was '%c' last time", to, s.m[to])
		}
		s.m[to] = oxygen
		s.pos = to
	}

	return moved, nil
}

func (s *state) draw() {
	var min, max vec2
	for pos := range s.m {
		min = geom.Min2(min, pos)
		max = geom.Max2(max, pos)
	}
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			pos := vec2{x, y}
			c := s.m[pos]
			if s.pos == pos {
				c = 'R'
			}
			if c == 0 {
				c = ' '
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func (s *state) pathTo(from vec2, target rune) (bool, []vec2, error) {
	if s.m[from] == 0 {
		return false, nil, fmt.Errorf("Cannot pathTo from square %v of unknown state", from)
	}
	if s.m[from] == wall {
		return false, nil, fmt.Errorf("Cannot pathTo from wall square %v", from)
	}
	todo := []vec2{from}
	seen := map[vec2]vec2{from: from}
	foundPos := from
	found := false
	for len(todo) > 0 {
		pos := todo[0]
		todo = todo[1:]
		if s.m[pos] == target {
			foundPos = pos
			found = true
			break
		}

		for dir := dirNorth; dir <= dirEast; dir++ {
			newPos := pos.Add(dirToVec[dir])
			if _, ok := seen[newPos]; ok {
				continue
			}
			if s.m[newPos] == wall {
				continue
			}
			seen[newPos] = pos
			todo = append(todo, newPos)
		}
	}
	if !found {
		return false, nil, nil
	}
	var result []vec2
	for foundPos != from {
		result = append([]vec2{foundPos}, result...)
		foundPos = seen[foundPos]
	}
	return true, result, nil
}

func (s *state) diffuse() (int, error) {
	var start vec2
	found := false
	for k, v := range s.m {
		if v == oxygen {
			start = k
			found = true
		}
	}
	if !found {
		return 0, fmt.Errorf("Could not find oxygen")
	}
	todo := []vec2{start}

	minutes := -1
	for len(todo) > 0 {
		minutes++
		var nextTodo []vec2
		for _, pos := range todo {
			for dir := dirNorth; dir <= dirEast; dir++ {
				newPos := pos.Add(dirToVec[dir])
				if s.m[newPos] != space {
					continue
				}
				s.m[newPos] = oxygen
				nextTodo = append(nextTodo, newPos)
			}
		}

		todo = nextTodo
	}

	return minutes, nil
}

func run() error {

	s, err := newState("input")
	if err != nil {
		return err
	}
	s.move(vec2{-1, 0})
	s.move(vec2{0, 0})
	s.move(vec2{1, 0})
	s.move(vec2{0, 0})
	s.move(vec2{0, -1})
	s.move(vec2{0, 0})
	s.move(vec2{0, 1})
	s.move(vec2{0, 0})

	for {
		found, path, err := s.pathTo(s.pos, 0)
		if err != nil {
			return err
		}
		if !found {
			break
		}
		for _, newPos := range path {
			fmt.Println()
			_, err := s.move(newPos)
			if err != nil {
				return err
			}
		}
	}

	s.draw()
	found, path, err := s.pathTo(vec2{0, 0}, oxygen)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("Cannot find oxygen")
	}
	fmt.Println(len(path))

	minutes, err := s.diffuse()
	if err != nil {
		return err
	}
	fmt.Println(minutes)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
