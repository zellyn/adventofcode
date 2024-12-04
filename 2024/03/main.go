package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

var mulRe = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)

func findMuls(s string) [][2]int {
	muls := mulRe.FindAllStringSubmatch(s, -1)
	return util.Map(muls, func(submatches []string) [2]int {
		i1, _ := strconv.Atoi(submatches[1])
		i2, _ := strconv.Atoi(submatches[2])
		return [2]int{i1, i2}
	})
}

var conditionalRe = regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`)

func part1(inputs []string) (int, error) {
	input := strings.Join(inputs, "\n")
	muls := findMuls(input)
	sum := 0
	for _, m := range muls {
		sum += m[0] * m[1]
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	input := strings.Join(inputs, "\n")
	pieces := conditionalRe.FindAllStringSubmatch(input, -1)
	enabled := true
	sum := 0
	for _, piece := range pieces {
		fmt.Println(piece)
		if piece[0] == "do()" {
			enabled = true
			continue
		}

		if piece[0] == "don't()" {
			enabled = false
			continue
		}

		if enabled {
			i1, _ := strconv.Atoi(piece[1])
			i2, _ := strconv.Atoi(piece[2])
			sum += i1 * i2
			// fmt.Printf("Multiplying %d * %d = %d; new sum = %d\n", i1, i2, i1*i2, sum)
		}
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
