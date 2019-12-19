package main

import "testing"

func TestFillWith(t *testing.T) {
	a := []string{"L", "10", "R"}
	b := []string{"6"}
	c := []string{"R", "12", "L", "6"}
	steps := []string{"L", "10", "R", "6", "R", "12", "L", "6", "R", "12", "L", "6", "L", "10", "R", "6"}
	want := []string{"A", "B", "C", "C", "A", "B"}
	ways := fillWith(steps, [][]string{a, b, c})
	if len(ways) == 0 {
		t.Errorf("want len(ways)>0")
	}
	found := false
	for _, way := range ways {
		if eql(way, want) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Wanted to find %v in ways: %v", want, ways)
	}
}
