package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/dgraph"
	"github.com/zellyn/adventofcode/util"
	"golang.org/x/exp/maps"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func newPart(name, part string) string {
	return strings.ToUpper(name[:1]) + strings.ToLower(name[1:2]) + strings.ToUpper(part[:1])
}

type state struct {
	elevator int
	floors   [][]string
}

const (
	sevenbits    = 0b111_1111
	fourteenbits = (sevenbits << 7) | sevenbits
	chipMask     = (sevenbits << 2) | (sevenbits << 16) | (sevenbits << 30) | (sevenbits << 44)
	genMask      = chipMask << 7
)

type state2 int64

func (s state) String() string {
	return fmt.Sprintf("%d:%v", s.elevator, s.floors)
}

func (s state) state2() state2 {
	matMap := make(map[string]bool)
	for _, parts := range s.floors {
		for _, part := range parts {
			matMap[part[:2]] = true
		}
	}
	keys := maps.Keys(matMap)
	sort.Strings(keys)

	bits := make(map[string]int)
	for i, key := range keys {
		bits[key] = i
	}

	res := s.elevator

	for floor, parts := range s.floors {
		baseOffset := floor*14 + 2
		for _, part := range parts {
			offset := baseOffset
			if part[2] == 'G' {
				offset += 7
			}
			offset += bits[part[:2]]
			res = res | (1 << offset)
		}
	}

	return state2(res)
}

func (s state) End() bool {
	if s.elevator != 3 {
		return false
	}
	for _, parts := range s.floors[:3] {
		if len(parts) > 0 {
			return false
		}
	}
	return true
}

func (s state2) End() bool {
	i := int(s)
	if i&3 != 3 {
		return false
	}

	mask := (1 << (2 + 3*14)) - 4
	return i&mask == 0
}

func (s state) sort() {
	for i := range s.floors {
		sort.Strings(s.floors[i])
	}
}

var tempGenerators = make(map[string]bool, 10)

func valid(parts []string) bool {
	if len(parts) == 0 {
		return true
	}
	maps.Clear(tempGenerators)
	for _, part := range parts {
		if part[2] == 'G' {
			tempGenerators[part[:2]] = true
		}
	}
	for _, part := range parts {
		if part[2] == 'G' {
			continue
		}
		name := part[:2]
		if len(tempGenerators) > 0 && !tempGenerators[name] {
			return false
		}
	}

	return true
}

func (s state) valid() bool {
	s2valid := s.state2().valid()
	if len(s.floors[s.elevator]) == 0 {
		if s2valid {
			panic(fmt.Sprintf("%s is invalid due to elevator on empty floor, but %d is not", s, s.state2()))
		}
		return false
	}
	for i, parts := range s.floors {
		if !valid(parts) {
			if s2valid {
				panic(fmt.Sprintf("%s is invalid due to invalid floor %d, but %d is not", s, i, s.state2()))
			}
			return false
		}
	}
	if !s2valid {
		panic(fmt.Sprintf("%s is valid but %d (%b) is not", s, s.state2(), s.state2()))
	}
	return true
}

func (s state2) valid() bool {
	i := int(s)
	ele := i & 3
	if i&(fourteenbits<<(2+ele*14)) == 0 {
		return false
	}

	chips := i & chipMask
	gens := i & genMask

	chips = chips &^ (gens >> 7)

	for i := 0; i < 4; i++ {
		mask := fourteenbits << (2 + i*14)
		if gens&mask > 0 && chips&mask > 0 {
			return false
		}
	}

	return true
}

func (s state) numParts() int {
	count := 0
	for _, parts := range s.floors {
		count += len(parts)
	}
	return count
}

func (s state) move(floor int, parts ...string) state {
	// printf(" move(%d→%d, %v)\n", s.elevator, floor, parts)
	res := state{
		elevator: floor,
		floors:   make([][]string, 4),
	}

	for i := range s.floors {
		switch {
		case i == s.elevator:
			res.floors[i] = util.Filter(s.floors[i], func(part string) bool {
				return !slices.Contains(parts, part)
			})

		case i == floor:
			ary := append(slices.Clip(s.floors[i]), parts...)
			sort.Strings(ary)
			res.floors[i] = ary

		default:
			res.floors[i] = s.floors[i]
		}
	}

	// printf(" → %s\n", res)
	return res
}

func (s state) Key() string {
	return s.String()
}

