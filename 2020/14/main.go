package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var memRe = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)

const allOnes = 1<<36 - 1

func parseMask1(s string) (int, int, error) {
	orMask := 0
	andMask := 0
	if len(s) != 36 {
		return 0, 0, fmt.Errorf("weird mask: %q", s)
	}
	for _, c := range s {
		orMask <<= 1
		andMask <<= 1
		switch c {
		case 'X':
			andMask |= 1
		case '0':
		case '1':
			orMask |= 1
			andMask |= 1
		default:
			return 0, 0, fmt.Errorf("weird mask: %q", s)
		}
	}

	return orMask, andMask, nil
}

func part1(inputs []string) (int, error) {
	mem := make(map[int]int)

	orMask := 0
	andMask := allOnes
	var err error
	for _, input := range inputs {
		if strings.HasPrefix(input, "mask = ") {
			orMask, andMask, err = parseMask1(input[7:])
			if err != nil {
				return 0, err
			}
		} else {
			matches := memRe.FindStringSubmatch(input)
			if len(matches) != 3 {
				return 0, fmt.Errorf("weird input: %q", input)
			}
			addr, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, err
			}
			arg, err := strconv.Atoi(matches[2])
			if err != nil {
				return 0, err
			}

			addr &= allOnes
			arg |= orMask
			arg &= andMask
			mem[addr] = arg
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}
	return sum, nil
}

func applyMask(addr int, mask string) []int {
	for i, c := range mask {
		if c == '1' {
			addr |= 1 << (35 - i)
		}
	}

	result := []int{addr}

	for i, c := range mask {
		if c != 'X' {
			continue
		}
		ones := make([]int, len(result))
		copy(ones, result)

		for j := range result {
			m := 1 << (35 - i)
			result[j] &^= m
			ones[j] |= m
		}
		result = append(result, ones...)
	}

	return result
}

func part2(inputs []string) (int, error) {
	mem := make(map[int]int)

	mask := ""
	for _, input := range inputs {
		if strings.HasPrefix(input, "mask = ") {
			mask = input[7:]
		} else {
			matches := memRe.FindStringSubmatch(input)
			if len(matches) != 3 {
				return 0, fmt.Errorf("weird input: %q", input)
			}
			addr, err := strconv.Atoi(matches[1])
			if err != nil {
				return 0, err
			}
			arg, err := strconv.Atoi(matches[2])
			if err != nil {
				return 0, err
			}

			all := applyMask(addr, mask)
			// fmt.Printf("applying mask %q to addr %d yields\n%v\n", mask, addr, all)
			for _, one := range all {
				mem[one] = arg
			}
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
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
