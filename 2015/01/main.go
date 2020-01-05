package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/ioutil"
)

func run() error {
	s, err := ioutil.ReadFileString("input")
	if err != nil {
		return err
	}

	floor := strings.Count(s, "(") - strings.Count(s, ")")
	fmt.Println("Floor:", floor)

	floor = 0

LOOP:
	for i, r := range s {
		switch r {
		case '(':
			floor++
		case ')':
			floor--
			if floor == -1 {
				fmt.Println("Basement at:", i+1)
				break LOOP
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
