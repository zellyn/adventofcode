package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/ioutil"
)

func which(input string, which int) (int, error) {
	ints, err := ioutil.ParseInts(input, ",")
	if err != nil {
		return 0, err
	}
	when := make(map[int]int)
	for i := 0; i < len(ints)-1; i++ {
		when[ints[i]] = i
	}
	last := ints[len(ints)-1]
	for i := len(ints) - 1; i < which-1; i++ {
		previous, found := when[last]
		when[last] = i
		if !found {
			last = 0
		} else {
			last = i - previous
		}
	}
	return last, nil
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
