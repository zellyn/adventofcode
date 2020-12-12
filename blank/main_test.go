package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
example_input
`)

var input = ioutil.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want1 int
		want2 int
	}{
		{
			name:  "example",
			input: example,
			want1: 42,
			want2: 42,
		},
		{
			name:  "input",
			input: input,
			want1: 42,
			want2: 42,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			want1, err := part1(tt.input)
			if err != nil {
				t.Error(err)
			}

			if want1 != tt.want1 {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want1, want1)
			}

			want2, err := part2(tt.input)
			if err != nil {
				t.Error(err)
			}

			if want2 != tt.want2 {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want2, want2)
			}
		})
	}
}
