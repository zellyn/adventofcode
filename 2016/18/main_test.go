package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
example_input
`)

var input = util.MustReadFileString("input")

func TestParts(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		rows  int
		want  int
	}{
		{
			name:  "example1",
			input: "..^^.",
			rows:  3,
			want:  6,
		},
		{
			name:  "example2",
			input: ".^^.^.^^^^",
			rows:  10,
			want:  38,
		},
		{
			name:  "input-40",
			input: input,
			rows:  40,
			want:  1987,
		},
		{
			name:  "input-400000",
			input: input,
			rows:  400000,
			want:  19984714,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.rows)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input, tt.rows)=%d; got %d", tt.want, got)
			}
		})
	}
}
