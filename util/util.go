package util

import "fmt"

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
