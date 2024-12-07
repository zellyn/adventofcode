package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
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
			want:  3749,
		},
		{
			name:  "input",
			input: input,
			want:  2654749936343,
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

func TestJoin(t *testing.T) {
	testdata := []struct {
		a, b int
		want int
	}{
		{11, 8, 118},
		{50, 4, 504},
		{5, 404, 5404},
	}

	for _, tt := range testdata {
		if got := join(tt.a, tt.b); got != tt.want {
			t.Errorf("want join(%d, %d) == %d; got %d", tt.a, tt.b, tt.want, got)
		}
	}
}

func TestCanMake2(t *testing.T) {
	got := canMake2(7290, 6, []int{8, 6, 15})
	if !got {
		t.Errorf("want canMake(7290, 6, []int{8, 6, 15}) == true; got false")
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
			want:  11387,
		},
		{
			name:  "input",
			input: input,
			want:  124060392153684,
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
