package main

import (
	"testing"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

var example = charmap.Parse(util.TrimmedLines(`
	..##.......
	#...#...#..
	.#....#..#.
	..#.#...#.#
	.#...##..#.
	..#.##.....
	.#.#.#....#
	.#........#
	#.##...#...
	#...##....#
	.#..#...#.#`))

var input = charmap.MustRead("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		m     map[geom.Vec2]rune
		step  geom.Vec2
		steps []geom.Vec2
		want1 int
		want2 int
	}{
		{
			name:  "example",
			m:     example,
			step:  geom.Vec2{X: 3, Y: 1},
			steps: []geom.Vec2{{X: 1, Y: 1}, {X: 3, Y: 1}, {X: 5, Y: 1}, {X: 7, Y: 1}, {X: 1, Y: 2}},
			want1: 7,
			want2: 336,
		},
		{
			name:  "input",
			m:     input,
			step:  geom.Vec2{X: 3, Y: 1},
			steps: []geom.Vec2{{X: 1, Y: 1}, {X: 3, Y: 1}, {X: 5, Y: 1}, {X: 7, Y: 1}, {X: 1, Y: 2}},
			want1: 259,
			want2: 42,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got1 := treeCount(tt.m, tt.step)
			if got1 != tt.want1 {
				t.Errorf("Want treeCount(tt.m, %v)=%d; got %d", tt.step, tt.want1, got1)
			}
			got2 := multiTreeCount(tt.m, tt.steps)
			if got2 != tt.want2 {
				t.Errorf("Want multiTreeCount(tt.m, tt.steps)=%d; got %d", tt.want2, got2)
			}
		})
	}
}
