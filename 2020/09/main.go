package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/ioutil"
)

type lastN struct {
	slice []int
	m     map[int]int
}

func newLastN(initial []int) lastN {

	ln := lastN{
		slice: make([]int, len(initial)),
		m:     make(map[int]int, len(initial)),
	}
	copy(ln.slice, initial)
	for _, i := range initial {
		ln.m[i] = 1
	}
	return ln
}

func (ln lastN) add(i int) bool {
	result := false
	for _, j := range ln.slice {
		if j*2 == i {
			if ln.m[j] > 1 {
				result = true
				break
			}
		} else {
			if ln.m[i-j] > 0 {
				result = true
				break
			}
		}
	}

	old := ln.slice[0]
	copy(ln.slice, ln.slice[1:])
	ln.slice[len(ln.slice)-1] = i
	ln.m[old] = ln.m[old] - 1
	ln.m[i] = ln.m[i] + 1
	return result
}

func firstInvalid(inputs []string, prefix int) (int, error) {
	ints, err := ioutil.StringsToInts(inputs)
	if err != nil {
		return 0, err
	}
	ln := newLastN(ints[:prefix])
	for _, i := range ints[prefix:] {
		if !ln.add(i) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("no invalid inputs")
}

func findRange(ints []int, target int) ([]int, error) {
	for i := 0; i < len(ints); i++ {
		sum := 0
		for j := i; j < len(ints) && sum < target; j++ {
			sum += ints[j]
			if sum == target {
				return ints[i : j+1], nil
			}
		}
	}
	return nil, fmt.Errorf("no range found summing to %d", target)
}

func weakness(inputs []string, prefix int) (int, error) {
	target, err := firstInvalid(inputs, prefix)
	if err != nil {
		return 0, err
	}
	ints, err := ioutil.StringsToInts(inputs)
	if err != nil {
		return 0, err
	}

	r, err := findRange(ints, target)
	if err != nil {
		return 0, err
	}
	min := r[0]
	max := r[0]
	for _, i := range r {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min + max, nil
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
