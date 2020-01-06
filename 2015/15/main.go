package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

func waysToAddTo(sum, parts int) [][]int {
	if parts == 1 {
		return [][]int{{sum}}
	}
	if sum == 0 {
		res := make([]int, parts)
		return [][]int{res}
	}

	var result [][]int

	for i := 0; i <= sum; i++ {
		for _, sub := range waysToAddTo(sum-i, parts-1) {
			result = append(result, append([]int{i}, sub...))
		}
	}
	return result
}

func score(amounts []int, info []util.StringsAndInts, calories int) int {
	if calories > 0 {
		cals := 0
		for ingredient := 0; ingredient < len(info); ingredient++ {
			cals += amounts[ingredient] * info[ingredient].Ints[4]
		}
		if cals != calories {
			return 0
		}
	}

	prod := 1
	for property := 0; property < len(info[0].Ints)-1; property++ {
		sum := 0
		for ingredient := 0; ingredient < len(info); ingredient++ {
			// fmt.Printf(" %d * %d\n", amounts[ingredient], info[ingredient].Ints[property])
			sum += amounts[ingredient] * info[ingredient].Ints[property]
		}
		// fmt.Printf("sum=%d\n", sum)
		if sum <= 0 {
			return 0
		}

		prod *= sum
	}
	return prod
}

func parseInput(input []string) ([]util.StringsAndInts, error) {
	for i, s := range input {
		input[i] = strings.ReplaceAll(s, ",", "")
	}
	return util.ParseStringsAndInts(input, 11, []int{0}, []int{2, 4, 6, 8, 10})
}

func best(input []string, calories int) (int, error) {
	parsed, err := parseInput(input)
	if err != nil {
		return 0, err
	}
	ways := waysToAddTo(100, len(parsed))

	max := 0

	for _, amounts := range ways {
		s := score(amounts, parsed, calories)
		if s > max {
			max = s
		}
	}
	return max, nil
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
