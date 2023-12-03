package main

import (
	"fmt"
	"os"

	"github.com/zellyn/adventofcode/charmap"
	"github.com/zellyn/adventofcode/geom"
)

type number struct {
	start, end geom.Vec2
	num        int
}

func (n number) hasAdjacentSymbol(m charmap.M) bool {
	for v := n.start; v.X <= n.end.X; v = v.E() {
		for _, neighbor := range v.Neighbors8() {
			if isSymbol(m[neighbor]) {
				return true
			}
		}
	}
	return false
}

func (n number) positions() []geom.Vec2 {
	var result []geom.Vec2

	for v := n.start; v.X <= n.end.X; v = v.E() {
		result = append(result, v)
	}
	return result
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r rune) bool {
	if r == 0 || r == '.' || isDigit(r) {
		return false
	}
	return true
}

func findNumbers(m charmap.M) []number {
	seen := make(map[geom.Vec2]bool)
	var result []number

	for pos, char := range m {
		if !isDigit(char) || seen[pos] {
			continue
		}
		seen[pos] = true
		start := pos
		end := pos
		for isDigit(m[start.W()]) {
			start = start.W()
			seen[start] = true
		}
		for isDigit(m[end.E()]) {
			end = end.E()
			seen[end] = true
		}

		n := number{
			start: start,
			end:   end,
		}

		for _, v := range n.positions() {
			n.num = n.num*10 + int(m[v]-'0')
		}

		result = append(result, n)
	}

	return result
}

func makeNumMap(nums []number) map[geom.Vec2]number {
	result := make(map[geom.Vec2]number)
	for _, num := range nums {
		for _, v := range num.positions() {
			result[v] = num
		}
	}

	return result
}

func neighborNums(pos geom.Vec2, numMap map[geom.Vec2]number) []number {
	seen := make(map[geom.Vec2]bool)
	var result []number

	for _, p := range pos.Neighbors8() {
		num, ok := numMap[p]
		if !ok || seen[num.start] {
			continue
		}
		seen[num.start] = true
		result = append(result, num)
	}

	return result
}

func part1(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	nums := findNumbers(m)
	sum := 0
	for _, num := range nums {
		if num.hasAdjacentSymbol(m) {
			sum += num.num
		}
	}
	return sum, nil
}

func part2(inputs []string) (int, error) {
	m := charmap.Parse(inputs)
	nums := findNumbers(m)
	numMap := makeNumMap(nums)
	sum := 0
	for pos, char := range m {
		if char != '*' {
			continue
		}
		nns := neighborNums(pos, numMap)
		if len(nns) != 2 {
			continue
		}
		sum += nns[0].num * nns[1].num
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
