package math

import (
	"fmt"
	"math"
	"sort"
)

// ModMul return a * b, mod m.
func ModMul(a, b, m int) int {
	result := 0
	a = a % m
	b = b % m
	if b < 0 {
		a, b = -a, -b
	}
	for b > 0 {
		if b&1 == 1 {
			result = (result + a) % m
		}
		a = (a * 2) % m
		b = b / 2
	}
	return result
}

// ModExp returns base ** exponent, mod m.
func ModExp(base, exponent, m int) int {
	// debug := (base == 17463217478) && (exponent == 5)
	debug := false
	if debug {
		fmt.Println("HERE!")
		fmt.Printf("m=%d\n", m)
	}
	prod := 1
	base = base % m
	for exponent > 0 {
		if base == 0 {
			return 0
		}
		if base == 1 {
			return prod
		}

		if exponent&1 > 0 {
			if debug {
				fmt.Printf("BEFORE: prod=%d, base=%d\n", prod, base)
			}
			prod = ModMul(prod, base, m)
			if debug {
				fmt.Printf("AFTER: prod=%d, base=%d\n", prod, base)
			}
		}
		exponent >>= 1
		if debug {
			fmt.Printf("BEFORE: base=%d\n", base)
		}
		base = ModMul(base, base, m)
		if debug {
			fmt.Printf("AFTER: base=%d\n", base)
		}
	}
	return prod
}

// ModInv calculates the modular multiplicative inverse of x, modulo m.
// If gcd(x,m) != 1, it returns an error.
func ModInv(x, m int) (int, error) {
	if m <= 2 {
		return 0, fmt.Errorf("ModInv expects m >= 3; got %d", m)
	}
	if x < 1 {
		return 0, fmt.Errorf("ModInv expects x >= 1; got %d", x)
	}

	if x >= m {
		return 0, fmt.Errorf("ModInv expects |x| < n; got %d >= %d", x, m)
	}

	rim1 := m
	ri := x
	tim1 := 0
	ti := 1

	for {
		// fmt.Printf("rim1=%d, ri=%d\n", rim1, ri)
		// fmt.Printf("tim1=%d, ti=%d\n", tim1, ti)
		q := rim1 / ri
		rip1 := rim1 % ri
		// fmt.Printf("q=%d, rip1=%d\n", q, rip1)
		if rip1 == 0 {
			break
		}
		tip1 := (tim1 - ModMul(q, ti, m)) % m

		rim1, ri = ri, rip1
		tim1, ti = ti, tip1
	}

	gcd := ri
	if gcd != 1 {
		return 0, fmt.Errorf("ModInv(%d,%d): need gcd==1; got %d", x, m, gcd)
	}

	if ti < 0 {
		ti = ti%m + m
	}

	return ti, nil
}

// ModGeometricSum returns the sum of the first n terms of the sequence rⁿ, starting with r⁰, modulo m.
// eg: ModGeometricSum(37, 5, 10000001) = (1 + 37 + 37² + 37³ + 37⁴) mod 10000001 = 1926221
// eg: ModGeometricSum(37, 5, 101) = (1 + 37 + 37² + 37³ + 37⁴) mod 101 = 50
func ModGeometricSum(r, n, m int) (int, error) {
	if r == 0 {
		return 0, nil
	}
	if r == 1 {
		return ModMul(r, n, m), nil
	}

	// 1-rⁿ / 1-r

	inv, err := ModInv(m+1-r, m)
	if err != nil {
		return 0, err
	}

	top := m + 1 - ModExp(r, n, m)

	return ModMul(top, inv, m), nil
}

// Sort3 sorts three ints in ascending order.
func Sort3(a, b, c int) (int, int, int) {
	if a < b {
		if b < c {
			return a, b, c
		}
		if a < c {
			return a, c, b
		}
		return c, a, b
	}

	if a < c {
		return b, a, c
	}
	if b < c {
		return b, c, a
	}
	return c, b, a
}

// MaxUint is the largest uint value.
const MaxUint = ^uint(0)

// MaxInt is the max int value.
const MaxInt = int(MaxUint >> 1)

// MinInt is the smallest (negative) int value.
const MinInt = -MaxInt - 1

