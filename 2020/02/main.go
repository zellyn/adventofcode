package main

import (
	"fmt"
	"os"
	"strings"

	"cuelang.org/go/pkg/strconv"
)

func parse(input string) (lower int, upper int, char byte, password string, err error) {
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		err = fmt.Errorf("weird input: %q", input)
		return
	}
	lims := strings.Split(parts[0], "-")
	if len(lims) != 2 {
		err = fmt.Errorf("weird input: %q", input)
		return
	}
	lower, err = strconv.Atoi(lims[0])
	if err != nil {
		err = fmt.Errorf("weird input: %q", input)
		return
	}
	upper, err = strconv.Atoi(lims[1])
	if err != nil {
		err = fmt.Errorf("weird input: %q", input)
		return
	}
	char = parts[1][0]
	password = parts[2]
	return
}

func valid1(input string) (bool, error) {
	lower, upper, char, password, err := parse(input)
	if err != nil {
		return false, err
	}
	count := 0
	for _, c := range []byte(password) {
		if c == char {
			count++
		}
	}
	if count < lower || count > upper {
		return false, nil
	}
	return true, nil
}

func valid1Count(inputs []string) (int, error) {
	count := 0
	for _, input := range inputs {
		v, err := valid1(input)
		if err != nil {
			return 0, err
		}
		if v {
			count++
		}
	}
	return count, nil
}

func valid2(input string) (bool, error) {
	lower, upper, char, password, err := parse(input)
	if err != nil {
		return false, err
	}

	count := 0
	if password[lower-1] == char {
		count++
	}
	if password[upper-1] == char {
		count++
	}
	return count == 1, nil
}

func valid2Count(inputs []string) (int, error) {
	count := 0
	for _, input := range inputs {
		v, err := valid2(input)
		if err != nil {
			return 0, err
		}
		if v {
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
