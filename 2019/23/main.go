package main

import (
	"context"
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/zellyn/adventofcode/2019/intcode"
	"github.com/zellyn/adventofcode/geom"
)

type vec2 = geom.Vec2
type vec3 = geom.Vec3

const empty = int64(-1)

func runOne(ctx context.Context, i int, program []int64, inputChan <-chan vec2, outputChan chan<- vec3, sharedErrChan chan<- error, readWaitCount *int32) {
	readChan := make(chan int64)
	writeChan := make(chan int64)
	errChan := make(chan error)
	updateChan := make(chan [2]int64)

	logErr := func(err error) {
		sharedErrChan <- err
	}
	go intcode.RunProgramChans(ctx, program, readChan, writeChan, errChan, updateChan, false, "")
	readChan <- int64(i)

	var queue []vec2
	output := empty
	atomic.AddInt32(readWaitCount, 1)
	for {
		select {
		case <-ctx.Done():
			logErr(nil)
			return
		case err := <-errChan:
			logErr(err)
			return
		case readChan <- output:
			if output != empty {
				readChan <- int64(queue[0].Y)
				queue = queue[1:]
				if len(queue) == 0 {
					atomic.AddInt32(readWaitCount, 1)
					output = empty
				} else {
					output = int64(queue[0].X)
				}
			}
		case input := <-inputChan:
			if output == empty {
				atomic.AddInt32(readWaitCount, -1)
				output = int64(input.X)
			}
			queue = append(queue, input)
		case dest := <-writeChan:
			x := <-writeChan
			y := <-writeChan
			outputChan <- vec3{X: int(x), Y: int(y), Z: int(dest)}
		}
	}
}

func runAll(ctx context.Context, program []int64, count int, returnFirst bool) (int, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	inputChan := make([]chan vec2, count)
	outputChan := make(chan vec3, count)
	errChan := make(chan error, count)

	for i := 0; i < count; i++ {
		inputChan[i] = make(chan vec2, count)
		defer close(inputChan[i])
	}

	var readWaitCount int32

	for i := 0; i < count; i++ {
		go runOne(ctx, i, program, inputChan[i], outputChan, errChan, &readWaitCount)
	}

	last255 := vec3{-1, -1, -1}
	sent := map[int]bool{}
	i := 0

	stuck := false

	for {
		select {
		case <-ctx.Done():
			return 0, nil
		case <-time.After(100 * time.Millisecond):
			waiting := int(atomic.LoadInt32(&readWaitCount))
			if waiting == count {
				if last255.Z == -1 {
					return 0, fmt.Errorf("deadlocked, but never seen a 255-packet")
				}
				if sent[last255.Y] {
					return last255.Y, nil
				}
				sent[last255.Y] = true
				fmt.Printf("Sending X=%d,Y=%d to computer 0\n", last255.X, last255.Y)
				inputChan[0] <- vec2{X: last255.X, Y: last255.Y}
			} else {
				if stuck {
					fmt.Println("Warning: stuck twice in a row... ?")
				}
				stuck = true
			}
		case err := <-errChan:
			return 0, err
		case packet := <-outputChan:
			stuck = false
			if packet.Z == 255 {
				if returnFirst {
					return int(packet.Y), nil
				}
				i++
				// fmt.Printf("NAT Packet %d: X=%d,Y=%d\n", i, packet.X, packet.Y)
				last255 = packet
			} else {
				if packet.Z < 0 || packet.Z >= count {
					return 0, fmt.Errorf("Weird packet address: %d (X=%d,Y=%d)", packet.Z, packet.X, packet.Y)
				}
				inputChan[int(packet.Z)] <- vec2{X: packet.X, Y: packet.Y}
			}
		}
	}
}

func run() error {
	program, err := intcode.ReadProgram("input")
	if err != nil {
		return err
	}

	result, err := runAll(context.Background(), program, 50, false)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
