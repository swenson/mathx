package mathx

import (
	"math"
)

var primes = []int64{2, 3}

func genPrimes(n int64) {
	if primes[len(primes)-1] >= n {
		return
	}

	// Use the primes we already generated.
	for q := primes[len(primes)-1] + 2; q <= n; q += 2 {
		prime := true
		for _, p := range primes {
			if q%p == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, q)
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
