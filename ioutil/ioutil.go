package ioutil

import (
	iu "io/ioutil"
	"strings"
)

// ReadFile is just ioutil.ReadFile
func ReadFile(filename string) ([]byte, error) {
	return iu.ReadFile(filename)
}

// ReadLines reads a file and returns a slice of strings, one per line.
func ReadLines(filename string) ([]string, error) {
	bb, err := iu.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bb)), "\n"), nil
}

// ReadFileString reads a file and returns it as a string, trimmed.
func ReadFileString(filename string) (string, error) {
	bb, err := iu.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bb)), nil
}

// MustReadFileString reads a string from a file or panics.
func MustReadFileString(filename string) string {
	s, err := ReadFileString(filename)
	if err != nil {
		panic(err)
	}
	return s
}
