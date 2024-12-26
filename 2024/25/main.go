package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type profile []int

func parse(inputs []string) ([]profile, []profile, error) {
	paras := util.LinesByParagraph(inputs)
	var keys []profile
	var locks []profile

	for _, para := range paras {
		m := charmap.Parse(para)

		if m[geom.Z2] == '#' {
			// lock
			lock := make(profile, 5)
			for _, pos := range m.FindAll('#') {
				lock[pos.X] = max(lock[pos.X], pos.Y)
			}
			locks = append(locks, lock)
		} else {
			// key
			key := make(profile, 5)
			for _, pos := range m.FindAll('#') {
				key[pos.X] = max(key[pos.X], 6-pos.Y)
			}
			keys = append(keys, key)
		}

	}
	return locks, keys, nil
}

func fit(lock, key profile) bool {
	for i := range lock {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func part1(inputs []string) (int, error) {
	locks, keys, err := parse(inputs)
	if err != nil {
		return 0, nil
	}

	fits := 0
	for _, key := range keys {
		for _, lock := range locks {
			if fit(lock, key) {
				fits++
			}
		}
	}
	return fits, nil
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
