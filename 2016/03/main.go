package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/math"
	"github.com/zellyn/adventofcode/util"
)

func readGrid(filename string) ([][]int, error) {
	lines, err := ioutil.ReadLines(filename)
	if err != nil {
		return nil, err
	}
	return util.ParseGrid(lines)
}

func possible(x, y, z int) bool {
	x, y, z = math.Sort3(x, y, z)
	return x+y > z
}

func countPossible(grid [][]int) int {
	count := 0
	for _, row := range grid {
		if possible(row[0], row[1], row[2]) {
			count++
		}
	}
	return count
}

func countPossible2(grid [][]int) int {
	count := 0
	for row := 0; row < len(grid); row += 3 {
		for col := 0; col < len(grid[0]); col++ {
			if possible(grid[row][col], grid[row+1][col], grid[row+2][col]) {
				count++
			}
		}
	}
	return count
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
