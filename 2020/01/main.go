package main

import (
	"fmt"
	"os"
)

func twoAddTo(sum int, ints []int) (int, error) {
	m := make(map[int]int)
	for _, i := range ints {
		// Check for squares
		if i == sum/2 && m[i] == i {
			return i * i, nil
		}
		m[i] = i
	}
	for _, i := range ints {
		if m[sum-i] > 0 {
			return i * (sum - i), nil
		}
	}
	return -1, fmt.Errorf("can't find numbers that sum to %d", sum)
}

func threeAddTo(sum int, ints []int) (int, error) {
	l := len(ints)
	for i, a := range ints[:len(ints)-2] {
		if a > sum {
			continue
		}
		for j := i + 1; j < l-1; j++ {
			b := ints[j]
			if a+b > sum {
				continue
			}

			for k := j + 1; k < l; k++ {
				c := ints[k]
				if a+b+c == sum {
					return a * b * c, nil
				}
			}
		}
	}
	return -1, fmt.Errorf("can't find numbers that sum to %d", sum)
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
