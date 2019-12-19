package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/2019/intcode"
	"github.com/zellyn/adventofcode/geom"
)

type vec2 = geom.Vec2

const (
	wall   = '#'
	space  = '.'
	oxygen = 'O'

	headingNorth = 1
	headingEast  = 2
	headingSouth = 3
	headingWest  = 4
)

var headingToVec = map[int]vec2{
	headingNorth: {0, -1},
	headingSouth: {0, 1},
	headingWest:  {-1, 0},
	headingEast:  {1, 0},
}

var vecToheading = map[vec2]int{
	{0, -1}: headingNorth,
	{0, 1}:  headingSouth,
	{-1, 0}: headingWest,
	{1, 0}:  headingEast,
}

var picToheading = map[rune]int{
	'^': headingNorth,
	'>': headingEast,
	'<': headingWest,
	'v': headingSouth,
}

var headingToPic = map[int]rune{
	headingNorth: '^',
	headingEast:  '>',
	headingWest:  '<',
	headingSouth: 'v',
}

type state struct {
	program    []int64
	readChan   chan int64
	writeChan  chan int64
	errChan    chan error
	m          map[vec2]rune
	pos        vec2
	heading    int
	progCancel context.CancelFunc
}

func (s *state) shutdown() {
	if s.progCancel != nil {
		s.progCancel()
	}
}

