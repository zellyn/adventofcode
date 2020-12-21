package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc
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
			want:  2,
		},
		{
			name:  "input",
			input: input,
			want:  645,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valid1Count(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Want valid1Count(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestValid2(t *testing.T) {
	testdata := []struct {
		input string
		want  bool
	}{
		{
			input: "1-3 a: abcde",
			want:  true,
		},
		{
			input: "1-3 b: cdefg",
			want:  false,
		},
		{
			input: "2-9 c: ccccccccc",
			want:  false,
		},
	}

	for _, tt := range testdata {
		got, err := valid2(tt.input)
		if err != nil {
			t.Fatal(err)
		}
		if got != tt.want {
			t.Errorf("Want valid2(%q)=%v; got %v", tt.input, tt.want, got)
		}

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
			want:  1,
		},
		{
			name:  "input",
			input: input,
			want:  737,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := valid2Count(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("Want valid2Count(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
