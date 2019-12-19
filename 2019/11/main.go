package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/zellyn/adventofcode/2019/intcode"
)

const (
	N = 0
	E = 1
	S = 2
	W = 3
)

var vectors = map[int][2]int{
	N: {0, -1},
	E: {1, 0},
	S: {0, 1},
	W: {-1, 0},
}

func runProgram(program []int64, startColor bool) (map[[2]int]bool, error) {
	errChan := make(chan error)
	readChan := make(chan int64)
	writeChan := make(chan int64)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	drawing := map[[2]int]bool{
		{0, 0}: startColor,
	}
	pos := [2]int{0, 0}
	heading := N

	go intcode.RunProgramChans(ctx, program, readChan, writeChan, errChan, nil, false, "")

OUTER:
	for {
		var read int64
		if drawing[pos] {
			read = 1
		}
		select {
		case err := <-errChan:
			if err != nil {
				return nil, err
			}
			break OUTER
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout")
		case readChan <- read:
			// success!
		}
		color := <-writeChan
		drawing[pos] = color == 1
		dir := int(<-writeChan)
		heading = (heading + (dir*2 - 1) + 4) % 4
		vector := vectors[heading]
		pos[0] += vector[0]
		pos[1] += vector[1]
	}
	return drawing, nil
}

func draw(drawing map[[2]int]bool) {
	minx, miny, maxx, maxy := 0, 0, 0, 0
	for k := range drawing {
		if k[0] < minx {
			minx = k[0]
		}
		if k[0] > maxx {
			maxx = k[0]
		}
		if k[1] < miny {
			miny = k[1]
		}
		if k[1] > maxy {
			maxy = k[1]
		}
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			c := " "
			if drawing[[2]int{x, y}] {
				c = "#"
			}
			fmt.Printf("%s", c)
		}
		fmt.Println()
	}
}

func run() error {
	program, err := intcode.ReadProgram("input")
	if err != nil {
		return err
	}

	drawing, err := runProgram(program, false)
	if err != nil {
		return err
	}

	fmt.Printf("painted squares: %d\n", len(drawing))

	drawing, err = runProgram(program, true)
	if err != nil {
		return err
	}
	fmt.Printf("painted squares: %d\n", len(drawing))
	draw(drawing)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
