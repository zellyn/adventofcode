package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/dgraph"
	"github.com/zellyn/adventofcode/geom"
)

// To make someone move up: 3  (<

var printf = func(string, ...any) (int, error) { return 0, nil }

// var printf = fmt.Printf

type arm struct {
	pos     geom.Vec2
	numeric bool
}

var overNumeric = map[geom.Vec2]rune{
	geom.V2(0, 0):   'A',
	geom.V2(-1, 0):  '0',
	geom.V2(0, -1):  '3',
	geom.V2(-1, -1): '2',
	geom.V2(-2, -1): '1',
	geom.V2(0, -2):  '6',
	geom.V2(-1, -2): '5',
	geom.V2(-2, -2): '4',
	geom.V2(0, -3):  '9',
	geom.V2(-1, -3): '8',
	geom.V2(-2, -3): '7',
}

var overDirectional = map[geom.Vec2]rune{
	geom.V2(0, 0):  'A',
	geom.V2(-1, 0): '^',
	geom.V2(-2, 1): '<',
	geom.V2(-1, 1): 'v',
	geom.V2(0, 1):  '>',
}

var posNumeric = map[rune]geom.Vec2{
	'A': geom.V2(0, 0),
	'0': geom.V2(-1, 0),
	'3': geom.V2(0, -1),
	'2': geom.V2(-1, -1),
	'1': geom.V2(-2, -1),
	'6': geom.V2(0, -2),
	'5': geom.V2(-1, -2),
	'4': geom.V2(-2, -2),
	'9': geom.V2(0, -3),
	'8': geom.V2(-1, -3),
	'7': geom.V2(-2, -3),
}

var posDirectional = map[rune]geom.Vec2{
	'A': geom.V2(0, 0),
	'^': geom.V2(-1, 0),
	'<': geom.V2(-2, 1),
	'v': geom.V2(-1, 1),
	'>': geom.V2(0, 1),
}

func (a *arm) over() rune {
	if a.numeric {
		return overNumeric[a.pos]
	} else {
		return overDirectional[a.pos]
	}
}

func numericMovesTo(target rune, pos geom.Vec2) (string, geom.Vec2) {
	if target < '0' || (target > '9' && target != 'A') {
		panic(fmt.Sprintf("weird numeric pad target: %c", target))
	}
	targetPos := posNumeric[target]

	var result string

	if pos.Y != 0 || targetPos.X > -2 {
		for targetPos.X < pos.X {
			result += "<"
			pos.X--
		}
	}

	if targetPos != (geom.Vec2{X: -2, Y: 0}) {
		for targetPos.Y > pos.Y {
			result += "v"
			pos.Y++
		}
	}

	for targetPos.X > pos.X {
		result += ">"
		pos.X++
	}
	for targetPos.Y < pos.Y {
		result += "^"
		pos.Y--
	}
	for targetPos.X < pos.X {
		result += "<"
		pos.X--
	}
	for targetPos.Y > pos.Y {
		result += "v"
		pos.Y++
	}
	return result, pos
}

const (
	downs  = "vvv"
	ups    = "^^^"
	lefts  = "<<<"
	rights = ">>>"
)

func numericMovesToMulti(target rune, pos geom.Vec2) ([]string, geom.Vec2) {
	if target < '0' || (target > '9' && target != 'A') {
		panic(fmt.Sprintf("weird numeric pad target: %c", target))
	}
	targetPos := posNumeric[target]

	diff := targetPos.Sub(pos)
	if diff == geom.Z2 {
		return []string{""}, pos
	}

	if diff.X == 0 {
		if diff.Y > 0 {
			return []string{downs[:diff.Y]}, targetPos
		}
		return []string{ups[:-diff.Y]}, targetPos
	}

	if diff.Y == 0 {
		if diff.X > 0 {
			return []string{rights[:diff.X]}, targetPos
		}
		return []string{lefts[:-diff.X]}, targetPos
	}

	res := make([]string, 0, 2)

	if diff.Y > 0 {
		if diff.X > 0 {
			res = append(res, rights[:diff.X]+downs[:diff.Y])
			if pos.X > -2 || targetPos.Y < 0 {
				res = append(res, downs[:diff.Y]+rights[:diff.X])
			}
		} else {
			res = append(res,
				downs[:diff.Y]+lefts[:-diff.X],
				lefts[:-diff.X]+downs[:diff.Y],
			)
		}
	} else {
		if diff.X > 0 {
			res = append(res,
				ups[:-diff.Y]+rights[:diff.X],
				rights[:diff.X]+ups[:-diff.Y],
			)
		} else {
			res = append(res, ups[:-diff.Y]+lefts[:-diff.X])
			if pos.Y < 0 || targetPos.X > -2 {
				res = append(res, lefts[:-diff.X]+ups[:-diff.Y])
			}
		}
	}

	return res, targetPos
}

