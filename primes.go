package primes

import (
	"math"
)

var (
	primes []int
	max    = 400000
	last   = 0
)

func InitPrimes(m int) {
	primes = make([]int, 8)
	primes[0] = 2
	primes[1] = 3
	primes[2] = 5
	primes[3] = 7
	primes[4] = 11
	primes[5] = 13
	primes[6] = 17
	primes[7] = 19
	if m < 20 {
		max = 20
	} else {
		max = m
	}
	for i := 23; i < max; i++ {
		Primes(i)
	}
	last = primes[len(primes)-1]
}

func Primes(p int) bool {
	if p <= 1 {
		return false
	}
	if p > max {
		return false
	}

	f := float64(p)
	if p <= last {

		n := int(math.Floor(f/math.Log(f)-1)) - 1

		for i := n; i < len(primes); i++ {
			if p == primes[i] {
				return true
			}
			if p < primes[i] {
				return false
			}
		}
		return false
	}

	for i := 0; i < len(primes); i++ {
		if p%primes[i] == 0 {
			return false
		}
		if int(math.Sqrt(f)) < primes[i] {
			break
		}
	}

	primes = append(primes, p)
	return true
}
