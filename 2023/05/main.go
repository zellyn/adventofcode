package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/zellyn/adventofcode/util"
)

type interval struct {
	start int
	end   int
}

type span struct {
	start  int
	end    int
	offset int
}

type spanMap struct {
	name  string
	spans []span
}

type def struct {
	seeds []int
	maps  []spanMap
}

func (i interval) offset(offset int) interval {
	return interval{
		start: i.start + offset,
		end:   i.end + offset,
	}
}

func (i interval) overlap(other interval) bool {
	if i.end < other.start || i.start > other.end {
		return false
	}
	return true
}

func (i interval) String() string {
	return fmt.Sprintf("(%d,%d)", i.start, i.end)
}

func (i interval) combine(other interval) interval {
	if !i.overlap(other) {
		panic(fmt.Sprintf("non-overlapping intervals %s and %s", i, other))
	}

	min := i.start
	if other.start < min {
		min = other.start
	}
	max := i.end
	if other.end > max {
		max = other.end
	}

	return interval{start: min, end: max}
}

func (sm spanMap) mapNum(i int) int {
	for _, s := range sm.spans {
		if i >= s.start && i <= s.end {
			return i + s.offset
		}
		if i < s.start {
			return i
		}
	}
	return i
}

func (sm spanMap) mapInterval(i interval) []interval {
	// fmt.Printf("    mapping %s with map: %s\n", i, sm.name)
	var result []interval

	for _, s := range sm.spans {
		// below this interval: identity
		if s.start > i.end {
			result = append(result, i)
			return result
		}
		// above this interval: try next one
		if s.end < i.start {
			continue
		}

		// s.start <= i.end
		// s.end >= i.start

		// i.start <= s.end
		// i.end >= s.start

		if i.start < s.start {
			result = append(result, interval{i.start, s.start - 1})
			i.start = s.start
		}

		// i.start >= s.start
		// i.start <= s.end
		// i.end >= s.start

		if i.end <= s.end {
			result = append(result, i.offset(s.offset))
			return result
		}

		// i.start >= s.start
		// i.start <= s.end
		// i.end > s.end

		result = append(result, interval{start: i.start + s.offset, end: s.end + s.offset})
		i.start = s.end + 1
	}

	if i.start < i.end {
		result = append(result, i)
	}

	return result
}

func (sm spanMap) mapIntervals(intervals []interval) []interval {
	// fmt.Printf("mapping %v\n", intervals)
	var accumulated []interval

	for _, i := range intervals {
		accumulated = append(accumulated, sm.mapInterval(i)...)
	}

	if len(accumulated) == 1 {
		// fmt.Printf("  accumulated: %v\n", accumulated)
		return accumulated
	}

	// now we need to coalesce
	sort.Slice(accumulated, func(i, j int) bool { return accumulated[i].start < accumulated[j].start })

	// fmt.Printf("  accumulated: %v\n", accumulated)

	var result []interval

	for len(accumulated) > 1 {
		if accumulated[0].overlap(accumulated[1]) {
			accumulated[1] = accumulated[0].combine(accumulated[1])
		} else {
			result = append(result, accumulated[0])
		}
		accumulated = accumulated[1:]
	}

	result = append(result, accumulated...)

	// fmt.Printf("  result: %v\n", result)

	return result
}

func (d def) mapNum(i int) int {
	for _, sm := range d.maps {
		i = sm.mapNum(i)
	}
	return i
}

func (d def) mapIntervals(intervals []interval) []interval {
	// fmt.Printf("%d intervals\n", len(intervals))
	for _, sm := range d.maps {
		intervals = sm.mapIntervals(intervals)
		// fmt.Printf("%d intervals\n", len(intervals))
	}
	return intervals
}

func (d def) seedIntervals() []interval {
	var result []interval
	for i := 0; i < len(d.seeds); i += 2 {
		start := d.seeds[i]
		length := d.seeds[i+1]
		result = append(result, interval{
			start: start,
			end:   start + length - 1,
		})
	}
	return result
}

func parseMap(inputs []string) (spanMap, error) {
	for inputs[len(inputs)-1] == "" {
		inputs = inputs[:len(inputs)-1]
	}

	var result spanMap

	if !strings.HasSuffix(inputs[0], " map:") {
		return result, fmt.Errorf("weird map name input line: %q", inputs[0])
	}
	result.name = inputs[0][:len(inputs[0])-5]
	// fmt.Println(result.name)

	grid, err := util.ParseGrid(inputs[1:])
	if err != nil {
		return result, err
	}

	for _, mapping := range grid {
		dest, source, length := mapping[0], mapping[1], mapping[2]
		s := span{
			start:  source,
			end:    source + length - 1,
			offset: dest - source,
		}
		result.spans = append(result.spans, s)
	}

	sort.Slice(result.spans, func(i, j int) bool { return result.spans[i].start < result.spans[j].end })
	// fmt.Println(result)

	return result, nil
}

func parseInput(inputs []string) (def, error) {
	var result def
	var err error

	pieces := util.SplitBefore(inputs, func(s string) bool { return strings.Contains(s, ":") })

	if !strings.HasPrefix(pieces[0][0], "seeds: ") {
		return result, fmt.Errorf("unexpected seed list input line: %q", pieces[0][0])
	}

	result.seeds, err = util.ParseInts(pieces[0][0][7:], " ")
	if err != nil {
		return result, err
	}

	for _, piece := range pieces[1:] {
		m, err := parseMap(piece)
		if err != nil {
			return result, err
		}
		result.maps = append(result.maps, m)
	}

	return result, nil
}

func part1(inputs []string) (int, error) {
	d, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}
	min := math.MaxInt

	for _, seed := range d.seeds {
		mapped := d.mapNum(seed)
		if mapped < min {
			min = mapped
		}
	}

	return min, nil
}

func part2(inputs []string) (int, error) {
	d, err := parseInput(inputs)
	if err != nil {
		return 0, err
	}

	intervals := d.mapIntervals(d.seedIntervals())

	return intervals[0].start, nil
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
