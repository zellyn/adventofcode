package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example4 = `deal into new stack
cut -2
deal with increment 7
cut 8
cut -4
deal with increment 7
cut 3
deal with increment 9
deal with increment 3
cut -1`

var inputLines = strings.Split(util.MustReadFileString("input"), "\n")

func TestPart1(t *testing.T) {
	prod, sum, err := shuffle(10007, inputLines)
	if err != nil {
		t.Fatal(err)
	}
	got := (2019*prod + sum) % 10007
	want := 1822
	if got != want {
		t.Errorf("Want %d; got %d", want, got)
	}
}

func TestPart2(t *testing.T) {
	size := 119315717514047
	times := 101741582076661
	prod, sum, err := shuffle(size, inputLines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := runBackward(prod, sum, times, 2020, size)
	if err != nil {
		t.Fatal(err)
	}
	want := 49174686993380
	if got != want {
		t.Errorf("Want %d; got %d", want, got)
	}

	forward, err := runForward(prod, sum, times, got, size)
	if err != nil {
		t.Error(err)
	}
	if forward != 2020 {
		t.Errorf("Want 2020 when I run it forward again; got %d", forward)
	}
}

func TestForwardAndBackward(t *testing.T) {
	testdata := []struct {
		lines   []string
		size    int
		card    int
		times   int
		forward int
	}{
		{
			lines:   strings.Split(example4, "\n"),
			size:    11,
			card:    3,
			times:   1,
			forward: 5,
		},
		{
			lines:   strings.Split(example4, "\n"),
			size:    11,
			card:    3,
			times:   2,
			forward: 10,
		},
		{
			lines:   strings.Split(example4, "\n"),
			size:    11,
			card:    0,
			times:   99,
			forward: 1,
		},
	}

	for i, tt := range testdata {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			prod, sum, err := shuffle(tt.size, tt.lines)
			if err != nil {
				t.Fatal(err)
			}
			pos := tt.card
			for i := 0; i < tt.times; i++ {
				pos = (pos*prod + sum) % tt.size
			}
			pos = (pos + tt.size) % tt.size

			if pos != tt.forward {
				t.Errorf("testcase is wrong: want pos=%d; got %d", tt.forward, pos)
			}

			forward, err := runForward(prod, sum, tt.times, tt.card, tt.size)
			if err != nil {
				t.Fatal(err)
			}

			if forward != tt.forward {
				t.Errorf("want runForward(%d, %d, %d, %d, %d)=%d; got %d", prod, sum, tt.times, tt.card, tt.size, tt.forward, forward)
			}

			backward, err := runBackward(prod, sum, tt.times, tt.forward, tt.size)
			if err != nil {
				t.Fatal(err)
			}
			if backward != tt.card {
				t.Errorf("want runBackward(%d, %d, %d, %d, %d)=%d; got %d", prod, sum, tt.times, tt.forward, tt.size, tt.card, backward)
			}
		})
	}
}
