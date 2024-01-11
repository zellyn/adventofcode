package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
	"golang.org/x/exp/maps"
)

type botrule struct {
	index      int
	lowBot     *int
	highBot    *int
	lowOutput  *int
	highOutput *int
}

func (b botrule) String() string {
	format := func(bot, output *int) string {
		if bot != nil {
			return "bot " + strconv.Itoa(*bot)
		}
		return "output " + strconv.Itoa(*output)
	}

	return fmt.Sprintf("[%d low:%s  high:%s]", b.index, format(b.lowBot, b.lowOutput), format(b.highBot, b.highOutput))
}

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parseInput(inputs []string) (map[int]botrule, map[int][]int, error) {

	ruleStrings, valueStrings := util.Split2(inputs, func(s string) bool {
		return strings.HasPrefix(s, "bot ")
	})

	ruleStringsAndInts, err := util.ParseStringsAndInts(ruleStrings, 12, []int{5, 10}, []int{1, 6, 11})
	if err != nil {
		return nil, nil, err
	}
	rules := make(map[int]botrule, len(ruleStringsAndInts))
	for _, sai := range ruleStringsAndInts {
		lowType, highType := sai.Strings[0], sai.Strings[1]
		index, lowIndex, highIndex := sai.Ints[0], sai.Ints[1], sai.Ints[2]

		r := botrule{index: index}

		switch lowType {
		case "output":
			r.lowOutput = &lowIndex
		case "bot":
			r.lowBot = &lowIndex
		default:
			return nil, nil, fmt.Errorf("weird low output type: %q", lowType)
		}

		switch highType {
		case "output":
			r.highOutput = &highIndex
		case "bot":
			r.highBot = &highIndex
		default:
			return nil, nil, fmt.Errorf("weird high output type: %q", highType)
		}

		rules[index] = r
	}

	valueStringsAndInts, err := util.ParseStringsAndInts(valueStrings, 6, nil, []int{1, 5})
	if err != nil {
		return nil, nil, err
	}
	values := make(map[int][]int, len(valueStringsAndInts))
	for _, sai := range valueStringsAndInts {
		values[sai.Ints[1]] = append(values[sai.Ints[1]], sai.Ints[0])
	}

	return rules, values, nil
}

func sort2(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func parts(inputs []string, val1, val2 int) (index int, product int, err error) {
	rules, values, err := parseInput(inputs)
	if err != nil {
		return 0, 0, err
	}

	maxBot := slices.Max(maps.Keys(rules))

	targetLow, targetHigh := sort2(val1, val2)
	outputs := make(map[int][]int)

	done := false

	for !done {
		done = true
		for bot := 0; bot <= maxBot; bot++ {
			vs := values[bot]
			if len(vs) < 2 {
				continue
			}
			if len(vs) > 2 {
				return 0, 0, fmt.Errorf("bot %d has %d>2 values: %v", bot, len(vs), vs)
			}
			done = false

			low, high := sort2(vs[0], vs[1])
			if low == targetLow && high == targetHigh {
				index = bot
			}
			values[bot] = nil
			rule := rules[bot]
			if rule.lowBot != nil {
				values[*rule.lowBot] = append(values[*rule.lowBot], low)
			} else {
				outputs[*rule.lowOutput] = append(outputs[*rule.lowOutput], low)
			}
			if rule.highBot != nil {
				values[*rule.highBot] = append(values[*rule.highBot], high)
			} else {
				outputs[*rule.highOutput] = append(outputs[*rule.highOutput], high)
			}
		}
	}

	oneChip := func(chips []int) int {
		if len(chips) == 0 {
			return 0
		}
		return chips[0]
	}

	product = oneChip(outputs[0]) * oneChip(outputs[1]) * oneChip(outputs[2])
	return index, product, nil
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
