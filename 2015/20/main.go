package main

import (
	"fmt"
	"os"

	"github.com/zellyn/project-euler/go/primes"
)

func score(i int) int {
	sum := i * 10
	for _, dv := range primes.ProperDivisors(i) {
		sum += 10 * dv
	}
	return sum
}

func score50(i int) int {
	sum := i * 11
	for _, dv := range primes.ProperDivisors(i) {
		count := i / dv
		if count > 50 {
			continue
		}
		sum += 11 * dv
	}
	return sum
}

func minHouse(minScore int) int {
	for i := 1; i <= minScore/10; i++ {
		if score(i) >= minScore {
			return i
		}
	}
	return -1
}

func minHouse50(minScore int) int {
	for i := 1; i <= minScore/11; i++ {
		if score50(i) >= minScore {
			return i
		}
	}
	return -1
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
