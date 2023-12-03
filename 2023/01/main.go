package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var intRegexp = regexp.MustCompile(`[0-9]`)
var intOrNameRegexp = regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)

var digitValues = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func lineValue(s string) (int, error) {
	digits := intRegexp.FindAllString(s, -1)

	return strconv.Atoi(digits[0] + digits[len(digits)-1])
}

func lineValue2(s string) int {
	firstDigit := -1
	lastDigit := -1

	for pos := range s {
		match := intOrNameRegexp.FindString(s[pos:])
		if match == "" {
			continue
		}
		d, ok := digitValues[match]
		if !ok {
			d = int(match[0] - '0')
		}
		if firstDigit == -1 {
			firstDigit = d
		}
		lastDigit = d
	}

	return firstDigit*10 + lastDigit
}

func part1(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		val, err := lineValue(input)
		if err != nil {
			return 0, err
		}
		sum += val
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	sum := 0
	for _, input := range inputs {
		sum += lineValue2(input)
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
