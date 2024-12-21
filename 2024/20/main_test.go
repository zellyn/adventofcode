package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name    string
		input   []string
		minSave int
		want    int
	}{
		{
			name:    "example",
			input:   example,
			minSave: 1,
			want:    44,
		},
		{
			name:    "input",
			input:   input,
			minSave: 100,
			want:    1355,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.minSave)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name    string
		input   []string
		minSave int
		want    int
	}{
		{
			name:    "example",
			input:   example,
			minSave: 50,
			want:    285,
		},
		{
			name:    "input",
			input:   input,
			minSave: 100,
			want:    1007335,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.minSave)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
