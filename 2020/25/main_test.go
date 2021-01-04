package main

import (
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var input = util.MustReadLines("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name string
		door int
		card int
		want int
	}{
		{
			name: "example",
			card: 5764801,
			door: 17807724,
			want: 14897079,
		},
		{
			name: "input",
			card: 11239946,
			door: 10464955,
			want: 711945,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got := part1(tt.card, tt.door)
			if got != tt.want {
				t.Errorf("Want part1(%d, %d)=%d; got %d", tt.card, tt.door, tt.want, got)
			}
		})
	}
}
