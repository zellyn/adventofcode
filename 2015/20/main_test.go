package main

import (
	"fmt"
	"testing"
)

func TestScore(t *testing.T) {
	testdata := []struct {
		house int
		want  int
	}{
		{
			house: 9,
			want:  130,
		},
	}

	for _, tt := range testdata {
		t.Run(fmt.Sprintf("%d", tt.house), func(t *testing.T) {
			got := score(tt.house)
			if got != tt.want {
				t.Errorf("Want score(%d)=%d; got %d", tt.house, tt.want, got)
			}
		})
	}
}

func TestScore50(t *testing.T) {
	got := score50(52)
	want := 1067
	if got != want {
		t.Errorf("want score50(52)=%d; got %d", want, got)
	}
}

func TestParts(t *testing.T) {
	testdata := []struct {
		min   int
		want1 int
		want2 int
	}{
		{
			min:   130,
			want1: 8,
			want2: 6,
		},
		{
			min:   34000000,
			want1: 786240,
			want2: 831600,
		},
	}

	for _, tt := range testdata {
		t.Run(fmt.Sprintf("%d", tt.min), func(t *testing.T) {
			got1 := minHouse(tt.min)
			if got1 != tt.want1 {
				t.Errorf("Want minHouse(%d)=%d; got %d", tt.min, tt.want1, got1)
			}
			got2 := minHouse50(tt.min)
			if got2 != tt.want2 {
				t.Errorf("Want minHouse50(%d)=%d; got %d", tt.min, tt.want2, got2)
			}
		})
	}
}
