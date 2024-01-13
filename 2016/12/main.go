package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type op struct {
	name string
	x    string
	y    string
}

type state struct {
	ops  []op
	ip   int
	regs [4]int
	err  error
}

func (s *state) getValue(val string) (int, error) {
	if val == "a" || val == "b" || val == "c" || val == "d" {
		return s.regs[int(val[0]-'a')], nil
	}

	i, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (s *state) setValue(reg string, val int) error {
	switch reg {
	case "a", "b", "c", "d":
		s.regs[int(reg[0]-'a')] = val
		return nil
	}
	return fmt.Errorf("unknown register: %q", reg)
}

func (s *state) doSimple(o op, f func(i int) (string, int)) bool {
	val, err := s.getValue(o.x)
	if err != nil {
		s.err = err
		return true
	}
	target, newVal := f(val)
	if err := s.setValue(target, newVal); err != nil {
		s.err = err
		return true
	}
	s.ip++
	return s.ip >= len(s.ops)
}

func (s *state) step() bool {
	if s.err != nil || s.ip < 0 || s.ip >= len(s.ops) {
		return true
	}

	o := s.ops[s.ip]
	switch o.name {
	case "cpy":
		return s.doSimple(o, func(i int) (string, int) {
			return o.y, i
		})
	case "inc":
		return s.doSimple(o, func(i int) (string, int) {
			return o.x, i + 1
		})
	case "dec":
		return s.doSimple(o, func(i int) (string, int) {
			return o.x, i - 1
		})
	case "jnz":
		val1, err := s.getValue(o.x)
		if err != nil {
			s.err = nil
			return true
		}
		if val1 == 0 {
			s.ip++
		} else {
			val2, err := s.getValue(o.y)
			if err != nil {
				s.err = nil
				return true
			}
			s.ip += val2
		}
		return s.ip < 0 || s.ip >= len(s.ops)
	default:
		s.err = fmt.Errorf("weird op: %q", o.name)
		return true
	}

	return false
}

func parse(inputs []string) *state {
	res := &state{}

	for _, input := range inputs {
		parts := strings.Split(input, " ")
		if len(parts) == 2 {
			parts = append(parts, "")
		}
		res.ops = append(res.ops, op{name: parts[0], x: parts[1], y: parts[2]})
	}

	return res
}

func part1(inputs []string) (int, error) {
	s := parse(inputs)
	done := false
	for !done {
		done = s.step()
	}
	return s.getValue("a")
}

func part2(inputs []string) (int, error) {
	s := parse(inputs)
	s.setValue("c", 1)
	done := false
	for !done {
		done = s.step()
	}
	return s.getValue("a")
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
