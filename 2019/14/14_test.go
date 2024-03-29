package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/zellyn/adventofcode/util"
)

var example1 = `10 ORE => 10 A
1 ORE => 1 B
7 A, 1 B => 1 C
7 A, 1 C => 1 D
7 A, 1 D => 1 E
7 A, 1 E => 1 FUEL`

var example2 = `9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL`

var example3 = `157 ORE => 5 NZVS
165 ORE => 6 DCFZ
44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL
12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ
179 ORE => 7 PSHF
177 ORE => 5 HKGWZ
7 DCFZ, 7 PSHF => 2 XJWVT
165 ORE => 2 GPVTF
3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT`

var example4 = `2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG
17 NVRVD, 3 JNWZP => 8 VPVL
53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL
22 VJHF, 37 MNCFX => 5 FWMGM
139 ORE => 4 NVRVD
144 ORE => 7 JNWZP
5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC
5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV
145 ORE => 6 MNCFX
1 NVRVD => 8 CXFTF
1 VJHF, 6 MNCFX => 4 RFSQX
176 ORE => 6 VJHF`

var example5 = `171 ORE => 8 CNZTR
7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL
114 ORE => 4 BHXH
14 VRPVC => 6 BMBT
6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL
6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT
15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW
13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW
5 BMBT => 4 WPTQ
189 ORE => 9 KTJDG
1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP
12 VRPVC, 27 CNZTR => 2 XDBXC
15 KTJDG, 12 BHXH => 5 XCVML
3 BHXH, 2 VRPVC => 7 MZWV
121 ORE => 7 VRPVC
7 XCVML => 6 RJRHP
5 BHXH, 4 VRPVC => 5 LTCX`

var mainInput = util.MustReadFileString("input")

var testdata = []struct {
	name     string
	input    string
	count    int
	maxTrill int
}{
	{
		name:  "example 1",
		input: example1,
		count: 31,
	},
	{
		name:  "example 2",
		input: example2,
		count: 165,
	},
	{
		name:     "example 3",
		input:    example3,
		count:    13312,
		maxTrill: 82892753,
	},
	{
		name:     "example 4",
		input:    example4,
		count:    180697,
		maxTrill: 5586022,
	},
	{
		name:     "example 5",
		input:    example5,
		count:    2210736,
		maxTrill: 460664,
	},
	{
		name:     "real input",
		input:    mainInput,
		count:    346961,
		maxTrill: 4065790,
	},
}

func TestCount(t *testing.T) {
	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			s, err := parseInput(tt.input)
			if err != nil {
				t.Fatal(err)
			}
			count, err := s.run(1)
			if err != nil {
				t.Fatal(err)
			}
			if count != tt.count {
				t.Errorf("want count=%d; got %d", tt.count, count)
			}
			if tt.maxTrill > 0 {
				count, err = s.calcMax(1000000000000)
				if err != nil {
					t.Fatal(err)
				}
				if count != tt.maxTrill {
					t.Errorf("want s.max(TRILLION)=%d; got %d", tt.maxTrill, count)
				}
			}
		})
	}

}

func TestParseRule(t *testing.T) {
	testdata := []struct {
		input string
		want  rule
	}{
		{
			input: "1 KZQC => 8 MPBVZ",
			want: rule{
				output: "MPBVZ",
				count:  8,
				inputs: map[string]int{
					"KZQC": 1,
				},
			},
		},
		{
			input: "1 KJCJL, 15 GTWKH => 2 SBLQ",
			want: rule{
				output: "SBLQ",
				count:  2,
				inputs: map[string]int{
					"KJCJL": 1,
					"GTWKH": 15,
				},
			},
		},
		{
			input: "1 ZDQRT, 3 CZLDF, 10 GDLFK, 1 BDRP, 10 VHMT, 6 XGVF, 1 RLFHL => 7 CVHR",
			want: rule{
				output: "CVHR",
				count:  7,
				inputs: map[string]int{
					"ZDQRT": 1,
					"CZLDF": 3,
					"GDLFK": 10,
					"BDRP":  1,
					"VHMT":  10,
					"XGVF":  6,
					"RLFHL": 1,
				},
			},
		},
	}

	for i, tt := range testdata {
		t.Run(fmt.Sprintf("%d:%s", i, tt.want.output), func(t *testing.T) {
			got, err := parseRule(tt.input)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want rule==%v; got %v", tt.want, got)
			}
		})
	}
}
