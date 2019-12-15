package ioutil

import (
	"io/ioutil"
	iu "io/ioutil"
	"strings"
)

// ReadFile is just ioutil.ReadFile
func ReadFile(filename string) ([]byte, error) {
	return iu.ReadFile(filename)
}

// ReadLines reads a file and returns a slice of strings, one per line.
func ReadLines(filename string) ([]string, error) {
	bb, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bb)), "\n"), nil
}

// ReadFileString reads a file and returns it as a string, trimmed.
func ReadFileString(filename string) (string, error) {
	bb, err := ioutil.ReadFile("input")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bb)), nil
}
