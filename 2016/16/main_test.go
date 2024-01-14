package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var input = util.MustReadFileString("input")

func TestParts(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		size  int
		want  string
	}{
		{
			name:  "example",
			input: "10000",
			size:  20,
			want:  "01100",
		},
		{
			name:  "input-part1",
			input: input,
			size:  272,
			want:  "11100111011101111",
		},
		{
			name:  "input-part2",
			input: input,
			size:  35651584,
			want:  "10001110010000110",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.size)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input, tt.size)=%q; got %q", tt.want, got)
			}
		})
	}
}
