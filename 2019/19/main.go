package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/2019/intcode"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

type vec2 = geom.Vec2

type state struct {
	program []int64
	y10k    int
}

func newState(filename string) (*state, error) {
	s := &state{}
	program, err := intcode.ReadProgram(filename)
	if err != nil {
		return s, err
	}
	s.program = program
	s.y10k, err = s.findy(10000)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *state) check(x, y int) (bool, error) {
	_, writes, err := intcode.RunProgram(s.program, []int64{int64(x), int64(y)}, false)
	if err != nil {
		return false, err
	}

	return writes[0] == 1, nil
}

func (s *state) findy(x int) (int, error) {
	start := x
	if s.y10k > 0 {
		start = (x + 1) * s.y10k / 9990
	}
	active, err := s.check(x, start)
	if err != nil {
		return 0, err
	}
	if active {
		return 0, fmt.Errorf("bad guess started true: x=%d, start=%d", x, start)
	}
	for y := start; y >= 0; y-- {
		active, err := s.check(x, y)
		if err != nil {
			return 0, err
		}
		if active {
			return y, nil
		}
	}
	return 0, fmt.Errorf("not found")
}

func run() error {
	s, err := newState("input")
	if err != nil {
		return err
	}

	count := 0

	var max vec2
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			active, err := s.check(x, y)
			if err != nil {
				return err
			}
			if active {
				fmt.Printf("#")
				count++
				max = vec2{x, y}
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println()
		if max.Y < y-2 {
			break
		}
	}
	fmt.Println(count)

	predicate := func(x int) (bool, error) {
		y, err := s.findy(x)
		if err != nil {
			return false, err
		}
		active, err := s.check(x+99, y-99)
		if err != nil {
			return false, err
		}
		return active, nil
	}

	x, err := util.LowestTrue(500, predicate)
	if err != nil {
		return err
	}

	// Search back a tiny bit in case there are gaps...
	for maybeX := x - 1; maybeX > x-10; maybeX-- {
		good, err := predicate(maybeX)
		if err != nil {
			continue
		}
		if good {
			x = maybeX
		}
	}

	y, err := s.findy(x)
	if err != nil {
		return err
	}

	fmt.Printf("%d\n", x*10000+y-99)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
