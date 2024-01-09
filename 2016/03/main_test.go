package main

import (
	"testing"
)

func TestParts(t *testing.T) {
	grid, err := readGrid("input")
	if err != nil {
		t.Fatal(err)
	}
	got1 := countPossible(grid)
	want1 := 993
	if got1 != want1 {
		t.Errorf("want countPossible(grid)=%d; got %d", want1, got1)
	}
	got2 := countPossible2(grid)
	want2 := 1849
	if got2 != want2 {
		t.Errorf("want countPossible2(grid)=%d; got %d", want2, got2)
	}
}
