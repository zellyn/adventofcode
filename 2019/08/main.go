package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/ioutil"
)

func count(s string, c rune) int {
	sum := 0
	for _, cc := range s {
		if cc == c {
			sum++
		}
	}
	return sum
}

func split(s string, size int) []string {
	var result []string
	for i := 0; i < len(s); i += size {
		result = append(result, s[i:i+size])
	}
	return result
}

func merge(front, back string) string {
	result := ""
	for i := range front {
		if front[i] == '2' {
			result += string(back[i])
		} else {
			result += string(front[i])
		}
	}
	return result
}

func render(image string) {
	parts := split(image, 25)
	for _, part := range parts {
		fmt.Println(strings.Replace(part, "1", "\033[7m1\033[m", -1))
	}
}

func run() error {
	input, err := ioutil.ReadFileString("input")
	if err != nil {
		return err
	}

	min := 25*6 + 1
	prod := 0
	parts := split(input, 25*6)
	for _, part := range parts {
		if c := count(part, '0'); c < min {
			min = c
			prod = count(part, '1') * count(part, '2')
		}
	}

	fmt.Println(prod)

	image := parts[0]
	for _, part := range parts {
		image = merge(image, part)
	}
	render(image)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
