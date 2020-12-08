package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type op struct {
	name string
	arg  int
}

func (o op) String() string {
	return fmt.Sprintf("%s %+d", o.name, o.arg)
}

func parseOp(input string) (op, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return op{}, fmt.Errorf("weird input: %q", input)
	}
	i, err := strconv.Atoi(parts[1])
	if err != nil {
		return op{}, fmt.Errorf("weird input: %q: %w", input, err)
	}
	return op{
		name: parts[0],
		arg:  i,
	}, nil
}

func parseProg(inputs []string) (map[int]op, error) {
	result := make(map[int]op)
	for i, input := range inputs {
		op, err := parseOp(input)
		if err != nil {
			return nil, err
		}
		result[i] = op
	}

	return result, nil
}

func printProg(prog map[int]op) {
	end := len(prog)
	for i := 0; i < end; i++ {
		fmt.Printf("%03d: %s\n", i, prog[i])
	}
}

func step(pc int, acc int, prog map[int]op) (int, int, error) {
	o, found := prog[pc]
	if !found {
		return 0, 0, fmt.Errorf("no op at address %d", pc)
	}
	switch o.name {
	case "nop":
		pc++
	case "acc":
		acc += o.arg
		pc++
	case "jmp":
		pc += o.arg
	default:
		return 0, 0, fmt.Errorf("unknown op %q at address %d", o.name, pc)
	}

	return pc, acc, nil
}

func accBeforeLoop(inputs []string) (int, error) {
	prog, err := parseProg(inputs)
	if err != nil {
		return 0, err
	}
	seen := make(map[int]bool)
	pc := 0
	acc := 0
	for {
		if seen[pc] {
			return acc, nil
		}
		seen[pc] = true
		pc, acc, err = step(pc, acc, prog)
		if err != nil {
			return 0, err
		}
	}
}

func flip(prog map[int]op, index int) (map[int]op, error) {
	cp := make(map[int]op, len(prog))
	for k, v := range prog {
		cp[k] = v
	}

	o, found := cp[index]
	if !found {
		return nil, fmt.Errorf("no op at address %d", index)
	}

	switch o.name {
	case "jmp":
		cp[index] = op{"nop", o.arg}
	case "nop":
		cp[index] = op{"jmp", o.arg}
	default:
		return nil, nil
	}
	return cp, nil
}

func accFixed(inputs []string) (int, error) {
	prog, err := parseProg(inputs)
	if err != nil {
		return 0, err
	}

	for i := range prog {
		i += 0
		fixed, err := flip(prog, i)
		if err != nil {
			return 0, err
		}
		if fixed == nil {
			continue
		}
		win, acc, err := terminatesNormally(fixed)
		if err != nil {
			return 0, err
		}
		if win {
			return acc, nil
		}
	}

	return 0, fmt.Errorf("No winning program :-(")
}

func terminatesNormally(prog map[int]op) (bool, int, error) {
	end := len(prog)
	var err error
	seen := make(map[int]bool)
	pc := 0
	acc := 0
	for {
		if pc == end {
			return true, acc, nil
		}
		if seen[pc] {
			return false, acc, nil
		}
		seen[pc] = true
		pc, acc, err = step(pc, acc, prog)
		if err != nil {
			return false, 0, err
		}
	}
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
