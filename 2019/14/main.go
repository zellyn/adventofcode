package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type rule struct {
	output string
	count  int
	inputs map[string]int
}

type state struct {
	rules      map[string]rule
	order      []string
	needs      map[string]int
	depthCache map[string]int
}

// state can be sorted, which sorts the chemical order

func (s *state) Len() int           { return len(s.order) }
func (s *state) Swap(i, j int)      { s.order[i], s.order[j] = s.order[j], s.order[i] }
func (s *state) Less(i, j int) bool { return s.depth(s.order[i]) > s.depth(s.order[j]) }

func (s *state) depth(chem string) int {
	if chem == "ORE" {
		return 1
	}
	if d := s.depthCache[chem]; d > 0 {
		return d
	}
	d := 0
	for need := range s.rules[chem].inputs {
		dd := s.depth(need)
		if dd > d {
			d = dd
		}
	}
	d++
	s.depthCache[chem] = d
	return d
}

func parseInput(input string) (*state, error) {
	result := &state{
		rules:      make(map[string]rule),
		depthCache: make(map[string]int),
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		r, err := parseRule(line)
		if err != nil {
			return &state{}, err
		}
		result.rules[r.output] = r
		result.order = append(result.order, r.output)
	}
	sort.Sort(result)
	return result, nil
}

func (s *state) run(fuelCount int) (int, error) {
	s.needs = map[string]int{
		"FUEL": fuelCount,
	}
	if s.order[0] != "FUEL" {
		return 0, fmt.Errorf("want s.order[0]==%q; got %q", "FUEL", s.order[0])
	}
	for _, chem := range s.order {
		needCount := s.needs[chem]
		r := s.rules[chem]
		count := (needCount + r.count - 1) / r.count
		for name, c := range r.inputs {
			s.needs[name] += c * count
		}
	}
	return s.needs["ORE"], nil
}

func highestTrue(lowTrue int, highFalse int, pred func(int) (bool, error)) (int, error) {
	if lowTrue >= highFalse {
		return 0, fmt.Errorf("highestTrue(%d,%d, pred): want arg1 < arg2", lowTrue, highFalse)
	}
	lt, err := pred(lowTrue)
	if err != nil {
		return 0, err
	}
	if !lt {
		return 0, fmt.Errorf("highestTrue(%d,%d, pred): pred(%d)==false", lowTrue, highFalse, lowTrue)
	}
	hf, err := pred(highFalse)
	if err != nil {
		return 0, err
	}
	if hf {
		return 0, fmt.Errorf("highestTrue(%d,%d, pred): pred(%d)==true", lowTrue, highFalse, highFalse)
	}

	for highFalse-lowTrue > 1 {
		mid := (lowTrue + highFalse) / 2
		mm, err := pred(mid)
		if err != nil {
			return 0, err
		}
		if mm {
			lowTrue = mid
		} else {
			highFalse = mid
		}
	}
	return lowTrue, nil
}

func (s *state) calcMax(oreCount int) (int, error) {
	max := 2
	for {
		f, err := s.run(max)
		if err != nil {
			return 0, err
		}
		if f > oreCount {
			break
		}
		max = max * 2
	}

	result, err := highestTrue(max/2, max, func(i int) (bool, error) {
		f, err := s.run(i)
		if err != nil {
			return false, err
		}
		return f <= oreCount, nil
	})
	return result, err
}

// 1 ZDQRT, 3 CZLDF, 10 GDLFK, 1 BDRP, 10 VHMT, 6 XGVF, 1 RLFHL => 7 CVHR
func parseRule(line string) (r rule, err error) {
	err = fmt.Errorf("weird format rule: %q", line)
	line = strings.TrimSpace(line)
	halves := strings.Split(line, "=>")
	if len(halves) != 2 {
		return
	}
	chem, count, err := parseChemCount(halves[1])
	if err != nil {
		return rule{}, fmt.Errorf("problem in rule %q: %v", line, err)
	}
	r.count = count
	r.output = chem
	r.inputs = make(map[string]int)

	parts := strings.Split(halves[0], ",")
	for _, part := range parts {
		chem, count, err := parseChemCount(part)
		if err != nil {
			return rule{}, fmt.Errorf("problem in rule %q: %v", line, err)
		}
		r.inputs[chem] = count
	}

	return r, nil
}

func parseChemCount(s string) (string, int, error) {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("weird format chem count: %q", s)
	}
	c, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return "", 0, fmt.Errorf("count in chem count %q: %v", parts[0], err)
	}
	return strings.TrimSpace(parts[1]), c, nil
}
