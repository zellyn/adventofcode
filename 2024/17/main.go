package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type state struct {
	regs    [3]int
	program []int
	ip      int
	output  []int
}

func (s state) clone() state {
	return state{
		regs:    s.regs,
		program: slices.Clone(s.program),
		ip:      s.ip,
		output:  slices.Clone(s.output),
	}
}

func parse(inputs []string) (state, error) {
	paras := util.LinesByParagraph(inputs)
	if len(paras) != 2 {
		return state{}, fmt.Errorf("want 2 paragraphs; got %d", len(paras))
	}
	regs, err := util.MapE(paras[0], func(s string) (int, error) {
		_, rest, ok := strings.Cut(s, ": ")
		if !ok {
			return 0, fmt.Errorf("weird input: %q", s)
		}
		return strconv.Atoi(rest)
	})
	if err != nil {
		return state{}, nil
	}

	if len(regs) != 3 {
		return state{}, fmt.Errorf("want 3 register values; got %d", len(regs))
	}

	if len(paras[1]) != 1 {
		return state{}, fmt.Errorf("weird second para: want length==1; got %d", len(paras[1]))
	}

	_, rest, ok := strings.Cut(paras[1][0], "Program: ")
	if !ok {
		return state{}, fmt.Errorf("weird program input: %q", paras[1][0])
	}
	ints, err := util.ParseInts(rest, ",")
	if err != nil {
		return state{}, err
	}

	s := state{
		program: ints,
	}
	copy(s.regs[:], regs)
	return s, nil
}

func (s *state) combo(c int) int {
	if c < 0 || c > 6 {
		panic(fmt.Sprintf("weird combo arg: %d", c))
	}
	if c < 4 {
		return c
	}
	return s.regs[c-4]
}

func (s *state) step() bool {
	if s.ip >= len(s.program) {
		return true
	}

	inst, arg := s.program[s.ip], s.program[s.ip+1]
	s.ip += 2

	switch inst {
	case 0: // adv
		numerator := s.regs[0]
		denominator := 1 << s.combo(arg)
		s.regs[0] = numerator / denominator
	case 1: // bxl
		s.regs[1] = s.regs[1] ^ arg
	case 2: // bst
		s.regs[1] = s.combo(arg) % 8

	case 3: // jnz
		if s.regs[0] != 0 {
			s.ip = arg
		}
	case 4: // bxc
		s.regs[1] = s.regs[1] ^ s.regs[2]
	case 5: // out
		s.output = append(s.output, s.combo(arg)%8)
	case 6: // bdv
		numerator := s.regs[0]
		denominator := 1 << s.combo(arg)
		s.regs[1] = numerator / denominator
	case 7: // cdv
		numerator := s.regs[0]
		denominator := 1 << s.combo(arg)
		s.regs[2] = numerator / denominator
	default:
		panic(fmt.Sprintf("weird instruction %d at %d", inst, s.ip-2))
	}

	return false
}

func intsToStr(ints []int) string {
	return strings.Join(util.Map(ints, strconv.Itoa), ",")
}

func part1(inputs []string) ([]int, error) {
	s, err := parse(inputs)
	if err != nil {
		return nil, nil
	}

	for !s.step() {
	}

	return s.output, nil
}

func (s_orig state) yields(a int, want []int) bool {
	s := s_orig.clone()
	s.regs[0] = a

	done := false
	length := 0
	for !done {
		done = s.step()
		if len(s.output) != length {
			length = len(s.output)
			if length > len(want) {
				return false
			}
			if s.output[length-1] != want[length-1] {
				return false
			}
		}
	}
	return length == len(want)
}

func part2Easy(s state) int {
	for a := 0; ; a++ {
		if s.yields(a, s.program) {
			return a
		}
	}
}

func part2(inputs []string, bitsPerOutput int) (int, error) {
	s, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	if len(s.program) == 6 {
		return part2Easy(s), nil
	}

	aChoices := map[int]bool{0: true}

	for l := len(s.program) - 1; l >= 0; l-- {
		newChoices := make(map[int]bool)

		fmt.Println(aChoices)

		for a := range aChoices {
			a = a << bitsPerOutput
			limit := 1 << bitsPerOutput
			if a == 0 {
				limit = 8
			}

			for b := 0; b < limit; b++ {
				if s.yields(a+b, s.program[l:]) {
					newChoices[a+b] = true
				}
			}
		}

		aChoices = newChoices
	}

	minA := math.MaxInt
	for a := range aChoices {
		minA = min(a, minA)
	}

	return minA, nil
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
