package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strconv"

	"golang.org/x/exp/maps"
)

var printf = func(string, ...any) {}

// var printf = fmt.Printf

type state struct {
	salt           string
	max            int
	extraLoops     int
	threes         map[int]map[rune]bool
	fives          map[int]map[rune]bool
	found          map[int]bool
	largestChecked map[rune]int
}

func newState(salt string) *state {
	return &state{
		salt:           salt,
		max:            -1,
		threes:         make(map[int]map[rune]bool),
		fives:          make(map[int]map[rune]bool),
		found:          make(map[int]bool),
		largestChecked: make(map[rune]int),
	}
}

func max(ints ...int) int {
	res := ints[0]
	for _, i := range ints[1:] {
		if i > res {
			res = i
		}
	}
	return res
}

func (s *state) md5(n int) string {
	md5Bytes := md5.Sum([]byte(s.salt + strconv.Itoa(n)))

	for i := 0; i < s.extraLoops; i++ {
		md5Bytes = md5.Sum([]byte(hex.EncodeToString(md5Bytes[:])))
	}
	return hex.EncodeToString(md5Bytes[:])
}

func (s *state) genNext() int {
	s.max++
	h := s.md5(s.max)
	var threes map[rune]bool
	var fives map[rune]bool

	last := ' '
	count := 0
	for _, r := range []rune(h) {
		if r == last {
			count++
		} else {
			last = r
			count = 1
		}

		if count >= 3 {
			if threes == nil {
				threes = make(map[rune]bool)
				threes[r] = true
			}

			if count >= 5 {
				if fives == nil {
					fives = make(map[rune]bool)
				}
				fives[r] = true
			}
		}
	}

	if threes != nil {
		s.threes[s.max] = threes
		if fives != nil {
			s.fives[s.max] = fives
		}
	}

	if len(fives) > 0 {
		printf("Checking for %d: %s\n", s.max, h)
	}
	for r := range fives {
		printf(" %c:\n", r)
		start := max(0, s.max-1000, s.largestChecked[r]+1)
		end := s.max
		s.largestChecked[r] = end - 1

		for i := start; i < end; i++ {
			if s.threes[i][r] {
				s.found[i] = true
				//printf("  %d: %s\n", i, s.md5(i))
				printf("  %d\n", i)
			}
		}
	}

	return s.max
}

func part1(salt string, target int) (int, error) {
	s := newState(salt)
	for len(s.found) < target {
		s.genNext()
	}
	for i := 0; i < 1000; i++ {
		s.genNext()
	}
	candidates := maps.Keys(s.found)
	sort.Ints(candidates)
	return candidates[target-1], nil
}

func part2(salt string, target int) (int, error) {
	s := newState(salt)
	s.extraLoops = 2016
	for len(s.found) < target {
		s.genNext()
	}
	for i := 0; i < 1000; i++ {
		s.genNext()
	}
	candidates := maps.Keys(s.found)
	sort.Ints(candidates)
	return candidates[target-1], nil
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
