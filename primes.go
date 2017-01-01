package primes

import (
	"math"
)

var (
	Primes []int
	max    = 400000
)

func InitPrimes(m int) {
	Primes = make([]int, 8)
	Primes[0] = 2
	Primes[1] = 3
	Primes[2] = 5
	Primes[3] = 7
	Primes[4] = 11
	Primes[5] = 13
	Primes[6] = 17
	Primes[7] = 19
	if m < 20 {
		max = 20
	} else {
		max = m
	}
	// calculate all primes up to max
	for i := 23; i <= max; i++ {
		IsPrime(i)
	}
}

func IsPrime(p int) bool {
	if p <= 1 {
		return false
	}
	if p > max {
		return false
	}

	f := float64(p)
	if p <= Primes[len(Primes)-1] {

		// this is the magic line
		// a is the indice very close to p
		a := int(f/math.Log(f)) - 2
		b := (a + 1) * 2
		if b > len(Primes)-1 {
			b = len(Primes) - 1
		}

		// binary search
		for a < b {
			k := a + (b-a)/2
			if p == Primes[a] {
				return true
			}
			if p == Primes[b] {
				return true
			}
			if p == Primes[k] {
				return true
			}
			if p > Primes[k] {
				a = k + 1
			}
			if p < Primes[k] {
				b = k - 1
			}
		}
		return false
	}

	for i := 0; i < len(Primes); i++ {
		if p%Primes[i] == 0 {
			return false
		}
		if int(math.Sqrt(f)) < Primes[i] {
			break // it is prime
		}
	}

	Primes = append(Primes, p)
	return true
}
