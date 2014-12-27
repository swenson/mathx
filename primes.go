package mathx

import (
	"math"
)

var primes = []int64{2, 3}

// generate primes up to and including n by Sieve of Eratosthenes
func genPrimes(n int64) {
	if primes[len(primes)-1] >= n {
		return
	}

	size := int((n / 2 / 64) + 1)
	mem := make([]uint64, size, size)
	primes = make([]int64, 0, int(math.Ceil(float64(n)/math.Log(float64(n)))))
	primes = append(primes, 2)
	maxn := int(n-1)/2 + 1

	for i := 1; i < maxn; i++ {
		loc := i >> 6
		bit := uint(i & 63)

		if mem[loc]&(uint64(1)<<bit) == 0 {
			p := 2*i + 1
			primes = append(primes, int64(p))
			for j := i + p; j < maxn; j += p {
				loc = j >> 6
				bit = uint(j & 63)
				mem[loc] |= uint64(1) << bit
			}
		}
	}
}

// Factorization64 returns the complete factorization of the given 64-bit
// argument using trial division.
func Factorization64(n int64) []factor64 {
	sqrtN := int64(math.Floor(math.Sqrt(float64(n))))
	genPrimes(sqrtN)

	factors := []factor64{}
	if n < 0 {
		factors = append(factors, factor64{-1, 1})
		n = -n
	}
	if n <= 1 {
		return factors
	}
	for _, p := range primes {
		x := 0
		for ; n%p == 0; x++ {
			n = n / p
		}
		if x > 0 {
			factors = append(factors, factor64{p, x})
		}
	}
	return factors
}

// IsSquareFree64 returns true if the 64-bit argument is square free.
func IsSquareFree64(n int64) bool {
	factors := Factorization64(n)
	for _, f := range factors {
		if f.exponent >= 2 {
			return false
		}
	}
	return true
}

// PosMod returns a % b, but always positive.
func PosMod(a, b int64) int64 {
	m := a % b
	if m < 0 {
		return m + b
	}
	return m
}
