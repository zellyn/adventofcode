package main

import (
	"fmt"
	"math/bits"
	"os"

	"github.com/zellyn/adventofcode/ioutil"
	"github.com/zellyn/adventofcode/math"
)

func readInput(filename string) ([]int, error) {
	s, err := ioutil.ReadFileString(filename)
	if err != nil {
		return nil, err
	}
	return ioutil.ParseInts(s, "\n")
}

func canAddTo(target int, mask uint32, sum int, ints []int, memo map[int]map[uint32]bool) bool {
	for m, res := range memo[target] {
		if mask&m == mask {
			return res
		}
	}
	if target > sum {
		return false
	}
	if target == 0 {
		return true
	}
	if mask == 0 {
		return false
	}
	last := bits.TrailingZeros32(mask)
	newMask := mask &^ (1 << last)

	result := canAddTo(target-ints[last], newMask, sum-ints[last], ints, memo) || canAddTo(target, newMask, sum-ints[last], ints, memo)
	if memo[target] == nil {
		memo[target] = map[uint32]bool{}
	}
	memo[target][mask] = result

	return result
}

func waysToAddTo(target int, availableMask uint32, usedMask uint32, sum int, ints []int, maxCount int) []uint32 {
	debug := false
	if debug {
		fmt.Printf("waysToAddTo(target=%d, mask=%d, sum=%d, ints, maxCount=%d) {\n", target, availableMask, sum, maxCount)
	}
	if target == 0 {
		if debug {
			fmt.Printf("return [%d]\n}\n", availableMask)
		}
		return []uint32{usedMask}
	}
	if target > sum || maxCount == 0 || availableMask == 0 {
		if debug {
			fmt.Printf("return []\n}\n")
		}
		return nil
	}

	last := bits.TrailingZeros32(availableMask)
	bit := uint32(1) << last
	theInt := ints[len(ints)-1-last]
	if debug {
		fmt.Printf("with %d\n", theInt)
	}
	with := waysToAddTo(target-theInt, availableMask&^bit, usedMask|bit, sum-theInt, ints, maxCount-1)
	if debug {
		fmt.Printf("with := %v\n", with)
	}
	if debug {
		fmt.Printf("without %d\n", theInt)
	}
	without := waysToAddTo(target, availableMask&^bit, usedMask, sum-theInt, ints, maxCount)
	if debug {
		fmt.Printf("without := %v\n", without)
	}
	result := make([]uint32, len(with)+len(without))
	copy(result, with)
	copy(result[len(with):], without)
	if debug {
		fmt.Printf("return %v\n}\n", result)
	}
	return result
}

func entanglement(ints []int, mask uint32) int {
	prod := 1
	pos := len(ints) - 1
	for m := mask; m > 0; m >>= 1 {
		if m&1 > 0 {
			prod *= ints[pos]
		}
		pos--
	}
	return prod
}

func sum(ints []int) int {
	sum := 0
	for _, ii := range ints {
		sum += ii
	}
	return sum
}

func canPartitionHelper(ways []uint32, partitionsLeft int, fullMask uint32, mask uint32, result map[uint32]bool) bool {
	if partitionsLeft == 0 {
		return fullMask == mask
	}
	success := false
	for first := 0; first <= len(ways)-partitionsLeft; first++ {
		if mask&ways[first] > 0 {
			continue
		}
		if canPartitionHelper(ways[first+1:], partitionsLeft-1, fullMask, mask|ways[first], result) {
			result[ways[first]] = true
			success = true
		}
	}
	return success
}

func canPartitionN(ways []uint32, partitions int, bitCount int) map[uint32]bool {
	fullMask := uint32(1)<<bitCount - 1
	result := map[uint32]bool{}

	canPartitionHelper(ways, partitions, fullMask, 0, result)

	return result
}

func othersCanPartitionNways(usedSoFar uint32, ways []uint32, n int) bool {
	if n == 1 {
		return true
	}
	for _, way := range ways {
		if way&usedSoFar > 0 {
			continue
		}
		if othersCanPartitionNways(usedSoFar|way, ways, n-1) {
			return true
		}
	}
	return false
}

func best(ints []int, partitions int) int {
	mask := uint32(1)<<len(ints) - 1
	sum := 0
	for _, ii := range ints {
		sum += ii
	}
	target := sum / partitions
	fmt.Printf("target=%d\n", target)

	ways := waysToAddTo(target, mask, 0, sum, ints, len(ints)/partitions)
	if len(ways) == 0 {
		return -1
	}
	fmt.Printf("len(ways)=%d\n", len(ways))

	min := len(ints)
	bestEntanglement := math.MaxInt

	for _, way := range ways {
		count := bits.OnesCount32(way)
		if count > min {
			continue
		}
		if !othersCanPartitionNways(way, ways, partitions-1) {
			continue
		}

		if count < min {
			min = count
			bestEntanglement = math.MaxInt
		}
		// fmt.Printf("Considering %b\n", way)
		ent := entanglement(ints, way)
		if ent < bestEntanglement {
			bestEntanglement = ent
		}
	}
	return bestEntanglement
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