func dirMovesToMulti(target rune, pos geom.Vec2) ([]string, geom.Vec2) {
	if !strings.ContainsRune("<>^vA", target) {
		panic(fmt.Sprintf("weird directional pad target: %c", target))
	}
	targetPos := posDirectional[target]

	diff := targetPos.Sub(pos)
	if diff == geom.Z2 {
		return []string{""}, pos
	}

	if diff.X == 0 {
		if diff.Y > 0 {
			return []string{downs[:diff.Y]}, targetPos
		}
		return []string{ups[:-diff.Y]}, targetPos
	}

	if diff.Y == 0 {
		if diff.X > 0 {
			return []string{rights[:diff.X]}, targetPos
		}
		return []string{lefts[:-diff.X]}, targetPos
	}

	res := make([]string, 0, 2)

	if diff.Y > 0 {
		if diff.X > 0 {
			res = append(res,
				rights[:diff.X]+downs[:diff.Y],
				downs[:diff.Y]+rights[:diff.X],
			)
		} else {
			res = append(res, downs[:diff.Y]+lefts[:-diff.X])
			if target != '<' {
				res = append(res, lefts[:-diff.X]+downs[:diff.Y])
			}
		}
	} else {
		if diff.X > 0 {
			res = append(res, rights[:diff.X]+ups[:-diff.Y])
			if pos.X > -2 {
				res = append(res, ups[:-diff.Y]+rights[:diff.X])
			}

		} else {
			res = append(res,
				ups[:-diff.Y]+lefts[:-diff.X],
				lefts[:-diff.X]+ups[:-diff.Y],
			)
		}
	}

	return res, targetPos
}

func numericMovesFor(s string) string {
	pos := geom.Z2
	var res string
	var moves string
	for _, r := range s {
		moves, pos = numericMovesTo(r, pos)
		res += moves + "A"
	}
	return res
}

func dirMovesTo(target rune, pos geom.Vec2) (string, geom.Vec2) {
	if !strings.ContainsRune("<>^vA", target) {
		panic(fmt.Sprintf("weird directional pad target: %c", target))
	}
	targetPos := posDirectional[target]

	var result string

	if pos.Y == 1 || targetPos.X > -2 {
		for targetPos.X < pos.X {
			result += "<"
			pos.X--
		}
	}

	for targetPos.Y > pos.Y {
		result += "v"
		pos.Y++
	}
	for targetPos.X < pos.X {
		result += "<"
		pos.X--
	}
	for targetPos.X > pos.X {
		result += ">"
		pos.X++
	}
	for targetPos.Y < pos.Y {
		result += "^"
		pos.Y--
	}
	return result, pos
}

func dirMovesFor(s string) string {
	pos := geom.Z2
	var res string
	var moves string
	for _, r := range s {
		moves, pos = dirMovesTo(r, pos)
		res += moves + "A"
	}
	return res
}

func describe(numeric bool) string {
	if numeric {
		return "numeric keypad"
	}
	return "non-numeric keypad"
}

type cacheKey struct {
	s     string
	depth int
}

func movesFor(s string, depth int, cache map[cacheKey]int) int {
	if depth == 0 {
		return len(s)
	}
	key := cacheKey{s: s, depth: depth}
	if cache[key] > 0 {
		return cache[key]
	}

	var count int
	pos, moves := geom.Z2, ""
	for _, r := range s {
		moves, pos = dirMovesTo(r, pos)
		count += movesFor(moves+"A", depth-1, cache)
	}
	cache[key] = count
	return count
}

