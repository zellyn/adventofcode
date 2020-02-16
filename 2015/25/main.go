package main

import (
	"fmt"
	"os"
)

func sumTo(n int) int {
	return n * (n + 1) / 2
}

func index(x, y int) int {
	return sumTo(x+y-1) - y + 1
}

func value(x, y int) int {
	ind := index(x, y)
	result := 20151125
	for i := 1; i < ind; i++ {
		result = (result * 252533) % 33554393
	}
	return result
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
