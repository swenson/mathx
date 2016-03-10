package numtheory

import "math"

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

func genPrimesAtkin(max int64) {
	realMax := max
	max = (max/60 + 1) * 60
	results := []int64{2, 3, 5}
	list := make([]bool, (max + 1))
	n := int64(0)
	// 4x^2 + y^2 = n
	for x := int64(1); x <= int64(math.Ceil(math.Sqrt(float64(max-1)/4))); x++ {
		n := 4*x*x + 1
		for y := int64(1); ; y += 2 {
			if n > max {
				break
			}
			switch n % 60 {
			case 1, 13, 17, 29, 37, 41, 49, 53:
				list[n] = !list[n]
			}
			n += 2 + 2*(y+y+1)
		}
	}
	// 3x^2 + y^2 = n
	for x := int64(1); x <= int64(math.Ceil(math.Sqrt(float64(max-1)/3))); x += 2 {
		n := 3*x*x + 2*2
		for y := int64(2); n < max; y += 2 {
			if n > max {
				break
			}
			switch n % 60 {
			case 7, 19, 31, 43:
				list[n] = !list[n]
			}

			n += 2 + 2*(y+y+1)
		}
	}
	// 3x^2 - y^2 = n
	for x := int64(2); x <= int64(math.Ceil(math.Sqrt(float64(max)/3))); x++ {
		n := 3*x*x - (x-1)*(x-1)
		for y := x - 1; y >= 0; y -= 2 {
			if n > max {
				break
			}
			switch n % 60 {
			case 11, 23, 47, 59:
				list[n] = !list[n]
			}
			n += 4 * (y - 1)

		}
	}

	s := []int{1, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49, 53, 59}

	for w := int64(0); 3600*w < max; w++ {
		for _, x := range s {
			n := 60*w + int64(x)
			if n > max {
				break
			}
			if list[n] {
				c := int64(0)
				for v := int64(0); c < max; v++ {
					for _, y := range s {
						c = n * n * (60*v + int64(y))
						if c > max {
							break
						}
						list[c] = false
					}
				}
			}
		}
	}

	n = 0
	for w := int64(0); n < max; w++ {
		for _, x := range s {
			n = w*60 + int64(x)
			if n > realMax {
				break
			}
			if n > max {
				break
			}
			if list[n] {
				results = append(results, n)
			}
		}
	}

	primes = results
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
