package main

import (
	"fmt"
	"github.com/zellyn/adventofcode/ioutil"
	"os"
)

func escape(s string) string {
	result := ""
	for _, r := range s {
		switch r {
		case '"':
			result += `\"`
		case '\\':
			result += `\\`
		default:
			result += string(r)
		}
	}
	return `"` + result + `"`
}

func unescape(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			if s[i+1] == 'x' {
				result += "x"
				i += 3
			} else {
				result += s[i+1 : i+2]
				i++
			}
		} else {
			result += s[i : i+1]
		}
	}
	return result
}

func run() error {
	lines, err := ioutil.ReadLines("input")
	if err != nil {
		return err
	}
	totalChars := 0
	unescapedChars := 0
	escapedChars := 0
	for _, line := range lines {
		totalChars += len(line)
		unescapedChars += len(unescape(line[1 : len(line)-1]))
		escapedChars += len(escape(line))
	}
	fmt.Println(totalChars - unescapedChars)
	fmt.Println(escapedChars - totalChars)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
