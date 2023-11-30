package main

import (
	"fmt"
	"os"
	"strconv"
)

func valid(s string) bool {
	last := ' '
	double := false
	for _, c := range s {
		double = double || (c == last)
		if c < last {
			return false
		}
		last = c
	}
	return double
}

func valid2(s string) bool {
	last := ' '
	repCount := 1
	double := false
	for _, c := range s {
		if c == last {
			repCount++
		} else {
			double = double || (repCount == 2)
			repCount = 1
		}
		if c < last {
			return false
		}
		last = c
	}
	double = double || (repCount == 2)
	return double
}

func countBetween(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if valid(strconv.Itoa(i)) {
			count++
		}
	}
	return count
}

func countBetween2(start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if valid2(strconv.Itoa(i)) {
			count++
		}
	}
	return count
}

func run() error {
	fmt.Printf("Count 231832-767346: %d\n", countBetween(231832, 767346))
	fmt.Printf("Count 231832-767346: %d\n", countBetween2(231832, 767346))
	fmt.Printf("Valid 112233? %v\n", valid2("112233"))
	fmt.Printf("Valid 123444? %v\n", valid2("123444"))
	fmt.Printf("Valid 111122? %v\n", valid2("111122"))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
