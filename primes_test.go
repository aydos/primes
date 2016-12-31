package primes

import (
	"math/big"
	"math/rand"
	"testing"
)

func init() {
	InitPrimes(max)
}

func Test(t *testing.T) {
	// Test method 1
	var tests = []struct {
		i int
		b bool
	}{
		{0, false},
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{11, true},
		{13, true},
		{17, true},
		{19, true},
		{101, true},
		{103, true},
		{107, true},
		{109, true},
		{1973, true},
		{2017, true},
		{2018, false},
	}
	for _, c := range tests {
		if got := Primes(c.i); got != c.b {
			t.Errorf("Prime calculation error: %d", c.i)
		}
	}

	// Test method 2
	for i := 0; i < max; i++ {
		p := Primes(i)
		b := big.NewInt(int64(i))
		if p != b.ProbablyPrime(20) {
			t.Errorf("Probably Prime calculation error: %v", i)
		}
	}

}

func BenchmarkPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := rand.Intn(max)
		_ = Primes(k)
	}
}

func BenchmarkProbablyPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := rand.Intn(max)
		b := big.NewInt(int64(k))
		_ = b.ProbablyPrime(20)
	}
}
