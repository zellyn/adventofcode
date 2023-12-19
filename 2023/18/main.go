package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
	"golang.org/x/exp/maps"
)

type step struct {
	dir      geom.Vec2
	distance int
	color    string
}

func parse(inputs []string) ([]step, error) {
	sais, err := util.ParseStringsAndInts(inputs, 3, []int{0, 2}, []int{1})
	if err != nil {
		return nil, err
	}

	res := make([]step, 0, len(sais))

	for _, sai := range sais {
		color := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(sai.Strings[1], "(", ""), ")", ""), "#", "")
		res = append(res, step{
			dir:      geom.NameToDir[sai.Strings[0]],
			distance: sai.Ints[0],
			color:    color,
		})
	}

	return res, nil
}

func parse2(inputs []string) ([]step, error) {
	numToDir := map[byte]geom.Vec2{
		'0': geom.E,
		'1': geom.S,
		'2': geom.W,
		'3': geom.N,
	}

	res := make([]step, 0, len(inputs))

	for _, input := range inputs {
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			return nil, fmt.Errorf("Weird input line: %q", input)
		}
		color := parts[2]
		if len(color) != 9 {
			return nil, fmt.Errorf("Weird color field in input line: %q", input)
		}
		dir := numToDir[color[7]]
		steps, err := strconv.ParseInt(color[2:7], 16, 64)
		if err != nil {
			return nil, fmt.Errorf("Cannot hex-parse color in line %q: %w", input, err)
		}

		res = append(res, step{
			dir:      dir,
			distance: int(steps),
			color:    "#",
		})
	}

	return res, nil
}

func walk(steps []step) (map[geom.Vec2]string, geom.Vec2, geom.Vec2, geom.Vec2) {
	m := make(map[geom.Vec2]string)
	pos := geom.V2(0, 0)
	min := pos
	max := pos
	seed := pos.SE()

	for _, step := range steps {
		dir := step.dir
		for i := 0; i < step.distance; i++ {
			pos = pos.Add(dir)
			m[pos] = "#"
			pos = pos.Add(dir)
			m[pos] = "#"

			if pos.Y < min.Y || (pos.Y == min.Y && pos.X < min.X) {
				seed = pos.SE()
			}

			min = geom.Min2(min, pos)
			max = geom.Max2(max, pos)
		}
	}

	return m, min, max, seed
}

func breaksToPositions(breaks map[int]bool) map[int]int {
	ints := maps.Keys(breaks)
	sort.Ints(ints)
	res := make(map[int]int, len(ints))
	for index, i := range ints {
		res[i] = index * 2
	}
	return res
}

func flip(breaks map[int]int) map[int]int {
	res := make(map[int]int, len(breaks))

	for k, v := range breaks {
		res[v] = k
	}

	return res
}

type mapping struct {
	xPosToIndex map[int]int
	yPosToIndex map[int]int

	xIndexToPos map[int]int
	yIndexToPos map[int]int
}

func (m mapping) realToIndexed(real geom.Vec2) geom.Vec2 {
	xIndexed, ok := m.xPosToIndex[real.X]
	if !ok {
		panic(fmt.Sprintf("Cannot map real x=%d to indexed position", real.X))
	}
	yIndexed, ok := m.yPosToIndex[real.Y]
	if !ok {
		panic(fmt.Sprintf("Cannot map real y=%d to indexed position", real.Y))
	}
	return geom.Vec2{X: xIndexed, Y: yIndexed}
}

func (m mapping) indexedToReal(indexed geom.Vec2) geom.Vec2 {
	xReal, ok := m.xIndexToPos[indexed.X]
	if !ok {
		panic(fmt.Sprintf("Cannot map indexed x=%d to real position", indexed.X))
	}
	yReal, ok := m.yIndexToPos[indexed.Y]
	if !ok {
		panic(fmt.Sprintf("Cannot map indexed y=%d to real position", indexed.Y))
	}
	return geom.Vec2{X: xReal, Y: yReal}
}

