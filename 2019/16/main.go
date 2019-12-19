package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zellyn/adventofcode/ioutil"
)

func split(s string) []int32 {
	result := make([]int32, 0, len(s))
	for _, c := range s {
		result = append(result, int32(c-'0'))
	}

	return result
}

func join(int32s []int32) string {
	result := ""
	for _, i := range int32s {
		result += string(rune(i) + '0')
	}
	return result
}

func phase(int32s []int32) []int32 {
	return subRep(int32s, 0)
}

func splitRep(s string, l int, offset int) []int32 {
	result := make([]int32, l)
	for i := 0; i < l; i++ {
		c := s[(i+offset)%len(s)]
		result[i] = int32(c - '0')
	}
	return result
}

func subRep(int32s []int32, offset int) []int32 {
	l := len(int32s)
	result := make([]int32, l)

	for i := 0; i < l; i++ {
		if i%(l/100+1) == 0 {
			fmt.Print(".")
		}
		// fmt.Printf("i=%d\n", i)
		// lastMul := int32(0)
		for j := i; j < l; {
			nextBlock := (j+1+offset)/(i+1+offset)*(i+1+offset) + i
			if nextBlock > l {
				nextBlock = l
			}
			mul := int32((2 - ((j+1+offset)/(i+1+offset))%4) % 2)
			if mul == 0 {
				j = nextBlock
				continue
			}
			for ; j < nextBlock; j++ {
				result[i] += int32s[j] * mul
			}
			j += i + offset + 1
		}
	}
	for i, j := range result {
		if j < 0 {
			j = -j
		}
		result[i] = j % 10
	}
	return result
}

func special(int32s []int32) {
	sum := int32(0)
	for i := len(int32s) - 1; i >= 0; i-- {
		sum = (sum + int32s[i]) % 10
		int32s[i] = sum
	}
}

func run() error {
	input := ioutil.MustReadFileString("input")
	offset64, err := strconv.ParseInt(input[:7], 10, 32)
	if err != nil {
		return err
	}
	offset := int(offset64)
	int32s := splitRep(input, 10000*len(input)-offset, offset)
	for i := 0; i < 100; i++ {
		special(int32s)
	}
	fmt.Println(join(int32s[:8]))
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
