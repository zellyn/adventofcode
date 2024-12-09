package main

import (
	"fmt"
	"os"
)

// var printf = func(string, ...any) {}
var printf = fmt.Printf

func parse(input string) []int {
	var result []int
	gap := false
	file := 0
	for _, c := range input {
		count := int(c - '0')
		fileToWrite := file
		if gap {
			fileToWrite = -1
			file++
		}
		for i := 0; i < count; i++ {
			result = append(result, fileToWrite)
		}
		gap = !gap
	}
	return result
}

func printDisk(disk []int) {
	for _, i := range disk {
		if i == -1 {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", i%10)
		}
	}
	fmt.Println()
}

func sum(disk []int) int {
	sum := 0
	for i, sector := range disk {
		if sector == -1 {
			continue
		}
		sum += i * sector
	}
	return sum
}

func part1(inputs []string) (int, error) {
	disk := parse(inputs[0])

	left := 0
	right := len(disk) - 1

	for {
		for disk[left] != -1 && left < right {
			left++
		}
		for disk[right] == -1 && left < right {
			right--
		}
		if left >= right {
			break
		}
		disk[left] = disk[right]
		disk[right] = -1
		// printDisk(disk)
	}

	return sum(disk), nil
}

func findGap(disk []int, want int) int {
	size := 0
	start := 0
	for i, sector := range disk {
		if sector == -1 {
			size++
			if size == want {
				return start
			}
		} else {
			size = 0
			start = i + 1
		}
	}
	return -1
}

func part2(inputs []string) (int, error) {
	disk := parse(inputs[0])

	moved := map[int]bool{-1: true}

	right := len(disk) - 1
	for {
		for right > 0 && moved[disk[right]] {
			right--
		}
		if right <= 0 {
			break
		}
		fileNo := disk[right]
		moved[fileNo] = true
		last := right
		for right > 0 && disk[right] == fileNo {
			right--
		}
		if right <= 0 {
			break
		}
		first := right + 1
		size := last - first + 1
		gapStart := findGap(disk[:first], size)
		if gapStart == -1 {
			continue
		}
		for i := 0; i < size; i++ {
			disk[gapStart+i] = fileNo
			disk[first+i] = -1
		}
	}

	return sum(disk), nil
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