// ChooseNUint32 returns distinct groups of n uint32s from the inputs.
// It assumes the inputs are distinct too.
func ChooseNUint32(ints []uint32, n int) [][]uint32 {
	if n == 0 {
		return [][]uint32{{}}
	}
	if len(ints) < n {
		return nil
	}
	if len(ints) == n {
		return [][]uint32{ints}
	}
	without := ChooseNUint32(ints[1:], n)
	with := ChooseNUint32(ints[1:], n-1)
	result := make([][]uint32, len(without)+len(with))
	copy(result[len(with):], without)
	for i, w := range with {
		result[i] = append([]uint32{ints[0]}, w...)
	}
	return result
}

// GCD uses the Euclidean algorithm to calculate the GCD of a and b.
func GCD(a, b int) int {
	if a > b {
		a, b = b, a
	}
	if a == 0 {
		return b
	}
	if a == 1 {
		return 1
	}
	m := b % a
	if m == 0 {
		return a
	}
	return GCD(m, a)
}

// LCM computes the least common multiple of a and b. It doesn't handle negative numbers.
// If either a or b is zero, it returns zero.
func LCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	return a * (b / GCD(a, b))
}

type Factor struct {
	Prime int
	Count int
}

var primes = []int{2, 3, 5, 7, 11, 13}

func Sieve(n int) {
	if primes[len(primes)-1] > n {
		return
	}
}

func SlowPrimesLessThanOrEqualTo(target int) []int {
OUTER:
	for n := primes[len(primes)-1] + 2; n <= target; n += 2 {
		for _, p := range primes[1:] {
			if n%p == 0 {
				continue OUTER
			}
		}
		primes = append(primes, n)
	}

	for i, p := range primes {
		if p > target {
			return primes[:i]
		}
	}
	return primes
}

func PrimeFactors(n int) []Factor {
	if n < 0 {
		n = -n
	}
	if n < 2 {
		return nil
	}
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))
	pp := SlowPrimesLessThanOrEqualTo(sqrt)

	factorCount := make(map[int]int)

	largest := 0
	for i := 0; i < len(pp) && n > 1; i++ {
		prime := pp[i]
		for n%prime == 0 {
			factorCount[prime]++
			n /= prime
			largest = i
		}
	}

	res := make([]Factor, 0, largest+1)

	for i, prime := range pp {
		if i > largest {
			break
		}
		if count := factorCount[prime]; count > 0 {
			res = append(res, Factor{Prime: prime, Count: count})
		}
	}

	if n > 1 {
		res = append(res, Factor{Prime: n, Count: 1})
	}

	return res
}

// MultiLCM computes the LCM of multiple numbers.
func MultiLCM(nums ...int) int {
	maxFactors := make(map[int]int)

	for _, num := range nums {
		for _, factor := range PrimeFactors(num) {
			if maxFactors[factor.Prime] < factor.Count {
				maxFactors[factor.Prime] = factor.Count
			}
		}
	}
	product := 1
	for factor, count := range maxFactors {
		for i := 0; i < count; i++ {
			product *= factor
		}
	}

	return product
}

// MultiGCD computes the GCD of multiple numbers.
func MultiGCD(nums ...int) int {
	noZeros := make([]int, len(nums))
	for _, num := range nums {
		if num != 0 {
			if num < 0 {
				num = -num
			}
			if num == 1 {
				return num
			}
			noZeros = append(noZeros, num)
		}
	}

	sort.Ints(noZeros)
	switch len(noZeros) {
	case 0:
		return 1
	case 1:
		return noZeros[0]
	case 2:
		return GCD(noZeros[0], noZeros[1])
	}

	var minFactors map[int]int

	first := true
	for _, num := range noZeros {
		theseFactors := make(map[int]int)
		for _, factor := range PrimeFactors(num) {
			theseFactors[factor.Prime] = factor.Count
		}
		if first {
			first = false
			minFactors = theseFactors
			continue
		}

		for factor, count := range minFactors {
			thisCount, found := theseFactors[factor]
			if !found {
				delete(minFactors, factor)
				if len(minFactors) == 0 {
					return 1
				}
			} else {
				if thisCount < count {
					minFactors[factor] = thisCount
				}
			}
		}
	}
	product := 1
	for factor, count := range minFactors {
		for i := 0; i < count; i++ {
			product *= factor
		}
	}

	return product
}
