package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.
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
			want:  11,
		},
		{
			name:  "input",
			input: input,
			want:  47,
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
			name:  "input",
			input: input,
			want:  71,
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

func TestValid(t *testing.T) {
	testdata := []struct {
		parts     []string
		wantValid bool
	}{
		{parts: []string{"AaM"}, wantValid: true},
		{parts: []string{"AaG"}, wantValid: true},
		{parts: []string{"AaM", "AaG"}, wantValid: true},
		{parts: []string{"AaM", "BbM"}, wantValid: true},
		{parts: []string{"AaM", "AaG", "BbG"}, wantValid: true},
		{parts: []string{"AaM", "AaG", "BbG", "BbM"}, wantValid: true},
		{parts: []string{"AaM", "AaG", "BbG", "BbM", "CcG"}, wantValid: true},

		{parts: []string{"AaM", "BbG"}, wantValid: false},
	}

	for _, tt := range testdata {
		if gotValid := valid(tt.parts); gotValid != tt.wantValid {
			t.Errorf("want valid(%v)==%v; got %v", tt.parts, tt.wantValid, gotValid)
		}
	}

}
