package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func moduleFuel(size int) int {
	fuel := size/3 - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func moduleFuelPlusFuel(size int) int {
	fuel := moduleFuel(size)
	total := fuel
	for fuel > 0 {
		fuel = moduleFuel(fuel)
		total += fuel
	}
	return total
}

func commaStringToInts(input string) ([]int, error) {
	input = strings.TrimSpace(input)
	entries := strings.Split(input, ",")
	ints := make([]int, len(entries))
	for i, v := range entries {
		intVal, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		ints[i] = intVal
	}
	return ints, nil
}

func readCommaDelimited(filename string) ([]int, error) {
	bb, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}
	return commaStringToInts(string(bb))
}

func runFromString(input string) error {
	state, err := commaStringToInts(input)
	if err != nil {
		return err
	}
	return runProgram(state)
}

func runProgram(state []int) error {
	pc := 0
	for {
		cmd := state[pc]
		switch cmd {
		case 99:
			return nil
		case 1:
			state[state[pc+3]] = state[state[pc+1]] + state[state[pc+2]]
			pc += 4
		case 2:
			state[state[pc+3]] = state[state[pc+1]] * state[state[pc+2]]
			pc += 4
		default:
			return fmt.Errorf("unknown opcode %d found at position %d", cmd, pc)
		}

	}
}

func runWith(originalState []int, noun int, verb int) (int, error) {
	state := make([]int, len(originalState))
	copy(state, originalState)
	state[1] = noun
	state[2] = verb
	if err := runProgram(state); err != nil {
		return 0, err
	}
	return state[0], nil
}

func run() error {
	state, err := readCommaDelimited("input")
	if err != nil {
		return err
	}

	result, err := runWith(state, 12, 2)
	if err != nil {
		return err
	}

	fmt.Printf("State at position 0: %d\n", result)

	for noun := 0; noun < len(state); noun++ {
		for verb := 0; verb < len(state); verb++ {
			result, err := runWith(state, noun, verb)
			if err != nil {
				return err
			}
			if result == 19690720 {
				fmt.Printf("noun=%d, verb=%d, answer=%d\n", noun, verb, 100*noun+verb)
				return nil
			}
		}
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
