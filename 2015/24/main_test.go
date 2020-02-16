package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRead(t *testing.T) {
	ints, err := readInput("input")
	if err != nil {
		t.Fatal(err)
	}
	got := len(ints)
	want := 29
	if got != want {
		t.Errorf("want %d ints; got %d", want, got)
	}
}

func TestCanAddTo(t *testing.T) {
	memo := map[int]map[uint32]bool{}

	testdata := []struct {
		ints   []int
		target int
		want   bool
	}{

		{
			ints:   []int{1, 2, 3, 4, 5},
			target: 15,
			want:   true,
		},
		{
			ints:   []int{3, 7, 9, 13, 21},
			target: 11,
			want:   false,
		},

		{
			ints:   []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
			target: 20,
			want:   true,
		},
	}

	for i, tt := range testdata {
		mask := uint32(1<<len(tt.ints) - 1)
		sum := 0
		for _, i := range tt.ints {
			sum += i
		}
		got := canAddTo(tt.target, mask, sum, tt.ints, memo)
		if got != tt.want {
			t.Errorf("%d: want %v; got %v", i, tt.want, got)
			continue
		}
		got = canAddTo(tt.target, mask, sum, tt.ints, memo)
		if got != tt.want {
			t.Errorf("%d: (pass2) want %v; got %v", i, tt.want, got)
		}
	}
}

func TestWaysToAddTo(t *testing.T) {
	testdata := []struct {
		ints     []int
		target   int
		maxCount int
		want     []uint32
	}{
		{
			ints:     []int{1, 2, 3, 4, 5},
			target:   9,
			maxCount: 3,
			want:     []uint32{0b00011, 0b10101, 0b01110},
		},
		{
			ints:     []int{1, 2, 3, 4, 5},
			target:   9,
			maxCount: 2,
			want:     []uint32{0b00011},
		},
		/*
			{
				ints: []int{
					1, 2, 3, 7, 11, 13, 17, 19, 23, 31, 37, 41, 43, 47, 53, 59,
					61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113,
				},
				target:   520,
				maxCount: 9,
			},
		*/
	}

	for i, tt := range testdata {
		mask := uint32(1)<<len(tt.ints) - 1
		sum := 0
		for _, anInt := range tt.ints {
			sum += anInt
		}
		got := waysToAddTo(tt.target, mask, 0, sum, tt.ints, tt.maxCount)
		if !cmp.Equal(got, tt.want) {
			t.Errorf("%d: want %#v; got %#v", i, tt.want, got)
			continue
		}
	}
}

func TestEntanglement(t *testing.T) {
	testdata := []struct {
		ints []int
		mask uint32
		want int
	}{
		{
			ints: []int{1, 2, 3, 4, 5},
			mask: 0b11111,
			want: 120,
		},
		{
			ints: []int{1, 2, 3, 4, 5},
			mask: 0b01001,
			want: 10,
		},
		{
			ints: []int{1, 2, 3, 4, 5},
			mask: 0b00011,
			want: 20,
		},
	}

	for i, tt := range testdata {
		got := entanglement(tt.ints, tt.mask)
		if got != tt.want {
			t.Errorf("%d: want %d; got %d", i, tt.want, got)
			continue
		}
	}
}

func TestParts(t *testing.T) {
	inputInts, err := readInput("input")
	if err != nil {
		t.Fatal(err)
	}

	testdata := []struct {
		name       string
		ints       []int
		partitions int
		want       int
	}{
		{
			name:       "example1",
			ints:       []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
			partitions: 3,
			want:       99,
		},
		{
			name:       "example2",
			ints:       []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11},
			partitions: 4,
			want:       44,
		},
		{
			name:       "input-part1",
			ints:       inputInts,
			partitions: 3,
			want:       11846773891,
		},
		{
			name:       "input-part2",
			ints:       inputInts,
			partitions: 4,
			want:       80393059,
		},
	}

	_ = inputInts
	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got := best(tt.ints, tt.partitions)
			if got != tt.want {
				t.Errorf("want best(ints, %d)=%d; got %d", tt.partitions, tt.want, got)
			}
		})
	}
}
