package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type arg struct {
	name string
	num  uint16
}

type def struct {
	op   string
	arg1 arg
	arg2 arg
}

func parseArg(s string) arg {
	i16, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return arg{name: s}
	}
	return arg{num: uint16(i16)}
}

func parseInput(input string) (map[string]def, error) {
	result := map[string]def{}
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i, line := range lines {
		parts := strings.Split(line, " ")
		l := len(parts)
		if parts[l-2] != "->" {
			return nil, fmt.Errorf("weird input at line %d: %q", i, line)
		}
		name := parts[l-1]
		switch len(parts) {
		case 3:
			a := parseArg(parts[0])
			result[name] = def{
				op:   "EQU",
				arg1: a,
			}
		case 4:
			if parts[0] != "NOT" {
				return nil, fmt.Errorf("weird input at line %d: %q", i, line)
			}
			result[name] = def{
				op:   "NOT",
				arg1: parseArg(parts[1]),
			}
		case 5:
			result[name] = def{
				op:   parts[1],
				arg1: parseArg(parts[0]),
				arg2: parseArg(parts[2]),
			}
		default:
			return nil, fmt.Errorf("weird input at line %d: %q", i, line)
		}
	}

	return result, nil
}

func argVal(a arg, defs map[string]def, memo map[string]uint16) (uint16, error) {
	if a.name == "" {
		return a.num, nil
	}
	return eval(a.name, defs, memo)
}

func eval(name string, defs map[string]def, memo map[string]uint16) (uint16, error) {
	if val, ok := memo[name]; ok {
		return val, nil
	}

	d, ok := defs[name]
	if !ok {
		return 0, fmt.Errorf("no rule for %q", name)
	}

	val1, err := argVal(d.arg1, defs, memo)
	if err != nil {
		return 0, err
	}
	val2, err := argVal(d.arg2, defs, memo)
	if err != nil {
		return 0, err
	}

	_ = val1
	_ = val2
	var val uint16
	switch d.op {
	case "EQU":
		val = val1
	case "OR":
		val = val1 | val2
	case "LSHIFT":
		val = val1 << val2
	case "RSHIFT":
		val = val1 >> val2
	case "NOT":
		val = ^val1
	case "AND":
		val = val1 & val2
	default:
		return 0, fmt.Errorf("unknown op when evaluation %q: %q", name, d.op)
	}
	memo[name] = val
	return val, nil
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