func findBreaks(steps []step) mapping {
	xs := make(map[int]bool, len(steps))
	ys := make(map[int]bool, len(steps))
	pos := geom.V2(0, 0)

	for _, step := range steps {
		dir := step.dir
		pos = pos.Add(dir.Mul(step.distance))
		xs[pos.X] = true
		ys[pos.Y] = true
	}

	var mpg mapping
	mpg.xPosToIndex = breaksToPositions(xs)
	mpg.yPosToIndex = breaksToPositions(ys)
	mpg.xIndexToPos = flip(mpg.xPosToIndex)
	mpg.yIndexToPos = flip(mpg.yPosToIndex)
	return mpg
}

func walk2(steps []step, mpg mapping) (map[geom.Vec2]string, geom.Vec2, geom.Vec2, geom.Vec2, int) {
	perimeter := 0
	m := make(map[geom.Vec2]string)
	posReal := geom.V2(0, 0)
	posIndexed := mpg.realToIndexed(posReal)
	printf("start: %s→%s\n", posReal, posIndexed)
	min := posIndexed
	max := posIndexed
	seed := posIndexed.SE()

	for _, step := range steps {
		perimeter += step.distance
		dir := step.dir
		newReal := posReal.Add(dir.Mul(step.distance))
		newIndexed := mpg.realToIndexed(newReal)
		// printf("going from %s→%s to %s→%s:\n", posReal, posIndexed, newReal, newIndexed)
		posReal = newReal

		for posIndexed != newIndexed {
			posIndexed = posIndexed.Add(dir)
			m[posIndexed] = "#"
			posIndexed = posIndexed.Add(dir)
			m[posIndexed] = "#"

			if posIndexed.Y < min.Y || (posIndexed.Y == min.Y && posIndexed.X < min.X) {
				seed = posIndexed.SE()
			}

			min = geom.Min2(min, posIndexed)
			max = geom.Max2(max, posIndexed)

			// printf(" at %s\n", posIndexed)
		}
	}

	printf("Final pos=%s→%s\n", posReal, posIndexed)

	return m, min, max, seed, perimeter
}

func fill(m map[geom.Vec2]string, min geom.Vec2) {
	queue := []geom.Vec2{min.SE()}
	size := 1
	for size > 0 {
		pos := queue[size-1]
		queue = queue[:size-1]
		size--

		if m[pos] != "" {
			continue
		}

		m[pos] = "_"
		for _, dir := range geom.Compass4 {
			queue = append(queue, pos.Add(dir))
			size++
		}
	}

	return
}

func display(m map[geom.Vec2]string, min, max geom.Vec2) {
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			s := m[geom.Vec2{X: x, Y: y}]
			if s == "" {
				s = "."
			}
			printf("%s", s)

		}
		printf("\n")
	}

}

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func part1(inputs []string) (int, error) {
	steps, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	m, min, max, seed := walk(steps)
	fill(m, seed)

	// display(m, min, max)

	count := 0
	for y := min.Y; y <= max.Y; y += 2 {
		for x := min.X; x <= max.X; x += 2 {
			if m[geom.Vec2{X: x, Y: y}] != "" {
				count++
			}
		}
	}

	return count, nil
}

func part2(inputs []string) (int, error) {
	steps, err := parse2(inputs)
	if err != nil {
		return 0, err
	}
	mpg := findBreaks(steps)
	m, min, max, seed, perimeter := walk2(steps, mpg)
	fill(m, seed)

	if max.X < 100 {
		display(m, min, max)
	}

	total := perimeter

	for y := min.Y + 1; y < max.Y; y += 2 {
		for x := min.X + 1; x < max.X; x += 2 {
			pos := geom.V2(x, y)
			if m[pos] != "_" {
				continue
			}

			countWest := m[pos.W()] == "_"
			countNorth := m[pos.N()] == "_"
			countNorthWest := m[pos.NW()] == "_"

			tl := mpg.indexedToReal(pos.NW())
			br := mpg.indexedToReal(pos.SE())

			width := br.X - tl.X - 1
			height := br.Y - tl.Y - 1

			area := width * height
			if countWest {
				area += height
			}
			if countNorth {
				area += width
			}
			if countNorthWest {
				area += 1
			}

			total += area
		}
	}

	return total, nil
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
