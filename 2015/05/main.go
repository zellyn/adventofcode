package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/ioutil"
)

func isNice(s string) bool {
	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
		return false
	}
	if strings.Count(s, "a")+strings.Count(s, "e")+strings.Count(s, "i")+strings.Count(s, "o")+strings.Count(s, "u") < 3 {
		return false
	}
	var last rune
	for _, r := range s {
		if last == r {
			return true
		}
		last = r
	}
	return false
}

func isNice2(s string) bool {
	repeat := false
	pair := false
	where := map[[2]byte]int{}
	for i, r := range s {
		if i >= 2 && rune(s[i-2]) == r {
			repeat = true
		}
		if i >= 1 {
			key := [2]byte{s[i-1], s[i]}
			if where[key] > 0 {
				if where[key] < i-1 {
					pair = true
				}
			} else {
				where[key] = i
			}
		}
	}
	return repeat && pair
}

func run() error {
	lines, err := ioutil.ReadLines("input")
	if err != nil {
		return err
	}

	nice := 0
	nice2 := 0
	for _, line := range lines {
		if isNice(line) {
			nice++
		}
		if isNice2(line) {
			nice2++
		}
	}
	fmt.Println("Nice lines:", nice)
	fmt.Println("Nice lines (new rules):", nice2)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