func newState(filename string) (*state, error) {
	program, err := intcode.ReadProgram("input")
	if err != nil {
		return &state{}, err
	}
	s := &state{
		program:   program,
		readChan:  make(chan int64),
		writeChan: make(chan int64),
		errChan:   make(chan error),
		m:         map[vec2]rune{},
		pos:       vec2{-1, -1},
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.progCancel = cancel
	go intcode.RunProgramChans(ctx, program, s.readChan, s.writeChan, s.errChan, nil, false, "")

	return s, nil
}

func (s *state) readMap() error {
	pos := vec2{0, 0}
	for {
		select {
		case i := <-s.writeChan:
			fmt.Printf("%c", rune(i))
			if i == 10 {
				pos.X = 0
				pos.Y++
			} else {
				s.m[pos] = rune(i)
				pos.X++
			}
		case err := <-s.errChan:
			return err
		}
	}
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
				c = headingToPic[s.heading]
			}
			if c == 0 {
				c = ' '
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func (s *state) findOneOf(chars string) (rune, vec2) {
	for pos, char := range s.m {
		if strings.ContainsRune(chars, char) {
			return char, pos
		}
	}
	return 0, vec2{-1, -1}
}

func (s *state) scoreIntersections() int {
	score := 0
OUTER:
	for pos, char := range s.m {
		if char != wall {
			continue
		}
		for heading := headingNorth; heading <= headingWest; heading++ {
			if s.m[pos.Add(headingToVec[heading])] != wall {
				continue OUTER
			}
		}
		score += pos.X * pos.Y
	}
	return score
}

func turn(currheading int, dir rune) int {
	d := 0
	if dir == 'L' {
		d = currheading - 1
	} else if dir == 'R' {
		d = currheading + 1
	}
	if d < headingNorth {
		d = headingWest
	} else if d > headingWest {
		d = headingNorth
	}
	return d
}

func (s *state) nextFrom(pos vec2, heading int) rune {
	for _, c := range "LR" {
		if s.m[pos.Add(headingToVec[turn(heading, c)])] == wall {
			return c
		}
	}
	return 0
}

func (s *state) maxSteps(pos vec2, heading int) (vec2, int) {
	vec := headingToVec[heading]
	steps := 0
	for {
		if s.m[pos.Add(vec)] != wall {
			return pos, steps
		}
		pos = pos.Add(vec)
		steps++
	}
}

type step struct {
	dir   rune
	steps int
}

func (s step) String() string {
	return fmt.Sprintf("%c,%d", s.dir, s.steps)
}

func (s *state) idealPath(pos vec2, heading int) []step {
	var result []step
	for {
		dir := s.nextFrom(pos, heading)
		if dir == 0 {
			return result
		}
		heading = turn(heading, dir)
		p, steps := s.maxSteps(pos, heading)
		pos = p
		result = append(result, step{dir: dir, steps: steps})
	}
}

func progToChars(prog []string) []int64 {
	s := strings.Join(prog, ",") + "\n"
	var result []int64
	for _, c := range s {
		result = append(result, int64(c))
	}
	return result
}

func (s *state) sendSoln(progs [][]string) (int, error) {
	program := intcode.Copy(s.program)
	program[0] = 2
	var reads []int64
	for _, prog := range progs {
		reads = append(reads, progToChars(prog)...)
	}
	reads = append(reads, progToChars([]string{"n"})...)
	_, writes, err := intcode.RunProgram(program, reads, false)
	if err != nil {
		return 0, err
	}
	for _, char := range writes {
		if char > 255 {
			return int(char), nil
		}
		fmt.Printf("%c", rune(char))
	}
	return 0, nil
}

func choices(steps []string, pos int) [][]string {
	var result [][]string
	for l := pos; l < len(steps); l++ {
		part := steps[pos : l+1]
		if len(strings.Join(part, ",")) > 20 {
			break
		}
		result = append(result, part)
	}
	return result
}

func allChoices(steps []string) [][]string {
	m := map[string]bool{}
	var result [][]string
	for i := 0; i < len(steps); i++ {
		for _, c := range choices(steps, i) {
			s := strings.Join(c, ",")
			if m[s] {
				continue
			}
			m[s] = true
			result = append(result, c)
		}
	}
	return result
}

func choiceCombos(steps []string) [][][]string {
	var result [][][]string
	ac := allChoices(steps)
	for _, a := range choices(steps, 0) {
		for _, b := range ac {
			for _, c := range ac {
				result = append(result, [][]string{a, b, c})
			}
		}
	}
	return result
}

func fillWith(steps []string, abc [][]string) [][]string {
	var result [][]string
	if len(steps) == 0 {
		return nil
	}
	for i, x := range abc {
		if !hasPrefix(steps, x) {
			continue
		}
		c := string('A' + rune(i))
		if len(x) == len(steps) {
			result = append(result, []string{c})
			continue
		}
		for _, way := range fillWith(steps[len(x):], abc) {
			result = append(result, append([]string{c}, way...))
		}
	}
	return result
}

func hasPrefix(strings, prefix []string) bool {
	if len(strings) < len(prefix) {
		return false
	}
	return eql(strings[:len(prefix)], prefix)
}

func eql(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, aa := range a {
		if aa != b[i] {
			return false
		}
	}
	return true
}

func findSolution(steps []string) [][]string {
	cc := choiceCombos(steps)
	for _, abc := range cc {
		ways := fillWith(steps, abc)
		for _, way := range ways {
			if len(way) < 11 {
				return [][]string{way, abc[0], abc[1], abc[2]}
			}
		}
	}
	return nil
}

func run() error {
	s, err := newState("input")
	if err != nil {
		return err
	}
	s.readMap()
	heading, pos := s.findOneOf("^><v")
	s.pos = pos
	s.heading = picToheading[heading]
	s.draw()
	fmt.Println(s.scoreIntersections())
	fmt.Printf("%c\n", s.nextFrom(s.pos, s.heading))
	path := s.idealPath(s.pos, s.heading)
	var steps []string
	for _, s := range path {
		steps = append(steps, string(s.dir))
		steps = append(steps, strconv.Itoa(s.steps))
	}
	fmt.Println(steps)
	soln := findSolution(steps)
	fmt.Println(soln)
	dust, err := s.sendSoln(soln)
	if err != nil {
		return err
	}
	fmt.Printf("dust: %d\n", dust)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