func (s state2) Key() string {
	return strconv.Itoa(int(s))
}

func (s state2) String() string {
	i := int(s)
	var res string
	for floor := 3; floor >= 0; floor-- {
		bits := (i >> (2 + floor*14) & fourteenbits)
		res += fmt.Sprintf("%07b %07b  ", bits>>7, bits&sevenbits)
	}
	res += fmt.Sprintf("%02b", i&3)
	return res
}

func setBits(i, min, max int) []int {
	var res []int
	for shift := min; shift <= max; shift++ {
		if i&(1<<shift) > 0 {
			res = append(res, shift)
		}
	}

	return res
}

func (s state) Neighbors() []dgraph.CostedNode {
	var res []dgraph.CostedNode

	for dest := s.elevator - 1; dest <= s.elevator+1; dest += 2 {
		if dest < 0 || dest > 3 {
			continue
		}

		for i, part1 := range s.floors[s.elevator] {
			moved := s.move(dest, part1)
			if moved.valid() {
				res = append(res, dgraph.CostedNode{N: moved, Steps: 1})
			}

			for _, part2 := range s.floors[s.elevator][i+1:] {
				moved := s.move(dest, part1, part2)
				if moved.valid() {
					res = append(res, dgraph.CostedNode{N: moved, Steps: 1})
				}
			}
		}
	}

	return res
}

func (s state2) Neighbors() []dgraph.CostedNode {
	var res []dgraph.CostedNode

	elevator := int(s) & 3
	fromShift := 2 + elevator*14
	ones := setBits(int(s), fromShift, fromShift+13)
	// printf(" elevator:%d fromShift:%d ones:%v\n", elevator, fromShift, ones)

	for dest := elevator - 1; dest <= elevator+1; dest += 2 {
		if dest < 0 || dest > 3 {
			continue
		}

		diff := 14
		if dest < elevator {
			diff = -14
		}

		toI := (int(s) &^ 3) | dest

		for i, bit1 := range ones {
			moved1 := state2(toI&^(1<<bit1) | (1 << (bit1 + diff)))
			if moved1.valid() {
				res = append(res, dgraph.CostedNode{N: moved1, Steps: 1})
			}

			for _, bit2 := range ones[i+1:] {
				moved2 := state2(int(moved1)&^(1<<bit2) | (1 << (bit2 + diff)))
				if moved2.valid() {
					res = append(res, dgraph.CostedNode{N: moved2, Steps: 1})
				}
			}
		}

	}

	return res
}

func parseInput(inputs []string) (state, error) {
	res := state{
		floors: make([][]string, 4),
	}

	for _, input := range inputs {
		floorName := strings.Split(input, " ")[1]
		floor := 0
		switch floorName {
		case "first":
			floor = 0
		case "second":
			floor = 1
		case "third":
			floor = 2
		case "fourth":
			floor = 3
		default:
			return state{}, fmt.Errorf("unknown floor: %q", floorName)
		}

		_, list, ok := strings.Cut(input, " contains ")
		if !ok {
			return state{}, fmt.Errorf("weird input line: %q", input)
		}

		cleanList := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(list, ", and ", " and "), " and ", ", "), ".", "")
		pieces := strings.Split(cleanList, ", ")

		if len(pieces) == 0 || len(pieces) == 1 && pieces[0] == "nothing relevant" {
			continue
		}

		var err error
		res.floors[floor], err = util.MapE(pieces, func(piece string) (string, error) {
			cleaned := strings.ReplaceAll(strings.ReplaceAll(piece, "-compatible", ""), "a ", "")
			name, part, ok := strings.Cut(cleaned, " ")
			if !ok {
				return "", fmt.Errorf("weird input: %q", input)
			}
			return newPart(name, part), nil
		})
		if err != nil {
			return state{}, err
		}
	}

	res.sort()
	return res, nil
}

func part1(inputs []string) (int, error) {
	s, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	return dgraph.Dijkstra(s)
}

func part2(inputs []string) (int, error) {
	s, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	s.floors[0] = append(s.floors[0], "ElG", "ElM", "DiG", "DiM")
	s.sort()

	s2 := s.state2() | 1
	printf("  %s\n", s2)
	for _, nn := range s2.Neighbors() {
		printf("→ %s\n", nn.N.(state2))
	}

	// return dgraph.Dijkstra(s)
	return dgraph.Dijkstra(s.state2())
	// return int(s.state2()), nil
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
