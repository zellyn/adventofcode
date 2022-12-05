package main

import (
	"strings"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = strings.Trim(`
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`, "\r\n")

var input = util.MustReadFileString("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: example,
			want:  "CMZ",
		},
		{
			name:  "input",
			input: input,
			want:  "RNZLFZSJH",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input)=%q; got %q", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: example,
			want:  "MCD",
		},
		{
			name:  "input",
			input: input,
			want:  "CNSFCGJSM",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%q; got %q", tt.want, got)
			}
		})
	}
}
