package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/zellyn/adventofcode/util"
	"golang.org/x/exp/maps"
)

var printf = func(string, ...any) {}

// var printf = fmt.Printf

var indentf = func(indent int, format string, rest ...any) {
	printf(strings.Repeat("  ", indent)+format, rest...)
}

var stringToField = map[string]int{
	"x": 0,
	"m": 1,
	"a": 2,
	"s": 3,
}

type part [4]int

func (p part) sum() int {
	return p[0] + p[1] + p[2] + p[3]
}

type condition struct {
	field  int
	op     string
	num    int
	target string
}

func (c condition) meet(p part) bool {
	num := p[c.field]
	switch c.op {
	case "<":
		return num < c.num
	case ">":
		return num > c.num
	default:
		panic("unexpected operator: " + c.op)
	}
}

func (c condition) inside(i interval) (interval, bool) {
	minMatch := c.meet(i.min)
	maxMatch := c.meet(i.max)
	if !minMatch && !maxMatch {
		return i, false
	}
	if minMatch && maxMatch {
		return i, true
	}
	if minMatch {
		if c.op == ">" {
			panic("min-only match on >")
		}

		i.max[c.field] = c.num - 1
		return i, true
	}
	if c.op == "<" {
		panic("max-only match on <")
	}

	i.min[c.field] = c.num + 1

	return i, true
}

func (c condition) outside(i interval) (interval, bool) {
	minMatch := c.meet(i.min)
	maxMatch := c.meet(i.max)
	if minMatch && maxMatch {
		return i, false
	}
	if !minMatch && !maxMatch {
		return i, true
	}
	if minMatch {
		if c.op == ">" {
			panic("min-only match on >")
		}
		i.min[c.field] = c.num

		return i, true
	}
	if c.op == "<" {
		panic("max-only match on <")
	}

	i.max[c.field] = c.num
	return i, true
}

type workflow struct {
	name       string
	conditions []condition
	recourse   string
}

func (w workflow) outcome(p part) string {
	for _, c := range w.conditions {
		if c.meet(p) {
			return c.target
		}
	}
	return w.recourse
}

func parseParts(inputs []string) ([]part, error) {
	return util.MapE(inputs, func(s string) (part, error) {
		var p part
		if s[0] != '{' || s[len(s)-1] != '}' {
			return p, fmt.Errorf("weird part: %q", s)
		}
		pieces := strings.Split(s[1:len(s)-1], ",")
		for _, piece := range pieces {
			name, val, ok := strings.Cut(piece, "=")
			if !ok {
				return p, fmt.Errorf("weird entry (%q) in part %q", piece, s)
			}
			i, err := strconv.Atoi(val)
			if err != nil {
				return p, fmt.Errorf("weird entry (%q) in part %q: %v", piece, s, err)
			}
			field := stringToField[name]
			p[field] = i
			switch name {
			case "x":
				p[0] = i
			case "m":
				p[1] = i
			case "a":
				p[2] = i
			case "s":
				p[3] = i
			default:
				return p, fmt.Errorf("weirdly named variable in entry %q in part %q: %v", piece, s, err)
			}
		}
		return p, nil
	})
}

var conditionRe = regexp.MustCompile(`([xmas])([<>])([0-9]+):([a-zA-Z]+)`)

func parseFlows(inputs []string) (map[string]workflow, error) {
	flows, err := util.MapE(inputs, func(s string) (workflow, error) {
		var w workflow
		if s[len(s)-1] != '}' {
			return w, fmt.Errorf("Weird workflow: should end in '}': %q", s)
		}
		name, seq, ok := strings.Cut(s[:len(s)-1], "{")
		if !ok {
			return w, fmt.Errorf("Weird workflow: sequence should start with '{': %q", s)
		}
		w.name = name
		conditions := strings.Split(seq, ",")
		w.recourse = conditions[len(conditions)-1]
		var err error
		w.conditions, err = util.MapE(conditions[:len(conditions)-1], func(cs string) (condition, error) {
			var c condition
			match := conditionRe.FindStringSubmatch(cs)
			if match == nil {
				return c, fmt.Errorf("Weird condition (%q) in input %q", cs, s)
			}
			c.field = stringToField[match[1]]
			c.op = match[2]
			i, err := strconv.Atoi(match[3])
			c.num = i
			c.target = match[4]

			return c, err
		})
		return w, err
	})
	if err != nil {
		return nil, err
	}
	res := make(map[string]workflow, len(flows))
	for _, flow := range flows {
		res[flow.name] = flow
	}
	return res, nil
}

func parse(inputs []string) (map[string]workflow, []part, error) {
	paras := util.LinesByParagraph(inputs)
	if len(paras) != 2 {
		return nil, nil, fmt.Errorf("expected 2 paragraphs; got %d", len(paras))
	}

	flows, err := parseFlows(paras[0])
	if err != nil {
		return nil, nil, err
	}

	parts, err := parseParts(paras[1])

	return flows, parts, err
}

