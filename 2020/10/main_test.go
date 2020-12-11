package main

import (
	"testing"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/util"
)

var example1 = ioutil.MustStringsToInts(util.TrimmedLines(`
16
10
15
5
1
11
7
19
6
12
4
`))

var example2 = ioutil.MustStringsToInts(util.TrimmedLines(`
28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`))

var input = ioutil.MustReadFileInts("input")

func TestPart1(t *testing.T) {
	testdata := []struct {
		name     string
		input    []int
		wantProd int
		wantWays int
	}{
		{
			name:     "example1",
			input:    example1,
			wantProd: 35,
			wantWays: 8,
		},
		{
			name:     "example2",
			input:    example2,
			wantProd: 220,
			wantWays: 19208,
		},
		{
			name:     "input",
			input:    input,
			wantProd: 1856,
			wantWays: 42,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			gotProd := prod(tt.input)
			if gotProd != tt.wantProd {
				t.Errorf("Want prod(tt.input)=%d; got %d", tt.wantProd, gotProd)
			}

			gotWays := ways(tt.input)
			if gotWays != tt.wantWays {
				t.Errorf("Want ways(tt.input)=%d; got %d", tt.wantWays, gotWays)
			}
		})
	}
}
