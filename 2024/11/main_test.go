package main

import (
	"strings"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
125 17
`)

var input = strings.TrimSpace(util.MustReadFileString("input"))

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		iters int
		want  int
	}{
		{
			name:  "first example",
			input: "0 1 10 99 999",
			iters: 1,
			want:  7,
		},
		{
			name:  "second example, 1 step",
			input: "125 17",
			iters: 1,
			want:  3,
		},
		{
			name:  "second example, 2 steps",
			input: "125 17",
			iters: 2,
			want:  4,
		},
		{
			name:  "second example, 6 steps",
			input: "125 17",
			iters: 6,
			want:  22,
		},
		{
			name:  "input-25",
			input: input,
			iters: 25,
			want:  200446,
		},
		{
			name:  "input-75",
			input: input,
			iters: 75,
			want:  238317474993392,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.iters)
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
