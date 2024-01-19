package main

import (
	"reflect"
	"strconv"
	"testing"
)

var sequentialTestdata = []struct {
	input      []int64
	wantSignal int64
	sequence   []int64
}{
	{
		input:      []int64{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		wantSignal: 43210,
		sequence:   []int64{4, 3, 2, 1, 0},
	},
	{
		input:      []int64{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
		wantSignal: 54321,
		sequence:   []int64{0, 1, 2, 3, 4},
	},
	{
		input:      []int64{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
		wantSignal: 65210,
		sequence:   []int64{1, 0, 4, 3, 2},
	},
}

func TestRunSequence(t *testing.T) {
	for i, tt := range sequentialTestdata {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			signal, err := runSequence(tt.input, tt.sequence, false)
			if err != nil {
				t.Fatal(err)
			}
			if signal != tt.wantSignal {
				t.Errorf("want signal=%d; got %d", tt.wantSignal, signal)
			}
		})
	}
}

func TestFindBestSequence(t *testing.T) {
	for i, tt := range sequentialTestdata {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			signal, sequence, err := bestSequence(tt.input, false)
			if err != nil {
				t.Fatal(err)
			}
			if signal != tt.wantSignal {
				t.Errorf("want signal=%d; got %d", tt.wantSignal, signal)
			}
			if !reflect.DeepEqual(tt.sequence, tt.sequence) {
				t.Errorf("want sequence=%v; got %v", tt.sequence, sequence)
			}
		})
	}
}

var parallelTestdata = []struct {
	input      []int64
	wantSignal int64
	sequence   []int64
}{
	{
		input:      []int64{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
		wantSignal: 139629729,
		sequence:   []int64{9, 8, 7, 6, 5},
	},
	{
		input:      []int64{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10},
		wantSignal: 18216,
		sequence:   []int64{9, 7, 8, 5, 6},
	},
}

func TestRunParallelSequence(t *testing.T) {
	for i, tt := range parallelTestdata {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			signal, err := runParallelSequence(tt.input, tt.sequence, true)
			if err != nil {
				t.Fatal(err)
			}
			if signal != tt.wantSignal {
				t.Errorf("want signal=%d; got %d", tt.wantSignal, signal)
			}
		})
	}
}

func TestFindBestParallelSequence(t *testing.T) {
	for i, tt := range parallelTestdata {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			signal, sequence, err := bestParallelSequence(tt.input, false)
			if err != nil {
				t.Fatal(err)
			}
			if signal != tt.wantSignal {
				t.Errorf("want signal=%d; got %d", tt.wantSignal, signal)
			}
			if !reflect.DeepEqual(tt.sequence, tt.sequence) {
				t.Errorf("want sequence=%v; got %v", tt.sequence, sequence)
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	want := [][]int64{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}

	got := permutations([]int64{0, 1, 2})
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want permutations({0,1,2})==%v; got %v", want, got)
	}
}
