package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
2333133121414131402
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example",
			input: example,
			want:  1928,
		},
		{
			name:  "input",
			input: input,
			want:  6330095022244,
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
			want:  2858,
		},
		{
			name:  "input",
			input: input,
			want:  6359491814941,
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
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestFindGap(t *testing.T) {
	//           1         2         3         4
	// 0123456789012345678901234567890123456789012
	// 00...111....2...333.44.5555.6666.777.888899
	testdata := []struct {
		size int
		want int
	}{
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 8},
		{5, -1},
	}
	disk := parse("2334133121414131402")

	for _, tt := range testdata {
		if got := findGap(disk, tt.size); got != tt.want {
			t.Errorf("want findGap(disk, %d) == %d; got %d", tt.size, tt.want, got)
		}
	}
}