func simplifyFlows(flows map[string]workflow) {
	allSame := make(map[string]string)

	done := false
	for !done {
		done = true
		for name, flow := range flows {

			// Replace names with equivalents
			if altRecourse := allSame[flow.recourse]; altRecourse != "" {
				flow.recourse = altRecourse
				done = false
			}
			for i := range flow.conditions {
				if altTarget := allSame[flow.conditions[i].target]; altTarget != "" {
					flow.conditions[i].target = altTarget
					done = false
				}
			}

			// Collapse unnecessary rules
			for i := len(flow.conditions) - 1; i >= 0; i-- {
				cond := flow.conditions[i]
				if cond.target == flow.recourse {
					flow.conditions = flow.conditions[:i]
					done = false
				} else {
					break
				}
			}
			if len(flow.conditions) == 0 {
				allSame[name] = flow.recourse
			}

			flows[name] = flow
		}
	}
}

const LIMIT = 10_000

func finalOutcome(p part, workflows map[string]workflow) string {
	name := "in"
	for i := 0; i < LIMIT; i++ {
		if name == "A" || name == "R" {
			return name
		}
		flow, ok := workflows[name]
		if !ok {
			panic("unknown workflow: " + name)
		}
		name = flow.outcome(p)
	}
	panic(fmt.Sprintf("part %v took > %d steps", p, LIMIT))
}

func part1(inputs []string) (int, error) {
	flows, parts, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	simplifyFlows(flows)

	sum := 0
	for _, p := range parts {
		outcome := finalOutcome(p, flows)
		if outcome == "A" {
			sum += p.sum()
		}
	}

	return sum, nil
}

type interval struct {
	min part
	max part
}

func (i interval) overlap(field, start, end int) bool {
	// No overlap at all for this field.
	if i.max[field] < start || i.min[field] > end {
		return false
	}
	return true
}

func (i interval) count() int {
	prod := 1
	for field := 0; field < 4; field++ {
		prod *= (i.max[field] - i.min[field] + 1)
	}
	return prod
}

func (i interval) intersect(j interval) (interval, bool) {
	var res interval

	for field := 0; field < 4; field++ {
		iMinVal := i.min[field]
		iMaxVal := i.max[field]
		jMinVal := j.min[field]
		jMaxVal := j.max[field]

		// No overlap at all for this field.
		if iMaxVal < jMinVal || iMinVal > jMaxVal {
			return interval{}, false
		}

		res.min[field] = max(iMinVal, jMinVal)
		res.max[field] = min(iMaxVal, jMaxVal)
	}

	return res, true
}

func allIntervals(name string, i interval, flows map[string]workflow) []interval {
	var res []interval

LOOP:
	flow := flows[name]

	for _, c := range flow.conditions {
		minMeet := c.meet(i.min)
		maxMeet := c.meet(i.max)
		if !minMeet && !maxMeet {
			continue
		}
		in, haveIn := c.inside(i)
		out, haveOut := c.outside(i)

		if haveIn {
			switch c.target {
			case "A":
				res = append(res, in)
			case "R":
				// rejected!
			default:
				res = append(res, allIntervals(c.target, in, flows)...)
			}
		}

		if !haveOut {
			return res
		}
		i = out
	}

	switch flow.recourse {
	case "A":
		return append(res, i)
	case "R":
		return res
	default:
		name = flow.recourse
		goto LOOP
	}
}

func findBreaks(intervals []interval) [4][]int {
	var res [4][]int

	for field := 0; field < 4; field++ {
		seen := make(map[int]bool, len(intervals)*2)
		for _, i := range intervals {
			seen[i.min[field]] = true
			seen[i.max[field]+1] = true
		}
		breaks := maps.Keys(seen)
		sort.Ints(breaks)
		res[field] = breaks
	}

	return res
}

func overlapping(field, start, end int, intervals []interval) []interval {
	res := make([]interval, 0, len(intervals))
	for _, i := range intervals {
		if i.overlap(field, start, end) {
			res = append(res, i)
		}
	}
	return res
}

func count(field int, intervals []interval, breaks [4][]int) int {
	// indentf(field*3, "field %d:\n", field)
	if len(intervals) == 0 {
		return 0
	}

	sum := 0

	for i := 0; i < len(breaks[field])-1; i++ {
		start := breaks[field][i]
		end := breaks[field][i+1] - 1
		// indentf(field*3+1, "considering %d-%d\n", start, end)

		newIntervals := overlapping(field, start, end, intervals)
		// indentf(field*3+2, "matching intervals: %v\n", newIntervals)
		if field == 3 && len(newIntervals) > 0 {
			sum += end - start + 1
			continue
		}

		sum += (end - start + 1) * count(field+1, newIntervals, breaks)
	}
	return sum
}

func part2(inputs []string) (int, error) {
	flows, _, err := parse(inputs)
	if err != nil {
		return 0, err
	}
	simplifyFlows(flows)
	start := interval{
		min: part{1, 1, 1, 1},
		max: part{4000, 4000, 4000, 4000},
	}
	intervals := allIntervals("in", start, flows)
	for _, i := range intervals {
		printf("x:%d-%d m:%d-%d a:%d-%d s:%d-%d\n",
			i.min[0], i.max[0], i.min[1], i.max[1], i.min[2], i.max[2], i.min[3], i.max[3])
	}

	// All this was unnecessary, and could have been saved by thinking. sigh.
	/*
		breaks := findBreaks(intervals)
		for _, br := range breaks {
			printf("%v\n", br)
		}

		return count(0, intervals, breaks), nil
	*/

	return util.MappedSum(intervals, interval.count), nil
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
