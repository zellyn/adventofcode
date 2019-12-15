package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readLines(filename string) ([]string, error) {
	bb, err := ioutil.ReadFile("input")
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bb)), "\n"), nil
}

func pointsForLine(line string) (map[[2]int]int, error) {
	result := map[[2]int]int{
		{0, 0}: 0,
	}
	x, y, steps := 0, 0, 0
	for _, part := range strings.Split(line, ",") {
		xi, yi := 0, 0
		switch part[0] {
		case 'L':
			xi, yi = -1, 0
		case 'R':
			xi, yi = 1, 0
		case 'U':
			xi, yi = 0, 1
		case 'D':
			xi, yi = 0, -1
		default:
			return nil, fmt.Errorf("unknown direction: '%c'", part[0])
		}
		dist, err := strconv.Atoi(part[1:])
		if err != nil {
			return nil, err
		}
		for i := 0; i < dist; i++ {
			x += xi
			y += yi
			steps++
			result[[2]int{x, y}] = steps
		}
	}
	return result, nil
}

func manhattan(x, y int) int {
	dist := 0
	if x >= 0 {
		dist += x
	} else {
		dist -= x
	}
	if y >= 0 {
		dist += y
	} else {
		dist -= y
	}
	return dist
}

func closestIntersection(line1 string, line2 string) (int, error) {
	p1, err := pointsForLine(line1)
	if err != nil {
		return 0, err
	}

	p2, err := pointsForLine(line2)
	if err != nil {
		return 0, err
	}

	min := int(^uint(0) >> 1)
	for k := range p1 {
		if p2[k] > 0 {
			x, y := k[0], k[1]
			dist := manhattan(x, y)
			if dist < min {
				min = dist
			}
		}
	}
	return min, nil
}

func soonestIntersection(line1 string, line2 string) (int, error) {
	p1, err := pointsForLine(line1)
	if err != nil {
		return 0, err
	}

	p2, err := pointsForLine(line2)
	if err != nil {
		return 0, err
	}

	min := int(^uint(0) >> 1)
	for k, v := range p1 {
		if p2[k] > 0 {
			steps := v + p2[k]
			if steps < min {
				min = steps
			}
		}
	}
	return min, nil
}

func run() error {

	testDistance, err := closestIntersection("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83")
	if err != nil {
		return err
	}
	fmt.Printf("Test distance 1 (want 159): %d\n", testDistance)

	testDistance, err = closestIntersection("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	if err != nil {
		return err
	}
	fmt.Printf("Test distance 2 (want 135): %d\n", testDistance)

	lines, err := readLines("input")
	if err != nil {
		return err
	}

	distance, err := closestIntersection(lines[0], lines[1])
	if err != nil {
		return err
	}

	fmt.Printf("Part 1 distance: %d\n", distance)

	testDistance, err = soonestIntersection("R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83")
	if err != nil {
		return err
	}
	fmt.Printf("Test distance 3 (want 610): %d\n", testDistance)

	testDistance, err = soonestIntersection("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")
	if err != nil {
		return err
	}
	fmt.Printf("Test distance 4 (want 410): %d\n", testDistance)

	distance, err = soonestIntersection(lines[0], lines[1])
	if err != nil {
		return err
	}

	fmt.Printf("Part 2 distance: %d\n", distance)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
