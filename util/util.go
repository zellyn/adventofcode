package util

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

// Number is a Float or Integer
type Number interface {
	constraints.Float | constraints.Integer
}

func LowestTrue(lowFalseStart int, pred func(int) (bool, error)) (int, error) {
	lf, err := pred(lowFalseStart)
	if err != nil {
		return 0, err
	}
	if lf {
		return 0, fmt.Errorf("lowestTrue expected pred(lowFalseStart)==false; got pred(%d)==true", lowFalseStart)
	}
	lowFalse := lowFalseStart
	highTrue := 0
	for lowFalse < MaxInt/2 {
		attempt := lowFalse * 2
		st, err := pred(attempt)
		if err != nil {
			return 0, err
		}
		if st {
			highTrue = attempt
			break
		}
		lowFalse <<= 1
	}
	if highTrue == 0 {
		return 0, fmt.Errorf("cannot find high enough value to make pred(value)==true")
	}

	for highTrue-lowFalse > 1 {
		mid := (lowFalse + highTrue) / 2
		mm, err := pred(mid)
		if err != nil {
			return 0, err
		}
		if mm {
			highTrue = mid
		} else {
			lowFalse = mid
		}
	}
	return highTrue, nil
}

func highestTrueRange(lowTrue int, highFalse int, pred func(int) (bool, error)) (int, error) {
	if lowTrue >= highFalse {
		return 0, fmt.Errorf("highestTrue(%d,%d, pred): want arg1 < arg2", lowTrue, highFalse)
	}
	lt, err := pred(lowTrue)
	if err != nil {
		return 0, err
	}
	if !lt {
		return 0, fmt.Errorf("highestTrue(%d,%d, pred): pred(%d)==false", lowTrue, highFalse, lowTrue)
	}
	hf, err := pred(highFalse)
	if err != nil {
		return 0, err
	}
	if hf {
		return 0, fmt.Errorf("highestTrue(%d,%d, pred): pred(%d)==true", lowTrue, highFalse, highFalse)
	}

	for highFalse-lowTrue > 1 {
		mid := (lowTrue + highFalse) / 2
		mm, err := pred(mid)
		if err != nil {
			return 0, err
		}
		if mm {
			lowTrue = mid
		} else {
			highFalse = mid
		}
	}
	return lowTrue, nil
}

// TrimmedLines takes a string, splits it into lines, and trims each line of starting
// and ending whitespace.
func TrimmedLines(s string) []string {
	result := strings.Split(strings.TrimSpace(s), "\n")
	for i, r := range result {
		result[i] = strings.TrimSpace(r)
	}
	return result
}

// RemoveBlanks returns a slice of strings, but trimmed, and with empty/all-whitespace
// strings removed.
func RemoveBlanks(ss []string) []string {
	var r []string

	for _, s := range ss {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		r = append(r, s)
	}

	return r
}

// GroupString returns the input string, broken into runs of consecutive characters
func GroupString(s string) []string {
	var result []string
	last := -1
	for i := range s {
		c := s[i : i+1]
		if len(result) == 0 || result[last][:1] != c {
			result = append(result, c)
			last++
		} else {
			result[last] = result[last] + c
		}
	}
	return result
}

// StringsAndInts is a set of parsed strings and ints from a line of input.
type StringsAndInts struct {
	Strings []string
	Ints    []int
}

// ParseStringsAndInts takes a slice of input lines, a slice of string field indexes,
// and a slice of int field indexes. It returns a slice of StringsAndInts structs,
// one per line.
func ParseStringsAndInts(lines []string, fields int, stringFields []int, intFields []int) ([]StringsAndInts, error) {
	var result []StringsAndInts

	for i, line := range lines {
		sai := StringsAndInts{}
		parts := strings.Split(line, " ")
		if len(parts) != fields {
			return nil, fmt.Errorf("want %d fields; got %d at line %d: %q", fields, len(parts), i, line)
		}

		for _, index := range stringFields {
			sai.Strings = append(sai.Strings, parts[index])
		}
		for _, index := range intFields {
			ii, err := strconv.Atoi(parts[index])
			if err != nil {
				return nil, fmt.Errorf("unparseable field %d at line %d (%q): %v", index, i, line, err)
			}
			sai.Ints = append(sai.Ints, ii)
		}
		result = append(result, sai)
	}

	return result, nil
}

