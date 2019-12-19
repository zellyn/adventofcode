package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/ioutil"
)

type vec2 = geom.Vec2

const (
	wall  = '#'
	space = '.'

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

var headingVectors = []vec2{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

type state struct {
	m         map[vec2]rune
	keys      string
	keyPos    map[rune]vec2
	startPos  vec2
	startPos4 [4]vec2
	infos     map[rune]map[vec2]info
}

func readMap(filename string) (m map[vec2]rune, startPos vec2, startPos4 [4]vec2, keyPos map[rune]vec2, err error) {
	m = map[vec2]rune{}
	keyPos = map[rune]vec2{}
	starts := 0
	lines, err := ioutil.ReadLines(filename)
	if err != nil {
		return
	}
	for y, line := range lines {
		for x, ch := range line {
			pos := vec2{x, y}
			m[pos] = ch
			if ch == '@' {
				if starts == 0 {
					startPos = pos
				} else {
					startPos4[0] = startPos
					startPos4[starts] = pos
				}
				starts++
			} else if isKey(ch) {
				keyPos[ch] = pos
			}
		}
	}
	return
}

func newState(filename string) (*state, error) {
	s := &state{
		m:     map[vec2]rune{},
		infos: map[rune]map[vec2]info{},
	}
	m, startPos, startPos4, keyPos, err := readMap(filename)
	if err != nil {
		return nil, err
	}
	s.m = m
	s.startPos = startPos
	s.startPos4 = startPos4
	s.keyPos = keyPos
	var keyStrings []string
	for key, pos := range s.keyPos {
		s.infos[key] = computeInfo(m, pos)
		keyStrings = append(keyStrings, string(key))
	}

	sort.Strings(keyStrings)
	s.keys = strings.Join(keyStrings, "")
	return s, nil
}

func drawMap(m map[vec2]rune) {
	var min, max vec2
	for pos := range m {
		min = geom.Min2(min, pos)
		max = geom.Max2(max, pos)
	}
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			pos := vec2{x, y}
			c := m[pos]
			if c == 0 {
				c = ' '
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}

}

func (s *state) draw() {
	drawMap(s.m)
}

func (s *state) findOneOf(chars string) (rune, vec2) {
	for pos, char := range s.m {
		if strings.ContainsRune(chars, char) {
			return char, pos
		}
	}
	return 0, vec2{-1, -1}
}

func isSimpleMap(m map[vec2]rune, start vec2) bool {
	froms := map[vec2]vec2{
		start: start,
	}
	path := []vec2{start}
OUTER:
	for len(path) > 0 {
		pos := path[len(path)-1]
		for _, vector := range headingVectors {
			newPos := pos.Add(vector)
			if m[newPos] == wall {
				continue
			}
			_, ok := froms[newPos]
			// fmt.Println(path, newPos)
			if ok {
				if froms[pos] == newPos {
					continue
				}
				if froms[newPos] == pos {
					continue
				}
				return false
			}
			froms[newPos] = pos
			path = append(path, newPos)
			continue OUTER
		}
		path = path[:len(path)-1]
	}
	return true
}

func isKey(ch rune) bool {
	return strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", ch)
}

func isDoor(ch rune) bool {
	return strings.ContainsRune("ABCDEFGHIJKLMNOPQRSTUVWXYZ", ch)
}

func keyToBit(ch rune) uint32 {
	return 1 << uint(ch-'a')
}

func doorToBit(ch rune) uint32 {
	return 1 << uint(ch-'A')
}

type info struct {
	seenDoors uint32
	seenKeys  uint32
	steps     int
}

func neighbors(pos vec2) []vec2 {
	return []vec2{
		pos.Add(vec2{0, -1}),
		pos.Add(vec2{1, 0}),
		pos.Add(vec2{0, 1}),
		pos.Add(vec2{-1, 0}),
	}
}

func computeInfo(m map[vec2]rune, startPos vec2) map[vec2]info {
	if !isKey(m[startPos]) {
		panic(fmt.Sprintf("computeInfo called with non-key start pos %v", startPos))
	}
	infos := map[vec2]info{
		startPos: info{
			seenKeys: keyToBit(m[startPos]),
		},
	}
	var todo []vec2
	for _, n := range neighbors(startPos) {
		if m[n] != wall {
			todo = append(todo, n)
		}
	}
	for len(todo) > 0 {
		pos := todo[0]
		todo = todo[1:]
		if _, ok := infos[pos]; ok {
			continue
		}
		ii := info{
			steps: -1,
		}
		for _, other := range neighbors(pos) {
			if m[other] == wall {
				continue
			}
			i, ok := infos[other]
			if !ok {
				todo = append(todo, other)
				continue
			}
			i.steps++
			ch := m[pos]
			if isKey(ch) {
				i.seenKeys |= keyToBit(ch)
			} else if isDoor(ch) {
				i.seenDoors |= doorToBit(ch)
			}
			if ii.steps > i.steps || ii.steps == -1 {
				ii = i
			}
		}
		if ii.steps == -1 {
			panic("BANG")
		}
		infos[pos] = ii
	}
	return infos
}

type args struct {
	pos  vec2
	seen uint32
}

type args4 struct {
	pos  [4]vec2
	seen uint32
}

func (s *state) best(debug bool) int {
	return s._best(s.startPos, s.keys, 0, "", debug, map[args]int{})
}

func (s *state) best4(debug bool) int {
	return s._best4(s.startPos4, s.keys, 0, "", debug, map[args4]int{})
}

func (s *state) _best(pos vec2, keys string, seen uint32, prefix string, debug bool, memo map[args]int) int {
	if keys == "" {
		return 0
	}
	if v := memo[args{pos, seen}]; v > 0 {
		return v
	}
	best := int(^uint(0) >> 1)
	if debug {
		fmt.Printf("%spos:%v\n", prefix, pos)
	}
	for i, key := range keys {
		if debug {
			fmt.Printf("%s  considering '%c'\n", prefix, key)
		}
		inf := s.infos[key][pos]
		if debug {
			fmt.Printf("%s    info: %v\n", prefix, inf)
		}
		bit := keyToBit(key)
		need := (inf.seenDoors | inf.seenKeys) &^ bit
		haveAll := seen&need == need
		if !haveAll {
			continue
		}
		if debug {
			fmt.Printf("%s  - '%c' is good\n", prefix, key)
		}
		steps := inf.steps
		score := s._best(s.keyPos[key], keys[:i]+keys[i+1:], seen|bit, prefix+"  ", debug, memo)
		if steps+score < best {
			best = steps + score
		}
	}
	memo[args{pos, seen}] = best
	return best
}

func (s *state) _best4(pos [4]vec2, keys string, seen uint32, prefix string, debug bool, memo map[args4]int) int {
	if keys == "" {
		return 0
	}
	if v := memo[args4{pos, seen}]; v > 0 {
		return v
	}
	best := int(^uint(0) >> 1)
	if debug {
		fmt.Printf("%spos:%v\n", prefix, pos)
	}
	for i, key := range keys {
		if debug {
			fmt.Printf("%s  considering '%c'\n", prefix, key)
		}
		bit := keyToBit(key)
		robot := 0
		var inf info
		for ; robot < 4; robot++ {
			inf = s.infos[key][pos[robot]]
			if inf.seenKeys&bit > 0 {
				break
			}
		}
		if robot > 3 {
			panic(fmt.Sprintf("Couldn't find robot for key: '%v'", key))
		}

		if debug {
			fmt.Printf("%s    info: %v\n", prefix, inf)
		}
		need := (inf.seenDoors | inf.seenKeys) &^ bit
		haveAll := seen&need == need
		if !haveAll {
			continue
		}
		if debug {
			fmt.Printf("%s  - '%c' is good\n", prefix, key)
		}
		steps := inf.steps
		posCopy := pos
		posCopy[robot] = s.keyPos[key]
		score := s._best4(posCopy, keys[:i]+keys[i+1:], seen|bit, prefix+"  ", debug, memo)
		if steps+score < best {
			best = steps + score
			if debug {
				fmt.Printf("%s  - robot %d moving to key %c is best\n", prefix, robot, key)
			}
		}
	}
	memo[args4{pos, seen}] = best
	return best
}

func run() error {
	s, err := newState("input")
	if err != nil {
		return err
	}

	fmt.Println(s.best(false))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
