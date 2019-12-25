package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/math"
)

func shuffle(size int, lines []string) (int, int, error) {
	product := 1
	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "cut" {
			cut, err := strconv.Atoi(parts[1])
			if err != nil {
				return 0, 0, err
			}
			sum -= cut

		} else if parts[1] == "into" {
			product = -product
			sum = size - 1 - sum

		} else if parts[1] == "with" {
			inc, err := strconv.Atoi(parts[3])
			if err != nil {
				return 0, 0, err
			}
			product = math.ModMul(product, inc, size) // product *= inc
			sum = math.ModMul(sum, inc, size)         // sum *= inc

		} else {
			return 0, 0, fmt.Errorf("Weird line: %q", line)
		}

		sum %= size
		if sum < 0 {
			sum += size
		}
		product %= size
	}
	if product < 0 {
		product += size
	}
	return product, sum, nil
}

func runForward(prod, sum, times, pos, size int) (int, error) {
	xterm := math.ModMul(pos, math.ModExp(prod, times, size), size)
	geom, err := math.ModGeometricSum(prod, times, size)
	if err != nil {
		return 0, err
	}
	sterm := math.ModMul(sum, geom, size)
	return (xterm + sterm) % size, nil
}

func runBackward(prod, sum, times, pos, size int) (int, error) {
	q, err := math.ModInv(prod, size)
	if err != nil {
		return 0, err
	}
	// fmt.Printf("q=%d, times=%d, size=%d\n", q, times, size)
	// fmt.Printf("math.ModExp(q, times, size)=%d\n", math.ModExp(q, times, size))
	xterm := math.ModMul(pos, math.ModExp(q, times, size), size)
	// fmt.Printf("xterm=%d\n", xterm)
	geom, err := math.ModGeometricSum(q, times+1, size)
	if err != nil {
		return 0, err
	}
	sterm := math.ModMul(sum, (geom - 1), size)
	return (xterm - sterm + size) % size, nil
}

func main() {
	fmt.Fprintf(os.Stderr, "Everything in tests\n")
	os.Exit(1)
}
