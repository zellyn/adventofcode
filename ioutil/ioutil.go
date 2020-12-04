package ioutil

import (
	iu "io/ioutil"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

// ReadFile is just ioutil.ReadFile
func ReadFile(filename string) ([]byte, error) {
	return iu.ReadFile(filename)
}

// ReadLines reads a file and returns a slice of strings, one per line.
func ReadLines(filename string) ([]string, error) {
	s, err := ReadFileString(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(s, "\n"), nil
}

// MustReadLines reads a file and returns a slice of strings, one per line, or dies.
// MustReadFileString reads a string from a file or panics.
func MustReadLines(filename string) []string {
	s, err := ReadLines(filename)
	if err != nil {
		panic(err)
	}
	return s
}

// ReadFileString reads a file and returns it as a string, trimmed.
func ReadFileString(filename string) (string, error) {
	bb, err := iu.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(bb), " \t\r\n"), nil
}

// MustReadFileString reads a string from a file or panics.
func MustReadFileString(filename string) string {
	s, err := ReadFileString(filename)
	if err != nil {
		panic(err)
	}
	return s
}

// ReadFileInts reads a file of ints, one per line
func ReadFileInts(filename string) ([]int, error) {
	all, err := ReadFileString(filename)
	if err != nil {
		return nil, err
	}
	lines := util.TrimmedLines(all)
	return StringsToInts(lines)
}

// MustReadFileInts reads a file of ints, one per line, or panics.
func MustReadFileInts(filename string) []int {
	ints, err := ReadFileInts(filename)
	if err != nil {
		panic(err)
	}
	return ints
}

// ParseInts parses a string of separated ints into a slice of ints.
func ParseInts(commaString string, separator string) ([]int, error) {
	input := strings.TrimSpace(commaString)
	entries := strings.Split(input, separator)
	ints := make([]int, len(entries))
	for i, v := range entries {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ints[i] = int(i64)
	}
	return ints, nil
}

// MustStringsToInts takes a slice of strings and returns a slice of ints, or panics.
func MustStringsToInts(strings []string) []int {
	ints, err := StringsToInts(strings)
	if err != nil {
		panic(err)
	}
	return ints
}

// StringsToInts takes a slice of strings and returns a slice of ints
func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))

	for i, v := range strings {
		i64, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		ints[i] = int(i64)
	}
	return ints, nil
}
