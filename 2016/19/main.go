package main

import (
	"fmt"
	"os"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func part1(count int) (int, error) {
	first := 1
	pow := 1
	for count > 1 {
		pow *= 2
		if count%2 == 1 {
			first += pow
		}
		count /= 2
	}
	return first, nil
}

func part2(count int) (int, error) {
	nextElf := make([]int, count+1)
	for i := 1; i < count; i++ {
		nextElf[i] = i + 1
	}
	nextElf[count] = 1

	munch := count / 2

	for count > 1 {
		// munch opposite
		nextElf[munch] = nextElf[nextElf[munch]]
		count--
		if count%2 == 0 {
			munch = nextElf[munch]
		}
	}

	return munch, nil
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
