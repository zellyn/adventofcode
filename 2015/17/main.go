package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func waysToAddTo(nums []int, goal, sum int) int {
	if sum == goal {
		return 1
	}
	if sum < goal {
		return 0
	}

	waysWithoutFirst := waysToAddTo(nums[1:], goal, sum-nums[0])

	if nums[0] == goal {
		return 1 + waysWithoutFirst
	}
	if nums[0] > goal {
		return waysWithoutFirst
	}

	return waysToAddTo(nums[1:], goal-nums[0], sum-nums[0]) + waysWithoutFirst
}

func waysToAddToUsing(nums []int, goal, sum, using int) int {
	if goal == 0 && using == 0 {
		return 1
	}
	if sum < goal {
		return 0
	}
	if using > len(nums) {
		return 0
	}
	if using <= 0 {
		return 0
	}
	waysWithoutFirst := waysToAddToUsing(nums[1:], goal, sum-nums[0], using)

	if nums[0] > goal {
		return waysWithoutFirst
	}

	if nums[0] == goal {
		if using == 1 {
			return 1 + waysWithoutFirst
		}
		return waysWithoutFirst
	}

	ways := waysToAddToUsing(nums[1:], goal-nums[0], sum-nums[0], using-1) + waysWithoutFirst
	return ways
}

func parse(input []string) ([]int, int, error) {
	var nums []int
	sum := 0
	for _, line := range input {
		ii, err := strconv.Atoi(line)
		if err != nil {
			return nil, 0, err
		}
		nums = append(nums, ii)
		sum += ii
	}
	sort.Ints(nums)
	for i := len(nums)/2 - 1; i >= 0; i-- {
		opp := len(nums) - 1 - i
		nums[i], nums[opp] = nums[opp], nums[i]
	}
	return nums, sum, nil
}

func combinations(input []string, goal int) (int, error) {
	nums, sum, err := parse(input)
	if err != nil {
		return 0, err
	}
	return waysToAddTo(nums, goal, sum), nil
}

func smallestCombinations(input []string, goal int) (int, error) {
	nums, sum, err := parse(input)
	if err != nil {
		return 0, err
	}
	for i := range nums {
		ways := waysToAddToUsing(nums, goal, sum, i+1)
		if ways > 0 {
			return ways, nil
		}
	}
	return 0, nil
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
