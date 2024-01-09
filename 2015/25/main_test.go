package main

import "testing"

const inputX = 3019
const inputY = 3010

func TestIndex(t *testing.T) {
	testdata := []struct {
		x    int
		y    int
		want int
	}{
		{x: 1, y: 1, want: 1},
		{x: 6, y: 1, want: 21},
		{x: 4, y: 3, want: 19},
		{x: 1, y: 6, want: 16},
		{x: inputX, y: inputY, want: 18168397},
	}

	for _, tt := range testdata {
		got := index(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("Want index(%d,%d)=%d; got %d", tt.x, tt.y, tt.want, got)
		}
	}
}

func TestPart1(t *testing.T) {
	testdata := []struct {
		x    int
		y    int
		want int
	}{
		{x: 1, y: 1, want: 20151125},
		{x: 6, y: 1, want: 33511524},
		{x: 4, y: 3, want: 7981243},
		{x: 1, y: 6, want: 33071741},
		{x: 6, y: 6, want: 27995004},
		{x: inputX, y: inputY, want: 8997277},
	}

	for _, tt := range testdata {
		got := value(tt.x, tt.y)
		if got != tt.want {
			t.Errorf("Want value(%d,%d)=%d; got %d", tt.x, tt.y, tt.want, got)
		}
	}
}
