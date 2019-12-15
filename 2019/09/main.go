package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/2019/intcode"
)

func run() error {
	state, err := intcode.ReadProgram("input")
	if err != nil {
		return err
	}

	_, writes, err := intcode.RunProgram(state, []int64{1}, false)
	if err != nil {
		return err
	}

	fmt.Printf("Output: %v\n", writes)

	_, writes, err = intcode.RunProgram(state, []int64{2}, false)
	if err != nil {
		return err
	}

	fmt.Printf("Output: %v\n", writes)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
