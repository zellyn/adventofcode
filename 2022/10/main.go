package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type machine struct {
	input        []string
	pc           int
	cycle        int
	x            int
	inc          *int
	lastStrength int
	lastX        int
}

func new(input []string) *machine {
	return &machine{
		input: input,
		x:     1,
	}
}

func (m *machine) step() bool {
	m.lastX = m.x
	if m.inc != nil {
		m.cycle++
		m.lastStrength = m.cycle * m.x
		// fmt.Printf("m.cycle=%d m.x=%d m.strength=%d\n", m.cycle, m.x, m.lastStrength)
		m.x += *m.inc
		m.inc = nil
		return true
	}

	if m.pc >= len(m.input) {
		return false
	}

	m.cycle++
	cmd := m.input[m.pc]
	m.pc++
	m.lastStrength = m.cycle * m.x
	// fmt.Printf("m.cycle=%d m.x=%d m.strength=%d\n", m.cycle, m.x, m.lastStrength)

	parts := strings.Split(cmd, " ")

	switch parts[0] {
	case "noop":
		return true
	case "addx":
		i, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Sprintf("weird command: %q: %v", cmd, err))
		}
		m.inc = &i
		return true

	default:
		panic(fmt.Sprintf("weird input: %q", cmd))
	}
}

func (m *machine) strength() int {
	return m.lastStrength
}

func (m *machine) pos() int {
	return m.lastX
}

func part1(inputs []string) (int, error) {
	m := new(inputs)
	sum := 0
	for i := 1; m.step(); i++ {
		if (i-20)%40 == 0 {
			// fmt.Printf("At cycle %d, strength=%d\n", i, m.strength())
			sum += m.strength()
		}
	}
	return sum, nil
}

func part2(inputs []string) ([]string, error) {
	var screen []string
	m := new(inputs)
	for i := 0; m.step(); i++ {
		x := i % 40
		if x == 0 {
			screen = append(screen, "")
		}
		diff := x - m.pos()
		if diff >= -1 && diff <= 1 {
			screen[len(screen)-1] += "#"
		} else {
			screen[len(screen)-1] += "."
		}
	}
	return screen, nil
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
