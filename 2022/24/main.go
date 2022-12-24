package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/math"
)

var oneone = geom.Vec2{X: 1, Y: 1}

func printf(format string, args ...any) {
	// fmt.Printf(format, args...)
}

var symToDir = map[rune]geom.Vec2{
	'>': {X: 1, Y: 0},
	'<': {X: -1, Y: 0},
	'^': {X: 0, Y: -1},
	'v': {X: 0, Y: 1},
}

var dirToSym = map[geom.Vec2]rune{
	{X: 1, Y: 0}:  '>',
	{X: -1, Y: 0}: '<',
	{X: 0, Y: -1}: '^',
	{X: 0, Y: 1}:  'v',
}

type blizzard struct {
	pos geom.Vec2
	inc geom.Vec2
}

func parseMap(m charmap.M) (geom.Rect, []blizzard) {
	min, max := m.MinMax()
	min = min.Add(oneone)
	max = max.Sub(oneone)
	rect := geom.MakeRect(min, max)
	var blizzards []blizzard
	for pos, c := range m {
		switch c {
		case '<', '>', '^', 'v':
			blizzards = append(blizzards, blizzard{pos: pos, inc: symToDir[c]})
		}
	}

	return rect, blizzards
}

func generateFields(field geom.Rect, blizzards []blizzard, count int) []charmap.M {
	var result []charmap.M

	walls := charmap.Empty()
	walls.DrawLine(field.Min.Add(geom.NW), field.BL().Add(geom.SW), '#')
	walls.DrawLine(field.Min.Add(geom.NE), field.TR().Add(geom.NE), '#')
	walls.DrawLine(field.BL().Add(geom.SW), field.Max.Add(geom.SW), '#')
	walls.DrawLine(field.TR().Add(geom.NE), field.Max.Add(geom.SE), '#')
	printf("%s\n", walls.AsString('.'))
	printf("field: %s, width=%d, height=%d\n", field, field.Width(), field.Height())
	width, height := field.Width(), field.Height()
	for i := 0; i < count; i++ {
		pic := walls.Copy()
		for j, b := range blizzards {
			pos := b.pos
			pos = pos.Add(b.inc)
			pos.X = (pos.X + width) % width
			pos.Y = (pos.Y + height) % height
			blizzards[j] = blizzard{pos: pos, inc: b.inc}

			switch pic[pos] {
			case 0:
				pic[pos] = dirToSym[b.inc]
			case '>', '<', '^', 'v':
				pic[pos] = '2'
			case '0', '1', '2', '3':
				pic[pos] += 1
			default:
				panic(fmt.Sprintf("weird char at pic[%s]: %c", pos, pic[pos]))
			}
		}
		result = append(result, pic)
	}

	_ = walls

	return result
}

func bfs(area geom.Rect, fields []charmap.M, pos geom.Vec2, part2 bool) int {
	phase := 3
	if part2 {
		phase = 1
	}
	start := area.Min.Add(geom.N)
	end := area.Max.Add(geom.S)
	goal := end
	current := []geom.Vec2{pos}
OUTER:
	for t := 0; ; t++ {
		seen := make(map[geom.Vec2]bool)
		next := make([]geom.Vec2, 0, len(current)*5)
		field := fields[t%len(fields)]
		for _, pos := range current {
			if !seen[pos] && !field.Has(pos) {
				next = append(next, pos)
				seen[pos] = true
			}
			for _, n := range pos.Neighbors4() {
				if n == goal {
					if phase == 3 {
						return t + 1
					}
					phase++
					goal = end
					if phase == 2 {
						goal = start
					}
					current = []geom.Vec2{n}
					continue OUTER
				}
				if !seen[n] && area.Contains(n) && !field.Has(n) {
					next = append(next, n)
					seen[n] = true
				}
			}
		}
		current = next
	}
}

func doit(inputs []string, part2 bool) (int, error) {
	m := charmap.ParseWithBackground(inputs, '.')
	m = m.Translated(geom.NW)
	area, blizzards := parseMap(m)
	lcm := math.LCM(area.Width(), area.Height())
	fields := generateFields(area, blizzards, lcm)
	printf("len(fields)=%d\n", len(fields))
	//	for i, field := range fields {
	//		printf("Step %d:\n%s\n\n", i+1, field.AsString('.'))
	//	}
	start := area.Min.Add(geom.N)
	steps := bfs(area, fields, start, part2)
	return steps, nil
}

func part1(inputs []string) (int, error) {
	return doit(inputs, false)
}

func part2(inputs []string) (int, error) {
	return doit(inputs, true)
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
