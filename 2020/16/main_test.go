package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = util.TrimmedLines(`
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`)

var example2 = util.TrimmedLines(`
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
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
			want:  71,
		},
		{
			name:  "input",
			input: input,
			want:  20231,
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
		name   string
		input  []string
		prefix string
		want   int
	}{
		{
			name:   "example1",
			input:  example1,
			prefix: "seat",
			want:   14,
		},
		{
			name:   "example2",
			input:  example2,
			prefix: "seat",
			want:   13,
		},
		{
			name:   "input",
			input:  input,
			prefix: "departure",
			want:   1940065747861,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.prefix)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input, %q)=%d; got %d", tt.prefix, tt.want, got)
			}
		})
	}
}