// ParseGrid parses a set of lines of whitespacespace-separated ints into a 2D grid.
func ParseGrid(lines []string) ([][]int, error) {
	var result [][]int
	var fields int
	for i, line := range lines {
		parts := strings.Fields(line)
		if i == 0 {
			fields = len(parts)
		} else {
			if len(parts) != fields {
				return nil, fmt.Errorf("line 0 has %d fields; line %d has %d: %q", fields, i+1, len(parts), line)
			}
		}
		ints := make([]int, 0, len(parts))

		for _, part := range parts {
			theInt, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error at line %d: %w", i+1, err)
			}
			ints = append(ints, theInt)
		}
		result = append(result, ints)
	}

	return result, nil
}

// LinesByParagraph takes a slice of strings, and returns a slice of slices of
// strings: it separates paragraphs (multiple newlines).
func LinesByParagraph(lines []string) [][]string {
	var result [][]string
	var chunk []string

	for _, line := range lines {
		if line == "" {
			if chunk != nil {
				result = append(result, chunk)
				chunk = nil
			}
		} else {
			chunk = append(chunk, line)
		}
	}
	if chunk != nil {
		result = append(result, chunk)
	}
	return result
}

// KeyValuePairs splits a space-separated sequence of colon-separated key:value
// pairs into a map.
func KeyValuePairs(input string) map[string]string {
	result := make(map[string]string)
	parts := strings.Split(input, " ")
	for _, part := range parts {
		pieces := strings.SplitN(part, ":", 2)
		if len(pieces) == 2 {
			result[pieces[0]] = pieces[1]
		} else {
			result[pieces[0]] = ""
		}
	}
	return result
}

func Transpose(input [][]int) [][]int {
	var result [][]int

	for col := 0; col < len(input[0]); col++ {
		var newRow []int
		for row := range input {
			newRow = append(newRow, input[row][col])
		}
		result = append(result, newRow)
	}

	return result
}

// ReadFile is just ioutil.ReadFile
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
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
	bb, err := ReadFile(filename)
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
	lines := TrimmedLines(all)
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

// ParseLinesOfInts runs ParseInts on every string in the passed slice, passing
// back a slice of slices of ints.
func ParseLinesOfInts(commaStrings []string, separator string) ([][]int, error) {
	result := make([][]int, 0, len(commaStrings))
	for _, commaString := range commaStrings {
		ints, err := ParseInts(commaString, separator)
		if err != nil {
			return nil, err
		}
		result = append(result, ints)
	}
	return result, nil
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

// StringsToInt32s takes a slice of strings and returns a slice of int32s
func StringsToInt32s(strings []string) ([]int32, error) {
	ints := make([]int32, len(strings))

	for i, v := range strings {
		i64, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, err
		}
		ints[i] = int32(i64)
	}
	return ints, nil
}

// Reverse reverses a string.
func Reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// Sum gives the sum of a slice of ints or floats.
func Sum[T Number](items []T) T {
	var sum T
	for _, item := range items {
		sum += item
	}
	return sum
}

// MapSum returns a slice of the Sums of sublists.
func MapSum[T Number](slicesOfItems [][]T) []T {
	result := make([]T, 0, len(slicesOfItems))

	for _, items := range slicesOfItems {
		result = append(result, Sum(items))
	}
	return result
}

// Max gives the max of a slice of ints or floats.
func Max[T Number](items []T) T {
	var max T
	if len(items) > 0 {
		max = items[0]
	}
	for _, item := range items[1:] {
		if item > max {
			max = item
		}
	}
	return max
}
