package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(inputs []string) ([][2]int, error) {
	var result [][2]int

	for _, s := range strings.Split(inputs[0], ",") {
		as, bs, _ := strings.Cut(s, "-")
		a, err := strconv.Atoi(as)
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(bs)
		if err != nil {
			return nil, err
		}
		result = append(result, [2]int{a, b})
	}
	return result, nil
}

func isPalindrome(i int) bool {
	s := strconv.Itoa(i)
	return s[0:len(s)/2] == s[len(s)/2:]
}

func isRepeating(i int) bool {
	s := strconv.Itoa(i)
	ll := len(s)
OUTER:
	for i := 1; i <= ll/2; i++ {
		if ll%i != 0 {
			continue
		}
		for j := 0; j < ll-i; j++ {
			if s[j] != s[j+i] {
				continue OUTER
			}
		}
		return true
	}
	return false
}

func part1(inputs []string) (int, error) {
	pairs, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	sum := 0

	for _, pair := range pairs {
		for i := pair[0]; i <= pair[1]; i++ {
			if isPalindrome(i) {
				sum += i
			}
		}
	}

	return sum, nil
}

func part2(inputs []string) (int, error) {
	pairs, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	sum := 0

	for _, pair := range pairs {
		for i := pair[0]; i <= pair[1]; i++ {
			if isRepeating(i) {
				sum += i
			}
		}
	}

	return sum, nil
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
