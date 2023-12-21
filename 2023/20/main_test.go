package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/zellyn/adventofcode/util"
	"kr.dev/diff"
)

var example1 = util.TrimmedLines(`
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`)

var example1Pulses = [][]string{
	util.TrimmedLines(`
button -low-> broadcaster
broadcaster -low-> a
broadcaster -low-> b
broadcaster -low-> c
a -high-> b
b -high-> c
c -high-> inv
inv -low-> a
a -low-> b
b -low-> c
c -low-> inv
inv -high-> a
`),
}

var example2 = util.TrimmedLines(`
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`)

var example2Pulses = [][]string{
	util.TrimmedLines(`
button -low-> broadcaster
broadcaster -low-> a
a -high-> inv
a -high-> con
inv -low-> b
con -high-> output
b -high-> con
con -low-> output
`),
	util.TrimmedLines(`
button -low-> broadcaster
broadcaster -low-> a
a -low-> inv
a -low-> con
inv -high-> b
con -high-> output
`),
	util.TrimmedLines(`
button -low-> broadcaster
broadcaster -low-> a
a -high-> inv
a -high-> con
inv -low-> b
con -low-> output
b -low-> con
con -high-> output
`),
	util.TrimmedLines(`
button -low-> broadcaster
broadcaster -low-> a
a -low-> inv
a -low-> con
inv -high-> b
con -high-> output
`),
}

var input = util.MustReadLines("input")

func TestOperation(t *testing.T) {
	testdata := []struct {
		name       string
		input      []string
		wantPulses [][]string
		subcircuit string
	}{
		{
			name:       "example1",
			input:      example1,
			wantPulses: example1Pulses,
		},
		{
			name:       "example2",
			input:      example2,
			wantPulses: example2Pulses,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {

			circ, err := parse(tt.input)
			if err != nil {
				t.Error(err)
				return
			}

			if tt.subcircuit != "" {
				circ = circ.subcircuit(tt.subcircuit)
			}

			if tt.name == "input" {
				for _, comp := range circ.components {
					printf("%#v\n", comp)
				}
			}

			circ.debug = true

			for i, wantPulseRecord := range tt.wantPulses {
				circ.pulseRecord = circ.pulseRecord[:0]
				circ.pushButton()
				gotPulseRecord := circ.pulseRecord
				if !reflect.DeepEqual(wantPulseRecord, gotPulseRecord) {
					t.Errorf("Difference in ouput after button press %d: want:\n[\n  [%s]\n]\ngot:\n[\n  [%s]\n]", i+1,
						strings.Join(wantPulseRecord, "]\n  ["), strings.Join(gotPulseRecord, "]\n  ["))
					diff.Test(t, t.Errorf, gotPulseRecord, wantPulseRecord)
				}
			}
		})
	}
}

func TestPart1(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "example1",
			input: example1,
			want:  32000000,
		},
		{
			name:  "example2",
			input: example2,
			want:  11687500,
		},
		{
			name:  "input",
			input: input,
			want:  747304011,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part1(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part1(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testdata := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "input",
			input: input,
			want:  220366255099387,
		},
	}

	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			got, err := part2(tt.input)
			if err != nil {
				t.Error(err)
			}

			if got != tt.want {
				t.Errorf("Want part2(tt.input)=%d; got %d", tt.want, got)
			}
		})
	}
}
