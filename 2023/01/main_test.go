package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var input = util.MustReadLines("input")

var example = util.TrimmedLines(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

var example2 = util.TrimmedLines(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  142,
		},
		{
			name:  "input",
			input: input,
			want:  54632,
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
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  142,
		},
		{
			name:  "example2",
			input: example2,
			want:  281,
		},
		{
			name:  "input",
			input: input,
			want:  54019,
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
