package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/zellyn/adventofcode/util"
)

func always(_ string) bool {
	return true
}

func re(r string) func(string) bool {
	rg := regexp.MustCompile("^" + r + "$")
	return func(s string) bool {
		return rg.MatchString(s)
	}
}

func in(min, max int) func(string) bool {
	return func(s string) bool {
		i, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return i >= min && i <= max
	}
}

func height(s string) bool {
	num, unit := s[:len(s)-2], s[len(s)-2:]
	i, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	if unit == "cm" {
		return i >= 150 && i <= 193
	} else if unit == "in" {
		return i >= 59 && i <= 76
	}
	return false
}

var needed = map[string]func(string) bool{
	"byr": in(1920, 2002),                    // Birth Year
	"iyr": re(`201\d|2020`),                  // Issue Year
	"eyr": re(`202\d|2030`),                  // Expiration Year
	"hgt": height,                            // Height
	"hcl": re(`#[0-9a-f]{6}`),                // Hair Color
	"ecl": re(`amb|blu|brn|gry|grn|hzl|oth`), // Eye Color
	"pid": re(`\d{9}`),                       // Passport ID
}

func valid(passport []string, checkFormat bool) bool {
	stillNeed := make(map[string]bool, len(needed))
	for k := range needed {
		stillNeed[k] = true
	}
	for _, line := range passport {
		for k, v := range util.KeyValuePairs(line) {
			if k == "cid" {
				continue
			}
			delete(stillNeed, k)
			if checkFormat {
				if !needed[k](v) {
					return false
				}
			}
		}
	}
	return len(stillNeed) == 0
}

func validCount(input []string, checkFormat bool) (int, error) {
	paras := util.LinesByParagraph(input)
	count := 0
	for _, para := range paras {
		if valid(para, checkFormat) {
			count++
		}
	}
	return count, nil
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
