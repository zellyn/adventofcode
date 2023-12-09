package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
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
			want:  6440,
		},
		{
			name:  "input",
			input: input,
			want:  249726565,
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
			want:  5905,
		},
		{
			name:  "input",
			input: input,
			want:  251135960,
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

func TestLess(t *testing.T) {
	testdata := []struct {
		card1 string
		card2 string
		joker bool
		want  bool
	}{
		{card1: "KTJJT", card2: "KK677", joker: false, want: true},
		{card1: "KK677", card2: "KTJJT", joker: false, want: false},

		{card1: "KTJJT", card2: "KK677", joker: true, want: false},
		{card1: "KK677", card2: "KTJJT", joker: true, want: true},
	}

	for _, tt := range testdata {
		t.Run(tt.card1+" vs "+tt.card2, func(t *testing.T) {
			h1, err := parseHand(util.StringsAndInts{
				Strings: []string{tt.card1},
				Ints:    []int{1},
			})
			if err != nil {
				t.Fatal(err)
			}
			h2, err := parseHand(util.StringsAndInts{
				Strings: []string{tt.card2},
				Ints:    []int{1},
			})
			if err != nil {
				t.Fatal(err)
			}

			got := h1.less(h2, tt.joker)

			if got != tt.want {
				t.Errorf("Want %v < %v == %v (with jokers:%v; got %v", h1, h2, tt.joker, tt.want, got)
			}
		})
	}
}
