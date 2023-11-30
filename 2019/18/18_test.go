package main

import (
	"testing"
)

var testdata = []struct {
	input string
	want  int
}{
	{
		input: "example1",
		want:  8,
	},
	{
		input: "example2",
		want:  86,
	},
	{
		input: "example3",
		want:  132,
	},
	{
		input: "example4",
		want:  136,
	},
	{
		input: "example5",
		want:  81,
	},
	{
		input: "input",
		want:  5858,
	},
}

func TestBest(t *testing.T) {
	for _, tt := range testdata {
		t.Run(tt.input, func(t *testing.T) {
			s, err := newState(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := s.best(false)
			if got != tt.want {
				t.Errorf("want best() for %q == %d; got %d", tt.input, tt.want, got)
			}
		})
	}
}

var testdata4 = []struct {
	input string
	want  int
}{
	{
		input: "example6",
		want:  8,
	},
	{
		input: "example7",
		want:  24,
	},
	{
		input: "example8",
		want:  72,
	},
	{
		input: "input2",
		want:  2144,
	},
}

func TestBest4(t *testing.T) {
	for _, tt := range testdata4 {
		t.Run(tt.input, func(t *testing.T) {
			s, err := newState(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			got := s.best4(false)
			if got != tt.want {
				t.Errorf("want best4() for %q == %d; got %d", tt.input, tt.want, got)
			}
		})
	}
}
