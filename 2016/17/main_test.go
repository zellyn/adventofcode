package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var input = util.MustReadFileString("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example1",
			input: "ihgpwlah",
			want:  "DDRRRD",
		},
		{
			name:  "example2",
			input: "kglvqrro",
			want:  "DDUDRLRRUDRD",
		},
		{
			name:  "example3",
			input: "ulqzkmiv",
			want:  "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
		},
		{
			name:  "input",
			input: input,
			want:  "DDURRLRRDD",
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
				return
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
		want  int
	}{
		// {
		// 	name:  "example1",
		// 	input: "ihgpwlah",
		// 	want:  370,
		// },
		// {
		// 	name:  "example2",
		// 	input: "kglvqrro",
		// 	want:  492,
		// },
		// {
		// 	name:  "example3",
		// 	input: "ulqzkmiv",
		// 	want:  830,
		// },
		{
			name:  "input",
			input: input,
			want:  436,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
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
