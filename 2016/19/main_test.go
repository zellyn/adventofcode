package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
example_input
`)

var input = util.MustReadSingleInt("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input int
		want  int
	}{
		{name: "example", input: 5, want: 3},
		{name: "mine-1", input: 1, want: 1},
		{name: "mine-2", input: 2, want: 1},
		{name: "mine-3", input: 3, want: 3},
		{name: "mine-4", input: 4, want: 1},
		{name: "mine-5", input: 5, want: 3},
		{name: "mine-6", input: 6, want: 5},
		{name: "mine-7", input: 7, want: 7},
		{name: "mine-8", input: 8, want: 1},
		{name: "mine-9", input: 9, want: 3},
		{name: "mine-10", input: 10, want: 5},
		{name: "mine-11", input: 11, want: 7},
		{name: "mine-12", input: 12, want: 9},
		{name: "mine-13", input: 13, want: 11},
		{name: "mine-14", input: 14, want: 13},
		{name: "mine-15", input: 15, want: 15},
		{name: "mine-16", input: 16, want: 1},
		{name: "mine-17", input: 17, want: 3},
		{name: "mine-18", input: 18, want: 5},
		{name: "mine-19", input: 19, want: 7},
		{name: "mine-20", input: 20, want: 9},
		{name: "mine-21", input: 21, want: 11},
		{name: "mine-22", input: 22, want: 13},
		{name: "mine-23", input: 23, want: 15},
		{name: "mine-24", input: 24, want: 17},
		{name: "mine-25", input: 25, want: 19},
		{name: "mine-26", input: 26, want: 21},
		{name: "mine-27", input: 27, want: 23},
		{name: "mine-28", input: 28, want: 25},
		{name: "input", input: input, want: 1834471},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
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
		name  string
		input int
		want  int
	}{
		{name: "example", input: 5, want: 2},
		{name: "input", input: input, want: 1420064},
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
