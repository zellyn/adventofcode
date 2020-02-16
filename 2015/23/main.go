package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type opType struct {
	opcode string
	reg    string
	intVal int
}

var lens = map[string]int{
	"hlf": 2,
	"tpl": 2,
	"inc": 2,
	"jmp": 2,
	"jie": 3,
	"jio": 3,
}

func parseReg(s string, trailer string) (string, error) {
	switch s {
	case "a", "b":
		return s, nil
	case "a,", "b,":
		return s[:1], nil
	default:
		return "", fmt.Errorf("weird register (%q) %s", s, trailer)
	}
}

func parseProgram(program []string) ([]opType, error) {
	var ops []opType
	for i, line := range program {
		trailer := fmt.Sprintf("in line %q (%d)", line, i+1)
		parts := strings.Split(line, " ")
		opcode := parts[0]
		if len(parts) != lens[opcode] {
			return nil, fmt.Errorf("%q requires 1 arg; got %d %s", opcode, len(parts)-1, trailer)
		}
		op := opType{
			opcode: opcode,
		}
		if opcode == "hlf" || opcode == "tpl" || opcode == "inc" || opcode == "jie" || opcode == "jio" {
			reg, err := parseReg(parts[1], trailer)
			if err != nil {
				return nil, err
			}
			op.reg = reg
		}

		if opcode == "jmp" || opcode == "jie" || opcode == "jio" {
			i, err := strconv.Atoi(parts[len(parts)-1])
			if err != nil {
				return nil, fmt.Errorf("%v %s", err, trailer)
			}
			op.intVal = i
		}

		ops = append(ops, op)
	}

	return ops, nil
}

func runProgram(program []string, startA int, startB int) (int, int, error) {
	ops, err := parseProgram(program)
	if err != nil {
		return 0, 0, err
	}
	pc := 0
	regs := map[string]int{
		"a": startA,
		"b": startB,
	}
	for {
		next := pc + 1
		if pc < 0 || pc >= len(program) {
			break
		}
		// fmt.Printf("%02d: A=%02d B=%02d - %s\n", pc, regs["a"], regs["b"], program[pc])
		op := ops[pc]
		switch op.opcode {
		case "hlf":
			regs[op.reg] = regs[op.reg] / 2
		case "tpl":
			regs[op.reg] = regs[op.reg] * 3
		case "inc":
			regs[op.reg]++
		case "jmp":
			next = pc + op.intVal
		case "jie":
			if regs[op.reg]%2 == 0 {
				next = pc + op.intVal
			}
		case "jio":
			if regs[op.reg] == 1 {
				next = pc + op.intVal
			}
		}

		pc = next
	}
	return regs["a"], regs["b"], nil
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
