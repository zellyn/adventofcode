package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name     string
		input    []string
		prefix   int
		want     int
		weakWant int
	}{
		{
			name:     "example",
			input:    example,
			prefix:   5,
			want:     127,
			weakWant: 62,
		},
		{
			name:     "input",
			input:    input,
			prefix:   25,
			want:     167829540,
			weakWant: 28045630,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := firstInvalid(tt.input, tt.prefix)
			if err != nil {
				t.Error(err)
			}
			if got != tt.want {
				t.Errorf("Want firstInvalid(tt.input. %d)=%d; got %d", tt.prefix, tt.want, got)
			}

			weakGot, err := weakness(tt.input, tt.prefix)
			if err != nil {
				t.Error(err)
			}
			if weakGot != tt.weakWant {
				t.Errorf("Want weakness(tt.input. %d)=%d; got %d", tt.prefix, tt.weakWant, weakGot)
			}
		})
	}
}
