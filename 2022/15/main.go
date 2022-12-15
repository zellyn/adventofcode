package main

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/zellyn/adventofcode/geom"
	"github.com/zellyn/adventofcode/lists"
)

type sensor struct {
	pos    geom.Vec2
	beacon geom.Vec2
}

type interval struct {
	start int
	end   int
}

func parseLine(s string) sensor {
	b := sensor{}

	_, err := fmt.Sscanf(s, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &b.pos.X, &b.pos.Y, &b.beacon.X, &b.beacon.Y)
	if err != nil {
		panic(fmt.Sprintf("weird input %q: %v", s, err))
	}
	return b
}

func (b sensor) distance() int {
	return b.pos.Sub(b.beacon).AbsSum()
}

func (b sensor) intervalAt(y int) *interval {
	inLine := geom.Vec2{X: b.pos.X, Y: y}
	diff := b.distance() - inLine.Sub(b.pos).AbsSum()
	if diff < 0 {
		return nil
	}
	return &interval{start: inLine.X - diff, end: inLine.X + diff}
}

func notCovered(intervals []*interval, min, max int) (int, bool) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	result := 0

	for _, interval := range intervals {
		if interval.end < result {
			continue
		}
		if interval.start > result {
			return result, true
		}
		result = interval.end + 1
	}
	return result, result <= max
}

func (i interval) contains(n int) bool {
	return n >= i.start && n <= i.end
}

func parseInput(inputs []string) []sensor {
	return lists.Map(inputs, func(s string) sensor {
		return parseLine(s)
	})
}

func part1(inputs []string, y int) (int, error) {
	sensors := parseInput(inputs)

	beaconMap := map[geom.Vec2]bool{}
	for _, s := range sensors {
		beaconMap[s.beacon] = true
	}
	// fmt.Println("beaconMap:", beaconMap)

	intervals := lists.Map(sensors, func(s sensor) *interval { return s.intervalAt(y) })
	// fmt.Println("intervals:", intervals)
	// fmt.Println("y=", y)
	intervals = lists.Filter(intervals, func(i *interval) bool { return i != nil })
	// fmt.Println("intervals:", intervals)
	min := intervals[0].start
	max := intervals[0].end
	for _, i := range intervals {
		if i.start < min {
			min = i.start
		}
		if i.end > max {
			max = i.end
		}
	}
	count := 0
	// fmt.Printf("min=%d, max=%d\n", min, max)
OUTER:
	for x := min; x <= max; x++ {
		// fmt.Printf("Considering (%d,%d)\n", x, y)
		if beaconMap[geom.Vec2{X: x, Y: y}] {
			// fmt.Println(" found beacon: continue")
			continue
		}
		for _, i := range intervals {
			if i.contains(y) {
				count++
				continue OUTER
			}
		}
	}

	return count, nil
}

func part2(inputs []string, max int) (int, error) {
	sensors := parseInput(inputs)
	for y := 0; y <= max; y++ {
		intervals := lists.Map(sensors, func(s sensor) *interval { return s.intervalAt(y) })
		intervals = lists.Filter(intervals, func(i *interval) bool { return i != nil })
		x, ok := notCovered(intervals, 0, max)
		if ok {
			return x*4000000 + y, nil
		}
	}
	return 42, errors.New("not found 😞")
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
