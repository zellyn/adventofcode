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
		want  int
	}{
		{
			name:  "example1",
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  7,
		},
		{
			name:  "example2",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
		{
			name:  "example3",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  6,
		},
		{
			name:  "example4",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  10,
		},
		{
			name:  "example5",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
		},
		{
			name:  "input",
			input: input,
			want:  1623,
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
		name  string
		input string
		want  int
	}{
		{
			name:  "example1",
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  19,
		},
		{
			name:  "example2",
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  23,
		},
		{
			name:  "example3",
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  23,
		},
		{
			name:  "example4",
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  29,
		},
		{
			name:  "example5",
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  26,
		},
		{
			name:  "input",
			input: input,
			want:  3774,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
