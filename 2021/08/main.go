package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/zellyn/adventofcode/myslices"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

/*
Notes:

Actual: abcefg cf acdeg acdfg bcdf abdfg abdefg acf abcdefg abcdfg

   4 e
   6 b
   7 d
   7 g
   8 a
   8 c
   9 f

Frequencies give you e,b,f
One ("cf") gives you c, hence a
Four ("bcdf") gives you d, hence g



dbcfeag cgaed fe bfgad aefcdb efa efgda gcef dcaebg dfeagc


   4 b → e
   6 c → b
   7 d → g
   7 g → d
   8 a → a
   8 f → c
   9 e → f
*/

var segmentsToDigit = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

type entry struct {
	segmentStrings []string
	codeStrings    []string
	mapping        map[rune]rune
}

func (e *entry) getMapping() map[rune]rune {
	if e.mapping != nil {
		return e.mapping
	}

	counts := make(map[rune]int)
	var one, four string
	mapping := make(map[rune]rune)

	for _, s := range e.segmentStrings {
		if len(s) == 2 {
			one = s
		} else if len(s) == 4 {
			four = s
		}
		for _, r := range s {
			counts[r]++
		}
	}

	sevens, eights := "", ""
	for r, count := range counts {
		switch count {
		case 4:
			mapping[r] = 'e'
		case 6:
			mapping[r] = 'b'
		case 7:
			sevens = sevens + string(r)
		case 8:
			eights = eights + string(r)
		case 9:
			mapping[r] = 'f'
		}
	}

	for _, r := range one {
		if mapping[r] == 0 {
			mapping[r] = 'c'
			for _, rr := range eights {
				if rr != r {
					mapping[rr] = 'a'
				}
			}
		}
	}

	for _, r := range four {
		if mapping[r] == 0 {
			mapping[r] = 'd'
			for _, rr := range sevens {
				if rr != r {
					mapping[rr] = 'g'
				}
			}
		}
	}

	e.mapping = mapping
	// for k, v := range mapping {
	// 	printf("%c → %c\n", k, v)
	// }
	return mapping
}

func (e *entry) translate(s string) int {
	mapping := e.getMapping()

	runes := util.Map([]rune(s), func(r rune) rune { return mapping[r] })
	slices.Sort(runes)
	return segmentsToDigit[string(runes)]
}

func (e *entry) codeDigits() []int {
	return util.Map(e.codeStrings, e.translate)
}

func parseOne(s string) entry {
	a, b, _ := strings.Cut(s, " | ")
	return entry{
		segmentStrings: strings.Split(a, " "),
		codeStrings:    strings.Split(b, " "),
	}
}

func part1(inputs []string, digits []int) (int, error) {
	wantDigits := myslices.ToSet(digits)

	entries := util.Map(inputs, parseOne)

	count := 0
	for _, e := range entries {
		for _, digit := range e.codeDigits() {
			if wantDigits[digit] {
				count++
			}
		}
	}

	return count, nil
}

func part2(inputs []string) (int, error) {
	entries := util.Map(inputs, parseOne)

	sum := 0
	for _, e := range entries {
		num := 0
		for _, digit := range e.codeDigits() {
			num = num*10 + digit
		}
		sum += num
	}

	return sum, nil
}

func run() error {
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
