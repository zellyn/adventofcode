package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = util.TrimmedLines(`
-L|F7
7S-7|
L|7||
-L-J|
L|-JF
`)

var example2 = util.TrimmedLines(`
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`)

var example3 = util.TrimmedLines(`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`)

var example4 = util.TrimmedLines(`
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  4,
		},
		{
			name:  "example2",
			input: example2,
			want:  8,
		},
		{
			name:  "input",
			input: input,
			want:  7066,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  1,
		},
		{
			name:  "example2",
			input: example2,
			want:  1,
		},
		{
			name:  "example3",
			input: example3,
			want:  4,
		},
		{
			name:  "example4",
			input: example4,
			want:  10,
		},
		{
			name:  "input",
			input: input,
			want:  401,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
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
