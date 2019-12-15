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

func run() error {
	bb, err := ioutil.ReadFile("input")
	if err != nil {
		return err
	}
	contents := string(bb)
	lines := strings.Split(contents, "\n")
	total1 := 0
	total2 := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		total1 += moduleFuel(num)
		total2 += moduleFuelPlusFuel(num)
	}
	fmt.Printf("Total fuel for modules: %d\n", total1)
	fmt.Printf("Total fuel for modules+fuel: %d\n", total2)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
