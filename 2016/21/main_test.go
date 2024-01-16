package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name    string
		input   []string
		initial string
		want    string
	}{
		{
			name:    "example",
			input:   example,
			initial: "abcde",
			want:    "decab",
		},
		{
			name:    "input",
			input:   input,
			initial: "abcdefgh",
			want:    "bdfhgeca",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.initial)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input, tt.initial)=%q; got %q", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name    string
		input   []string
		initial string
		want    string
	}{
		{
			name:    "input",
			input:   input,
			initial: "bdfhgeca",
			want:    "abcdefgh",
		},
		{
			name:    "input",
			input:   input,
			initial: "fbgdceah",
			want:    "gdfcabeh",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.initial)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input, tt.initial)=%q; got %q", tt.want, got)
			}
		})
	}
}
