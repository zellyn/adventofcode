package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var r = regexp.MustCompile(`-?[0-9]+`)

func sum(s string) (int, error) {
	total := 0
	matches := r.FindAllString(s, -1)
	for _, match := range matches {
		i, err := strconv.Atoi(match)
		if err != nil {
			return 0, err
		}
		total += i
	}
	return total, nil
}

func subsum(data interface{}) int {
	switch v := data.(type) {
	case float64:
		return int(v)
	case string, nil:
		return 0
	case []interface{}:
		sum := 0
		for _, dd := range v {
			sum += subsum(dd)
		}
		return sum
	case map[string]interface{}:
		sum := 0
		for _, dd := range v {
			if s, ok := dd.(string); ok && s == "red" {
				return 0
			}
			sum += subsum(dd)
		}
		return sum
	default:
		panic("weird type")
	}
}

func sumRed(s string) (int, error) {
	var payload interface{}
	if err := json.Unmarshal([]byte(s), &payload); err != nil {
		return 0, err
	}
	return subsum(payload), nil
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
