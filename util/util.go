package util

import (
	"fmt"
	"strings"
)

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func LowestTrue(lowFalseStart int, pred func(int) (bool, error)) (int, error) {
	lf, err := pred(lowFalseStart)
	if err != nil {
		return 0, err
	}
	if lf {
		return 0, fmt.Errorf("lowestTrue expected pred(lowFalseStart)==false; got pred(%d)==true", lowFalseStart)
	}
	lowFalse := lowFalseStart
	highTrue := 0
	for lowFalse < MaxInt/2 {
		attempt := lowFalse * 2
		st, err := pred(attempt)
		if err != nil {
			return 0, err
		}
		if st {
			highTrue = attempt
			break
		}
		lowFalse <<= 1
	}
	if highTrue == 0 {
		return 0, fmt.Errorf("cannot find high enough value to make pred(value)==true")
	}

	for highTrue-lowFalse > 1 {
		mid := (lowFalse + highTrue) / 2
		mm, err := pred(mid)
		if err != nil {
			return 0, err
		}
		if mm {
			highTrue = mid
		} else {
			lowFalse = mid
		}
	}
	return highTrue, nil
}

func highestTrueRange(lowTrue int, highFalse int, pred func(int) (bool, error)) (int, error) {
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

// TrimmedLines takes a string, splits it into lines, and trims each line of starting
// and ending whitespace.
func TrimmedLines(s string) []string {
	result := strings.Split(strings.TrimSpace(s), "\n")
	for i, r := range result {
		result[i] = strings.TrimSpace(r)
	}
	return result
}

// GroupString returns the input string, broken into runs of consecutive characters
func GroupString(s string) []string {
	var result []string
	last := -1
	for i := range s {
		c := s[i : i+1]
		if len(result) == 0 || result[last][:1] != c {
			result = append(result, c)
			last++
		} else {
			result[last] = result[last] + c
		}
	}
	return result
}
