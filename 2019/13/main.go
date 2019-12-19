package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/zellyn/adventofcode/2019/intcode"
	"github.com/zellyn/adventofcode/geom"
)

type vec2 = geom.Vec2

const (
	itemEmpty = iota
	itemWall
	itemBlock
	itemHorizontalPaddle
	itemBall
)

var chars = map[int]rune{
	0: ' ',
	1: '#',
	2: 'B',
	3: '=',
	4: 'o',
}

func runProgram(program []int64, startColor bool) (map[vec2]int, error) {
	errChan := make(chan error)
	readChan := make(chan int64)
	writeChan := make(chan int64)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	drawing := map[vec2]int{}

	go intcode.RunProgramChans(ctx, program, readChan, writeChan, errChan, nil, false, "")

OUTER:
	for {
		select {
		case err := <-errChan:
			if err != nil {
				return nil, err
			}
			break OUTER
		case <-ctx.Done():
			return nil, fmt.Errorf("timeout")
		case x := <-writeChan:
			y := <-writeChan
			id := <-writeChan
			drawing[vec2{int(x), int(y)}] = int(id)
		}
	}
	return drawing, nil
}

func playProgram(program []int64) (map[vec2]int, int, error) {
	program = intcode.Copy(program)
	program[0] = 2

	errChan := make(chan error)
	readChan := make(chan int64)
	writeChan := make(chan int64)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	drawing := map[vec2]int{}
	score := 0

	go intcode.RunProgramChans(ctx, program, readChan, writeChan, errChan, nil, false, "")

	var ballPos vec2
	var output int64
OUTER:
	for {
		// fmt.Printf("ball: %d, ballvec: %d, paddle: %d, output: %d\n", ballPos.X, ballVec.X, paddlePos.X, output)
		select {
		case err := <-errChan:
			if err != nil {
				return nil, 0, err
			}
			break OUTER
		case <-ctx.Done():
			return nil, 0, fmt.Errorf("timeout")
		case readChan <- output:
			// sweet!

		case x := <-writeChan:
			// success!
			y := <-writeChan
			z := <-writeChan
			if x == -1 && y == 0 {
				score = int(z)
				continue
			}
			drawing[vec2{int(x), int(y)}] = int(z)
			if z != itemBall {
				continue
			}

			ballPos = find(drawing, itemBall)
			paddlePos := find(drawing, itemHorizontalPaddle)
			if paddlePos.X == -1 {
				continue
			}
			// draw(drawing)
			output = int64(ballPos.Sub(paddlePos).Sgn().X)

			// time.Sleep(300 * time.Millisecond)
		}

	}
	return drawing, score, nil
}

func find(drawing map[vec2]int, item int) vec2 {
	for k, v := range drawing {
		if v == item {
			return k
		}
	}
	return vec2{-1, -1}
}

func draw(drawing map[vec2]int) {
	minx, miny, maxx, maxy := 0, 0, 0, 0
	for k := range drawing {
		if k.X < minx {
			minx = k.X
		}
		if k.X > maxx {
			maxx = k.X
		}
		if k.Y < miny {
			miny = k.Y
		}
		if k.Y > maxy {
			maxy = k.Y
		}
	}
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			id := drawing[vec2{x, y}]
			c, ok := chars[id]
			if !ok {
				c = '?'
			}
			fmt.Printf("%c", c)
		}
		fmt.Println()
	}
}

func counts(drawing map[vec2]int) map[int]int {
	c := make(map[int]int)
	for _, v := range drawing {
		c[v]++
	}
	return c
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

	draw(drawing)
	c := counts(drawing)
	fmt.Println(c[itemBlock])

	drawing, score, err := playProgram(program)
	if err != nil {
		return err
	}
	fmt.Println(score)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
