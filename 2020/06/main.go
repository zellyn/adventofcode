package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/stringset"
	"github.com/zellyn/adventofcode/util"
)

func countAny(inputs []string) int {
	all := stringset.New()
	for _, input := range inputs {
		s := stringset.OfRunes(input)
		all.AddAll(s)
	}
	return len(all)
}

func countAll(inputs []string) int {
	var all stringset.S
	for i, input := range inputs {
		s := stringset.OfRunes(input)
		if i == 0 {
			all = s
		} else {
			all = stringset.Intersect(all, s)
		}
	}
	return len(all)
}

func sumAny(input []string) int {
	paras := util.LinesByParagraph(input)
	total := 0
	for _, para := range paras {
		total += countAny(para)
	}
	return total
}

func sumAll(input []string) int {
	paras := util.LinesByParagraph(input)
	total := 0
	for _, para := range paras {
		total += countAll(para)
	}
	return total
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
