package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/zellyn/adventofcode/2019/intcode"
)

func runSequence(program []int64, phases []int64, debug bool) (signal int64, err error) {
	signal = 0

	for _, phase := range phases {
		_, writes, err := intcode.RunProgram(program, []int64{phase, signal}, debug)
		if err != nil {
			return 0, err
		}
		if len(writes) != 1 {
			return 0, fmt.Errorf("want 1 write at phase %d; got %d (%v)", phase, len(writes), writes)
		}
		signal = writes[0]
	}
	return signal, nil
}

func bestSequence(program []int64, debug bool) (signal int64, sequence []int64, err error) {
	var bestSeq []int64
	bestSignal := int64(-1)

	for _, sequence := range permutations([]int64{0, 1, 2, 3, 4}) {
		signal, err := runSequence(program, sequence, debug)
		if err != nil {
			return 0, nil, err
		}
		if signal > bestSignal {
			bestSignal = signal
			bestSeq = sequence
		}
	}
	return bestSignal, bestSeq, nil
}

func runParallelSequence(program []int64, phases []int64, debug bool) (signal int64, err error) {
	readChans := make([]chan int64, len(phases))
	errChans := make([]chan error, len(phases))
	for i, phase := range phases {
		errChans[i] = make(chan error)
		readChans[i] = make(chan int64, 1)
		defer close(readChans[i])
		readChans[i] <- phase
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := range phases {
		writeChan := readChans[(i+1)%len(phases)]
		go intcode.RunProgramChans(ctx, program, readChans[i], writeChan, errChans[i],
			debug, "AMP"+strconv.Itoa(i))
	}
	readChans[0] <- 0

	for i, errChan := range errChans {
		select {
		case err := <-errChan:
			if err != nil {
				return 0, err
			}
		case <-ctx.Done():
			return 0, fmt.Errorf("timeout waiting for error channel %d to close", i)
		}
	}

	signal = <-readChans[0]
	return signal, nil
}

func bestParallelSequence(program []int64, debug bool) (signal int64, sequence []int64, err error) {
	var bestSeq []int64
	bestSignal := int64(-1)

	for _, sequence := range permutations([]int64{5, 6, 7, 8, 9}) {
		signal, err := runParallelSequence(program, sequence, debug)
		if err != nil {
			return 0, nil, err
		}
		if signal > bestSignal {
			bestSignal = signal
			bestSeq = sequence
		}
	}
	return bestSignal, bestSeq, nil
}

func permutations(items []int64) [][]int64 {
	if len(items) <= 1 {
		return [][]int64{items}
	}

	var result [][]int64
	for i, item := range items {
		others := make([]int64, len(items)-1)
		copy(others, items[:i])
		copy(others[i:], items[i+1:])
		ps := permutations(others)
		for _, p := range ps {
			val := append([]int64{item}, p...)
			result = append(result, val)
		}
	}
	return result
}

func run() error {
	program, err := intcode.ReadProgram("input")
	if err != nil {
		return err
	}

	signal, seq, err := bestSequence(program, false)
	if err != nil {
		return err
	}
	fmt.Printf("Best signal: %d (for sequence %v)\n", signal, seq)

	signal, seq, err = bestParallelSequence(program, false)
	if err != nil {
		return err
	}
	fmt.Printf("Best signal: %d (for sequence %v)\n", signal, seq)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
