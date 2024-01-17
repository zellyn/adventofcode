package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
3,4,3,1,2
`)

var input = util.MustReadLines("input")

func TestParts(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		days  int
		want  int
	}{
		{
			name:  "example-18",
			input: example,
			days:  18,
			want:  26,
		},
		{
			name:  "example-80",
			input: example,
			days:  80,
			want:  5934,
		},
		{
			name:  "input-80",
			input: input,
			days:  80,
			want:  352195,
		},
		{
			name:  "example-256",
			input: example,
			days:  256,
			want:  26984457539,
		},
		{
			name:  "input-256",
			input: input,
			days:  256,
			want:  1600306001288,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.days)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input, tt.days)=%d; got %d", tt.want, got)
			}
		})
	}
}
