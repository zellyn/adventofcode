package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type term struct {
	name  string
	val   *float64
	term1 string
	term2 string
	op    string
}

func parseTerm(input string) term {
	parts := strings.Split(input, ": ")
	t := term{
		name: parts[0],
	}

	terms := strings.Split(parts[1], " ")
	if len(terms) == 1 {
		i, err := strconv.Atoi(terms[0])
		if err != nil {
			panic(err)
		}
		f := float64(i)
		t.val = &f
	} else {
		t.term1, t.op, t.term2 = terms[0], terms[1], terms[2]
	}

	return t
}

func parseInput(inputs []string) map[string]term {
	result := make(map[string]term, len(inputs))
	for _, input := range inputs {
		t := parseTerm(input)
		result[t.name] = t
	}
	return result
}

var div0 = errors.New("divide by zero")

func eval(terms map[string]term, which string) (float64, error) {
	t, ok := terms[which]
	if !ok {
		panic(fmt.Sprintf("can't find term %q", which))
	}
	if t.val != nil {
		return *t.val, nil
	}
	val1, err := eval(terms, t.term1)
	if err != nil {
		return 0, err
	}
	val2, err := eval(terms, t.term2)
	if err != nil {
		return 0, err
	}
	var result float64
	switch t.op {
	case "+":
		result = val1 + val2
	case "-":
		result = val1 - val2
	case "*":
		result = val1 * val2
	case "/":
		if val2 == 0 {
			return 0, div0
		}
		result = val1 / val2
	default:
		panic(fmt.Sprintf("weird op %q", t.op))
	}
	return result, nil
}

func part1(inputs []string) (int, error) {
	terms := parseInput(inputs)

	result, err := eval(terms, "root")
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

func sgn(f float64) int {
	if f > 0 {
		return 1
	}
	if f < 0 {
		return -1
	}
	return 0
}

func abs(f float64) float64 {
	if f < 0 {
		return -f
	}
	return f
}

func part2(inputs []string) (int, error) {
	terms := parseInput(inputs)
	delete(terms, "humn")
	root := terms["root"]
	target, err := eval(terms, root.term2)
	if err != nil {
		return 0, err
	}
	fmt.Printf("target: %.6g\n", target)

	name := root.term1

	e := func(f float64) (float64, error) {
		terms["humn"] = term{val: &f}
		return eval(terms, name)
	}

	var lower, vlower float64
	for ; ; lower++ {
		vlower, err = e(lower)
		if err == nil {
			break
		}
	}
	upper := lower + 1
	var vupper float64

	for ; ; upper++ {
		vupper, err = e(upper)
		if err == nil && vupper != vlower {
			break
		}
	}

	if sgn(vlower-target) != sgn(vupper-target) {
		return 0, fmt.Errorf("cannot bisect [%.6g,%.6g]", vlower, vupper)
	}

	fmt.Printf("e(%.6g)=%.6g\n", lower, vlower)
	fmt.Printf("e(%.6g)=%.6g\n", upper, vupper)
	sign := sgn(vlower - target)

	mag1 := abs(vlower - target)
	mag2 := abs(vupper - target)
	searchSign := sgn(mag1 - mag2)

	var search, vsearch float64
	for search = 2 * float64(searchSign); ; search *= 2 {
		for ; ; search += 1 {
			vsearch, err = e(search)
			fmt.Printf("e(%.6g)=%.6g\n", search, vsearch)
			if err == nil {
				break
			}
		}
		if sgn(vsearch-target) != sign {
			break
		}
	}

	if sgn(vsearch-target) == 0 {
		return int(vsearch), nil
	}

	if searchSign > 0 {
		upper = search
		vupper = vsearch
	} else {
		lower = search
		vlower = vsearch
	}

	for lower != upper {
		fmt.Printf("bisecting for target=%.6g from e(%.6g)=%.6g to e(%.6g)=%.6g\n", target, lower, vlower, upper, vupper)
		mid := upper/2 + lower/2
		var vmid float64

		for ; ; mid++ {
			if vmid, err = e(mid); err != nil {
				mid++
				continue
			}
			break
		}

		switch sgn(vmid - target) {
		case 0:
			fmt.Printf("match on %.6g\n", mid)
			return int(mid), nil
		case sgn(vupper - target):
			upper = mid
			vupper = vmid
		case sgn(vlower - target):
			lower = mid
			vlower = vmid
		}
	}

	return 42, nil
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
