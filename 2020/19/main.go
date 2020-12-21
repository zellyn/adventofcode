package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

type rule struct {
	ints    []int
	altInts []int
	s       string
}

var intRe = regexp.MustCompile(`\d+`)

func parseRules(inputs []string) (map[int]rule, error) {
	result := make(map[int]rule, len(inputs))
	for _, input := range inputs {
		intStrs := intRe.FindAllString(input, -1)
		ints, err := util.StringsToInts(intStrs)
		if err != nil {
			return nil, fmt.Errorf("weird rule: %q: %w", input, err)
		}
		r := rule{
			ints: ints[1:],
		}
		parts := strings.Split(input, " ")
		if strings.Contains(input, "|") {
			switch len(parts) {
			case 4:
				if parts[2] != "|" {
					return nil, fmt.Errorf("weird alternating rule: %q", input)
				}
				r.ints = ints[1:2]
				r.altInts = ints[2:3]
				result[ints[0]] = r
				continue
			case 6:
				if parts[3] != "|" {
					return nil, fmt.Errorf("weird alternating rule: %q", input)
				}
				r.ints = ints[1:3]
				r.altInts = ints[3:5]
				result[ints[0]] = r
				continue
			default:
				return nil, fmt.Errorf("weird alternating rule: %q", input)
			}
		}
		switch len(ints) {
		case 1:
			if len(parts) != 2 || len(parts[1]) != 3 || parts[1][0] != '"' || parts[1][2] != '"' {
				return nil, fmt.Errorf("weird no-reference rule: %q", input)
			}
			r.s = parts[1][1:2]
		case 2, 3, 4:
			// normal rule
		default:
			return nil, fmt.Errorf("weird input: %q", input)
		}
		result[ints[0]] = r
	}
	return result, nil
}

func parse(inputs []string) (map[int]rule, []string, error) {
	paras := util.LinesByParagraph(inputs)
	if len(paras) != 2 {
		return nil, nil, fmt.Errorf("want 2 paras; got %d", len(paras))
	}
	rules, err := parseRules(paras[0])
	if err != nil {
		return nil, nil, err
	}
	return rules, paras[1], nil
}

func altReFor(i int, rules map[int]rule) (string, error) {
	r42, err := reFor(42, rules, true)
	if err != nil {
		return "", err
	}
	r42 = "(?:" + r42 + ")"
	// Rule 8 is easy: just one or more rule 42s.
	if i == 8 {
		return r42 + "+", nil
	}
	// Rule 11 is nested rule 42s and 31s. Eight is enough.
	r31, err := reFor(31, rules, true)
	if err != nil {
		return "", err
	}
	r31 = "(?:" + r31 + ")"
	return fmt.Sprintf("%s%s|%s%s%s%s|%s%s%s%s%s%s|%s%s%s%s%s%s%s%s", r42, r31, r42, r42, r31, r31, r42, r42, r42, r31, r31, r31, r42, r42, r42, r42, r31, r31, r31, r31), nil
}

func reFor(i int, rules map[int]rule, sub bool) (string, error) {
	if sub && (i == 8 || i == 11) {
		return altReFor(i, rules)
	}
	r, ok := rules[i]
	if !ok {
		return "", fmt.Errorf("cannot find rule %d", i)
	}

	if r.s != "" {
		return r.s, nil
	}

	result := ""

	for _, ii := range r.ints {
		s, err := reFor(ii, rules, sub)
		if err != nil {
			return "", fmt.Errorf("error while building rule %d: %w", i, err)
		}
		if len(s) == 1 {
			result += s
		} else {
			result += fmt.Sprintf("(?:%s)", s)
		}
	}

	if len(r.altInts) > 0 {
		result += "|"
		for _, ii := range r.altInts {
			s, err := reFor(ii, rules, sub)
			if err != nil {
				return "", fmt.Errorf("error while building rule %d: %w", i, err)
			}
			if len(s) == 1 {
				result += s
			} else {
				result += fmt.Sprintf("(?:%s)", s)
			}
		}
	}

	return result, nil
}

func minLen(i int, rules map[int]rule) (int, error) {
	r, ok := rules[i]
	if !ok {
		return 0, fmt.Errorf("cannot find rule %d", i)
	}

	if r.s != "" {
		return 1, nil
	}

	count := 0
	for _, ii := range r.ints {
		cc, err := minLen(ii, rules)
		if err != nil {
			return 0, err
		}
		count += cc
	}

	if len(r.altInts) > 0 {
		count2 := 0
		for _, ii := range r.altInts {
			cc, err := minLen(ii, rules)
			if err != nil {
				return 0, err
			}
			count2 += cc
		}
		if count2 < count {
			return count2, nil
		}
	}

	return count, nil
}

func part1(inputs []string) (int, error) {
	rules, strings, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	reStr, err := reFor(0, rules, false)
	if err != nil {
		return 0, err
	}
	re, err := regexp.Compile("^" + reStr + "$")
	if err != nil {
		return 0, err
	}

	count := 0
	for _, s := range strings {
		if re.MatchString(s) {
			count++
		}
	}
	return count, nil
}

func part2(inputs []string) (int, error) {
	rules, strings, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	reStr, err := reFor(0, rules, true)
	if err != nil {
		return 0, err
	}
	re, err := regexp.Compile("^" + reStr + "$")
	if err != nil {
		return 0, err
	}

	count := 0
	for _, s := range strings {
		if re.MatchString(s) {
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
