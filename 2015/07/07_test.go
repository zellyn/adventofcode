package main

import (
	"strings"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

func TestCircuit(t *testing.T) {
	example1 := `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`
	testdata := []struct {
		name  string
		input string
		want  map[string]uint16
	}{
		{
			name:  "example",
			input: example1,
			want: map[string]uint16{
				"d": 72,
				"e": 507,
				"f": 492,
				"g": 114,
				"h": 65412,
				"i": 65079,
				"x": 123,
				"y": 456,
			},
		},
		{
			name:  "part 1",
			input: util.MustReadFileString("input"),
			want: map[string]uint16{
				"a": 3176,
			},
		},
		{
			name:  "part 2",
			input: strings.Replace(util.MustReadFileString("input"), "44430 -> b", "3176 -> b", 1),
			want: map[string]uint16{
				"a": 14710,
			},
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			parsed, err := parseInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			memo := map[string]uint16{}
			for name, want := range tt.want {
				got, err := eval(name, parsed, memo)
				if err != nil {
					t.Errorf("cannot evaluate %q: %v", name, err)
				} else if got != want {
					t.Errorf("want eval(%q, parsed, memo)=%d; got %d", name, want, got)
				}
			}
		})
	}
}
