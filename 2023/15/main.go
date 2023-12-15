package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type lens struct {
	name string
	f    int
}

func hash(s string) int {
	var h byte

	for _, b := range []byte(s) {
		h = (h + b) * 17
	}

	return int(h)
}

func doPart(cmd string, boxes *[256][]lens) error {
	if strings.Contains(cmd, "-") {
		name := cmd[:len(cmd)-1]
		boxNum := hash(name)
		for i, lens := range boxes[boxNum] {
			if lens.name == name {
				boxes[boxNum] = slices.Delete(boxes[boxNum], i, i+1)
				return nil
			}
		}
	} else {
		parts := strings.Split(cmd, "=")
		name := parts[0]
		f, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}

		boxNum := hash(name)
		for i, lens := range boxes[boxNum] {
			if lens.name == name {
				boxes[boxNum][i].f = f
				return nil
			}
		}
		boxes[boxNum] = append(boxes[boxNum], lens{name: name, f: f})
	}
	return nil
}

func score(boxes [256][]lens) int {
	sum := 0
	for i, lenses := range boxes {
		for j, lens := range lenses {
			sum += (1 + i) * (1 + j) * lens.f
		}
	}

	return sum
}

func part1(inputs []string) (int, error) {
	parts := strings.Split(inputs[0], ",")
	return util.MappedSum(parts, hash), nil
}

func part2(inputs []string) (int, error) {
	parts := strings.Split(inputs[0], ",")
	var boxes [256][]lens
	for _, part := range parts {
		doPart(part, &boxes)
	}
	return score(boxes), nil
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
