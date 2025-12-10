package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([]geom.Vec2, error) {
	return util.MapE(inputs, func(s string) (geom.Vec2, error) {
		first, second, ok := strings.Cut(s, ",")
		if !ok {
			return geom.Vec2{}, fmt.Errorf("weird input line: %q", s)
		}
		x, err := strconv.Atoi(first)
		if err != nil {
			return geom.Vec2{}, err
		}
		y, err := strconv.Atoi(second)
		if err != nil {
			return geom.Vec2{}, err
		}
		return geom.V2(x, y), nil
	})
}

func part1(inputs []string) (int, error) {
	tiles, err := parse(inputs)
	if err != nil {
		return 0, nil
	}
	most := 0

	for i, tile1 := range tiles {
		for _, tile2 := range tiles[i+1:] {
			wh := tile1.Sub(tile2).Abs()
			area := (wh.X + 1) * (wh.Y + 1)
			most = max(area, most)
		}
	}

	return most, nil
}

func maybe(p geom.Vec2, m charmap.M) {
	if m[p] == 0 {
		m[p] = '.'
	}
}

func draw(p1, p2 geom.Vec2, m charmap.M) {
	dir := p2.Sub(p1).Sgn()
	for p := p1; p != p2; p = p.Add(dir) {
		m[p] = '#'
	}
	m[p2] = '#'

	xMin := min(p1.X, p2.X) - 1
	xMax := max(p1.X, p2.X) + 1
	yMin := min(p1.Y, p2.Y) - 1
	yMax := max(p1.Y, p2.Y) + 1

	for x := xMin; x <= xMax; x++ {
		maybe(geom.Vec2{X: x, Y: yMin}, m)
		maybe(geom.Vec2{X: x, Y: yMax}, m)
	}
	for y := yMin + 1; y < yMax; y++ {
		maybe(geom.Vec2{X: xMin, Y: y}, m)
		maybe(geom.Vec2{X: xMax, Y: y}, m)
	}
}

func trace(m charmap.M) (map[int][]int, map[int][]int) {
	atX := make(map[int][]int)
	atY := make(map[int][]int)

	var start geom.Vec2
	for pos := range m {
		start = pos.WithY(-2)
		break
	}

	for m[start] == 0 {
		start = start.S()
	}

	todo := []geom.Vec2{start}

	ll := len(todo)
	for ll > 0 {
		pos := todo[ll-1]
		todo = todo[:ll-1]
		ll--

		if m[pos] != '.' {
			continue
		}
		m[pos] = '/'
		atX[pos.X] = append(atX[pos.X], pos.Y)
		atY[pos.Y] = append(atY[pos.Y], pos.X)
		todo = append(todo, pos.N(), pos.S(), pos.E(), pos.W())
		ll += 4
	}

	for _, v := range atX {
		slices.Sort(v)
	}
	for _, v := range atY {
		slices.Sort(v)
	}

	return atX, atY
}

func next(pos int, poses []int) int {
	for _, maybe := range poses {
		if maybe > pos {
			return maybe
		}
	}
	panic("not found!")
}

func safe(r geom.Rect, atX, atY map[int][]int) bool {
	if next(r.Min.X, atY[r.Min.Y]) <= r.Max.X {
		return false
	}
	if next(r.Min.X, atY[r.Max.Y]) <= r.Max.X {
		return false
	}
	if next(r.Min.Y, atX[r.Min.X]) <= r.Max.Y {
		return false
	}
	if next(r.Min.Y, atX[r.Max.X]) <= r.Max.Y {
		return false
	}

	return true
}

func part2(inputs []string) (int, error) {
	tiles, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	m := make(charmap.M)
	for i, p1 := range tiles {
		p2 := tiles[(i+1)%len(tiles)]
		draw(p1, p2, m)
	}

	atX, atY := trace(m)

	most := 0

	for i, tile1 := range tiles {
		for _, tile2 := range tiles[i+1:] {
			r := geom.MakeRect(tile1, tile2)
			wh := tile1.Sub(tile2).Abs()
			area := (wh.X + 1) * (wh.Y + 1)
			if area < most {
				continue
			}
			if !safe(r, atX, atY) {
				continue
			}
			most = area
		}
	}

	return most, nil
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
