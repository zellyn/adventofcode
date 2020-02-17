package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/zellyn/adventofcode/ioutil"
)

var bracketRe = regexp.MustCompile(`\[([^][]+)\]`)

func hasAbba(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i] != s[i+1] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

func tls(s string) bool {
	ins := bracketRe.FindAllStringSubmatch(s, -1)
	for _, match := range ins {
		for _, group := range match[1:] {
			if hasAbba(group) {
				return false
			}
		}
	}
	out := bracketRe.ReplaceAllString(s, "[]")
	return hasAbba(out)
}

func countValid(filename string) (int, int, error) {
	lines, err := ioutil.ReadLines(filename)
	if err != nil {
		return 0, 0, err
	}
	count1 := 0
	count2 := 0
	for _, line := range lines {
		if tls(line) {
			count1++
		}
		if ssl(line) {
			count2++
		}
	}

	return count1, count2, nil
}

func addABAs(s string, m map[string]bool) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			m[s[i:i+3]] = true
		}
	}
}

func ssl(s string) bool {
	inMap := map[string]bool{}
	ins := bracketRe.FindAllStringSubmatch(s, -1)
	for _, match := range ins {
		for _, group := range match[1:] {
			addABAs(group, inMap)
		}
	}
	out := bracketRe.ReplaceAllString(s, "^&*")
	outMap := map[string]bool{}
	addABAs(out, outMap)

	for k := range outMap {
		if inMap[k[1:2]+k[0:1]+k[1:2]] {
			return true
		}
	}

	return false
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
