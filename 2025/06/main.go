package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

type calc struct {
	ints []int
	op   rune
}

func parse1(inputs []string) ([]calc, error) {
	ll := len(strings.Fields(inputs[0]))
	res := make([]calc, ll)

	for _, input := range inputs[:len(inputs)-1] {
		ints, err := util.MapE(strings.Fields(input), strconv.Atoi)
		if err != nil {
			return nil, err
		}

		for i, anInt := range ints {
			res[i].ints = append(res[i].ints, anInt)
		}

		ops := strings.Fields(inputs[len(inputs)-1])
		for i, op := range ops {
			res[i].op = []rune(op)[0]
		}
	}
	return res, nil
}

func parse2(inputs []string) ([]calc, error) {
	ll := len(inputs)
	last := inputs[ll-1]
	var res []calc

	for col := 0; col < len(last); {
		op := []rune(last)[col]
		if op != '+' && op != '*' {
			return nil, fmt.Errorf("unexpected char '%c' at col %d", op, col)
		}
		next := col + 1
		for next < len(last) && last[next] == ' ' {
			next++
		}
		c := calc{op: op}

		for x := col; x < next; x++ {
			num := 0
			allSpace := true
			for y := 0; y < ll-1; y++ {
				ch := inputs[y][x]
				if ch == ' ' {
					continue
				}
				allSpace = false
				num = num*10 + int(ch-'0')
			}
			if !allSpace {
				c.ints = append(c.ints, num)
			}
		}

		col = next
		res = append(res, c)
	}

	return res, nil
}

func (c calc) compute() int {
	n := c.ints[0]
	for _, i := range c.ints[1:] {
		switch c.op {
		case '+':
			n += i
		case '*':
			n *= i
		default:
			panic(fmt.Sprintf("weird op: %c", c.op))
		}
	}

	return n
}

func part1(inputs []string) (int, error) {
	calcs, err := parse1(inputs)
	if err != nil {
		return 0, nil
	}

	sum := 0
	for _, c := range calcs {
		sum += c.compute()
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	calcs, err := parse2(inputs)
	if err != nil {
		return 0, nil
	}

	sum := 0
	for _, c := range calcs {
		sum += c.compute()
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