func movesForCode(code string, depth int, cache map[cacheKey]int) int {
	if depth == 0 {
		return len(code)
	}
	count := 0
	pos, moves := geom.Z2, ""
	for _, r := range code {
		moves, pos = numericMovesTo(r, pos)
		count += movesFor(moves+"A", depth-1, cache)
	}
	return count
}

func movesForMulti(s string, depth int, cache map[cacheKey]int) int {
	if depth == 0 {
		return len(s)
	}
	key := cacheKey{s: s, depth: depth}
	if cache[key] > 0 {
		return cache[key]
	}

	var count int
	var moveses []string
	pos := geom.Z2
	for _, r := range s {
		moveses, pos = dirMovesToMulti(r, pos)
		minCount := math.MaxInt - 1
		for _, moves := range moveses {
			minCount = min(minCount, movesForMulti(moves+"A", depth-1, cache))
		}
		count += minCount
	}
	cache[key] = count
	return count
}

func movesForCodeMulti(code string, depth int, cache map[cacheKey]int) int {
	if depth == 0 {
		return len(code)
	}
	count := 0
	pos := geom.Z2
	var moveses []string
	for _, r := range code {
		moveses, pos = numericMovesToMulti(r, pos)
		minCount := math.MaxInt - 1
		for _, moves := range moveses {
			minCount = min(minCount, movesForMulti(moves+"A", depth-1, cache))
		}
		count += minCount
	}
	return count
}

func (a arm) simulate(press rune) (arm, rune) {
	res := a
	switch press {
	case '^':
		res.pos = res.pos.N()
	case 'v':
		res.pos = res.pos.S()
	case '<':
		res.pos = res.pos.W()
	case '>':
		res.pos = res.pos.E()
	}
	o := res.over()
	if o == 0 {
		return res, -1
	}

	if press == 'A' {
		return res, o
	}
	return res, 0
}

type stack []arm

func (s stack) simulate(press rune) (stack, rune) {
	res := slices.Clone(s)
	for i, a := range res {
		a, press = a.simulate(press)
		if press == -1 {
			return nil, -1
		}
		res[i] = a
		if press == 0 {
			return res, 0
		}
	}
	return res, press
}

type state struct {
	want  []rune
	stack stack
}

func (s state) End() bool {
	return len(s.want) == 0
}

func (s state) Key() string {
	var rs []rune
	for _, a := range s.stack {
		rs = append(rs, a.over())
	}
	return string(s.want) + ":" + string(rs)
}

var _ dgraph.Node = state{}

func (s state) Neighbors() []dgraph.CostedNode {
	if len(s.want) == 0 {
		return nil
	}
	res := make([]dgraph.CostedNode, 0, 5)
	for _, r := range "^v<>A" {
		next, output := s.stack.simulate(r)
		if next == nil {
			continue
		}

		if output != 0 {
			if output != s.want[0] {
				continue
			}
			res = append(res, dgraph.CostedNode{
				N: state{
					stack: next,
					want:  s.want[1:],
				},
				Steps: 1,
			})
		} else {
			res = append(res, dgraph.CostedNode{
				N: state{
					stack: next,
					want:  s.want,
				},
				Steps: 1,
			})
		}
	}
	return res
}

func newStack(size int) stack {
	res := make(stack, size)
	res[size-1].numeric = true
	return res
}

func part1(inputs []string) (int, error) {
	cache := make(map[cacheKey]int)
	sum := 0
	for _, input := range inputs {
		count := movesForCodeMulti(input, 3, cache)
		i, err := strconv.Atoi(input[:len(input)-1])
		if err != nil {
			return 0, err
		}
		sum += i * count
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	cache := make(map[cacheKey]int)
	sum := 0
	for _, input := range inputs {
		count := movesForCodeMulti(input, 26, cache)
		i, err := strconv.Atoi(input[:len(input)-1])
		if err != nil {
			return 0, err
		}
		sum += i * count
	}
	return sum, nil
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
