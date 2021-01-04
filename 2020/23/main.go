package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

func move(next []int, curr int) int {
	// fmt.Printf("move(%v, %d)\n", next, curr)
	a := next[curr-1]
	b := next[a-1]
	c := next[b-1]
	d := next[c-1]
	target := curr
	for {
		target = target - 1
		if target == 0 {
			target = len(next)
		}
		if target != a && target != b && target != c {
			break
		}
	}
	next[c-1] = next[target-1]
	next[target-1] = a
	next[curr-1] = d
	return d
}

func toNext(ints []int, ll int) ([]int, int) {
	l := len(ints)
	if ll == -1 {
		ll = l
	}
	first := ints[0]
	last := ints[l-1]
	result := make([]int, ll)
	for i := 0; i < l-1; i++ {
		result[ints[i]-1] = ints[i+1]
	}
	result[last-1] = first
	if l == ll {
		return result, first
	}

	for i := l + 1; i < ll; i++ {
		result[i-1] = i + 1
	}
	result[ll-1] = first
	result[last-1] = l + 1
	return result, first
}

func score(next []int) string {
	l := len(next)
	pos := 1
	var result bytes.Buffer
	for j := 1; j < l; j++ {
		pos = next[pos-1]
		fmt.Fprintf(&result, "%d", pos)
	}
	return result.String()
}

func part1(input string, moves int) (string, error) {
	intStrs := strings.Split(input, "")
	ints, err := util.StringsToInts(intStrs)
	if err != nil {
		return "", err
	}
	ints, pos := toNext(ints, -1)
	for i := 0; i < moves; i++ {
		pos = move(ints, pos)
	}
	return score(ints), nil
}

func part2(input string, moves int) (int, error) {
	intStrs := strings.Split(input, "")
	ints, err := util.StringsToInts(intStrs)
	if err != nil {
		return 0, err
	}
	ints, pos := toNext(ints, 1000000)
	for i := 0; i < moves; i++ {
		pos = move(ints, pos)
	}
	return ints[0] * ints[ints[0]-1], nil
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
