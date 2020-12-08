package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagcount struct {
	name  string
	count int
}

type baginfo struct {
	name     string
	contains []bagcount
}

func waysToGold(rules []string) (int, error) {
	info, err := parseRules(rules)
	if err != nil {
		return 0, err
	}
	count := 0
	goldCache := make(map[string]bool)
	for _, i := range info {
		if hasGold(i.name, info, goldCache) {
			count++
		}
	}
	return count, nil
}

func sizeOfGold(rules []string) (int, error) {
	info, err := parseRules(rules)
	if err != nil {
		return 0, err
	}
	sizeCache := make(map[string]int)
	return size("shiny gold", info, sizeCache) - 1, nil
}

func hasGold(name string, info map[string]baginfo, goldCache map[string]bool) bool {
	if has, ok := goldCache[name]; ok {
		return has
	}

	for _, bc := range info[name].contains {
		if bc.name == "shiny gold" || hasGold(bc.name, info, goldCache) {
			goldCache[name] = true
			return true
		}

	}
	goldCache[name] = false
	return false
}

func size(name string, info map[string]baginfo, sizeCache map[string]int) int {
	if totalSize, ok := sizeCache[name]; ok {
		return totalSize
	}

	totalSize := 1
	for _, bc := range info[name].contains {
		subSize, ok := sizeCache[bc.name]
		if !ok {
			subSize = size(bc.name, info, sizeCache)
		}
		totalSize += bc.count * subSize
	}
	sizeCache[name] = totalSize
	return totalSize
}

func parseRules(inputs []string) (map[string]baginfo, error) {
	result := make(map[string]baginfo, len(inputs))
	for _, input := range inputs {
		bi, err := parseBag(input)
		if err != nil {
			return nil, err
		}
		result[bi.name] = bi
	}
	return result, nil
}

func parseBag(input string) (baginfo, error) {
	result := baginfo{}

	halves := strings.Split(input, " bags contain ")
	if len(halves) != 2 || input[len(input)-1] != '.' {
		return result, fmt.Errorf("weird input: %q", input)
	}
	result.name = halves[0]
	if halves[1] == "no other bags." {
		return result, nil
	}

	contains := strings.Split(halves[1][:len(halves[1])-1], ", ")
	for _, contain := range contains {
		parts := strings.SplitN(contain, " ", 2)
		i, err := strconv.Atoi(parts[0])
		if err != nil || len(parts) != 2 {
			return result, fmt.Errorf("weird bag contents: %q", contain)
		}
		name := parts[1][:strings.LastIndex(parts[1], " ")]
		result.contains = append(result.contains, bagcount{name: name, count: i})
	}

	return result, nil
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
