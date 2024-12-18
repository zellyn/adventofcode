package main

import (
	"testing"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

var example = util.TrimmedLines(`
5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0
`)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		steps int
		size  geom.Vec2
		want  int
	}{
		{
			name:  "example",
			input: example,
			steps: 12,
			size:  geom.V2(7, 7),
			want:  22,
		},
		{
			name:  "input",
			input: input,
			steps: 1024,
			size:  geom.V2(71, 71),
			want:  248,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input, tt.steps, tt.size)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input, tt.steps, tt.size)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		size  geom.Vec2
		steps int
		want  geom.Vec2
	}{
		{
			name:  "example",
			input: example,
			size:  geom.V2(7, 7),
			want:  geom.V2(6, 1),
		},
		{
			name:  "input",
			input: input,
			size:  geom.V2(71, 71),
			want:  geom.V2(32, 55),
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input, tt.size)
			if err != nil {
				t.Error(err)
				return
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input, tt.size)=%s; got %s", tt.want, got)
			}
		})
	}
}
