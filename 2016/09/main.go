package main

import (
	"fmt"
	"os"
	"strconv"
)

func foo(s string) (int, error) {
	return strconv.Atoi(s)
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
