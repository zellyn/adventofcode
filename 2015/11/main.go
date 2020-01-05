package main

import (
	"fmt"
	"os"
)

//         "abcdefghijklmnopqrstuvwxyz"
var next = "bcdefghjjkmmnppqrstuvwxyza"

func increment(s string) string {
	r := []byte(s)

	for i := len(r) - 1; i >= 0; i-- {
		r[i] = next[r[i]-'a']
		if r[i] != 'a' {
			return string(r)
		}
	}

	return string(r)
}

func valid(s string) bool {
	lastIncrement := false
	triple := false
	lastDouble := -1
	twoDoubles := false
	for i := 1; i < len(s); i++ {
		if s[i]-s[i-1] == 1 {
			if lastIncrement {
				triple = true
			}
			lastIncrement = true
		} else {
			lastIncrement = false
		}

		if s[i] == s[i-1] {
			if lastDouble == -1 {
				lastDouble = i
			} else if lastDouble < i-1 {
				twoDoubles = true
			}
		}

		if triple && twoDoubles {
			return true
		}
	}
	return false
}

func nextValid(s string) string {
	for {
		s = increment(s)
		if valid(s) {
			return s
		}
	}
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
