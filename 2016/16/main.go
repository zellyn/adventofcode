package main

import (
	"fmt"
	"os"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func lengthen(s string) string {
	ll := len(s)
	l2 := ll * 2
	res := make([]byte, l2+1)
	copy(res, []byte(s))
	res[len(s)] = '0'
	for i, b := range res[:ll] {
		res[l2-i] = b ^ 1
	}
	return string(res)
}

func checksum(s string) string {
	if len(s)%2 == 1 || s == "" {
		return s
	}

	res := make([]byte, 0, len(s)/2)

	for i := 0; i < len(s)/2; i++ {
		if s[i*2] == s[i*2+1] {
			res = append(res, '1')
		} else {
			res = append(res, '0')
		}
	}

	return checksum(string(res))
}

func part1(input string, size int) (string, error) {
	s := input
	for len(s) < size {
		s = lengthen(s)
	}

	s = s[:size]
	return checksum(s), nil
}

func part2(inputs []string) (int, error) {
	return 42, nil
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
