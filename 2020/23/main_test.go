package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		moves int
		want  string
	}{
		{
			name:  "example",
			input: "389125467",
			moves: 10,
			want:  "92658374",
		},
		{
			name:  "example",
			input: "389125467",
			moves: 100,
			want:  "67384529",
		},
		{
			name:  "example",
			input: "389125467",
			moves: 10000,
			want:  "64572893",
		},
		{
			name:  "input",
			input: "962713854",
			moves: 100,
			want:  "65432978",
		},
	}

	for _, tt := range testdata {
		t.Run(fmt.Sprintf("%s-%d", tt.name, tt.moves), func(t *testing.T) {
			got, err := part1(tt.input, tt.moves)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part1(%q, %d)=%q; got %q", tt.input, tt.moves, tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input string
		moves int
		want  int
	}{
		{
			name:  "example",
			input: "389125467",
			moves: 10000000,
			want:  149245887792,
		},
		{
			name:  "input",
			input: "962713854",
			moves: 10000000,
			want:  149245887792,
		},
	}

	for _, tt := range testdata {
		t.Run(fmt.Sprintf("%s-%d", tt.name, tt.moves), func(t *testing.T) {
			got, err := part2(tt.input, tt.moves)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part1(%q, %d)=%d; got %d", tt.input, tt.moves, tt.want, got)
			}
		})
	}
}

func TestStep(t *testing.T) {
	testdata := []struct {
		input []int
		pos   int
		want  []int
	}{
		{
			input: []int{3, 8, 9, 1, 2, 5, 4, 6, 7},
			pos:   0,
			want:  []int{5, 8, 2, 6, 4, 7, 3, 9, 1},
		},
		{
			input: []int{3, 2, 8, 9, 1, 5, 4, 6, 7},
			pos:   1,
			want:  []int{3, 5, 2, 6, 4, 7, 8, 9, 1},
		},
		/*
			{
				input: []int{3, 2, 5, 4, 6, 7, 8, 9, 1},
				pos:   2,
			},
			{
				input: []int{7, 2, 5, 8, 9, 1, 3, 4, 6},
				pos:   3,
			},
			{
				input: []int{3, 2, 5, 8, 4, 6, 7, 9, 1},
				pos:   4,
			},
			{
				input: []int{9, 2, 5, 8, 4, 1, 3, 6, 7},
				pos:   5,
			},
			{
				input: []int{7, 2, 5, 8, 4, 1, 9, 3, 6},
				pos:   6,
			},
			{
				input: []int{8, 3, 6, 7, 4, 1, 9, 2, 5},
				pos:   7,
			},
			{
				input: []int{7, 4, 1, 5, 8, 3, 9, 2, 6},
				pos:   8,
			},
			{
				input: []int{5, 7, 4, 1, 8, 3, 9, 2, 6},
				pos:   0,
			},
		*/
	}

	for i, tt := range testdata {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			ary, _ := toNext(tt.input, -1)
			//fmt.Printf("input: %v\nnext: %v\n", tt.input, ary)
			// next12, _ := toNext(tt.input, 12)
			// fmt.Printf("next12: %v\n", next12)
			pos := tt.input[tt.pos]
			move(ary, pos)
			if !reflect.DeepEqual(ary, tt.want) {
				t.Errorf("Want move(toNext(%v), %d)=%v; got %v", tt.input, tt.input[tt.pos], tt.want, ary)
			}
		})
	}
}
