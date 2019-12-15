package main

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/zellyn/adventofcode/2019/intcode"
)

var testdata = []struct {
	program []int64
	writes  []int64
	debug   bool
}{
	{
		program: []int64{104, 1125899906842624, 99},
		writes:  []int64{1125899906842624},
	},
	{
		program: []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		writes:  []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		debug:   true,
	},
	{
		program: []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
		writes:  []int64{1219070632396864},
	},
}

func TestBigNumbers(t *testing.T) {
	for i, tt := range testdata {
		t.Run("test "+strconv.Itoa(i), func(t *testing.T) {
			// (originalState []int64, reads []int64, debug bool) (state []int64, writes []int64, err error) {
			_, writes, err := intcode.RunProgram(tt.program, nil, tt.debug)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(writes, tt.writes) {
				t.Errorf("want writes == %v; got %v", tt.writes, writes)
			}
		})
	}
}
