package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
cpy 2 a
tgl a
tgl a
tgl a
cpy 1 a
dec a
dec a
`)

var input = util.MustReadLines("input")

func TestParts(t *testing.T) {
	testdata := []struct {
		name     string
		input    []string
		initialA int
		want     int
	}{
		{
			name:     "example",
			input:    example,
			initialA: 0,
			want:     3,
		},
		{
			name:     "input-7",
			input:    input,
			initialA: 7,
			want:     10500,
		},
		{
			name:     "input-12",
			input:    input,
			initialA: 12,
			want:     479007060,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.initialA)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input, tt.initialA)=%d; got %d", tt.want, got)
			}
		})
	}
}
