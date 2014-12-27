package mathx

import (
	"testing"
)

func TestPrimeGeneration(t *testing.T) {
	primes = []int64{2, 3}
	genPrimes(30)
	expected := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	if len(primes) != len(expected) {
		t.Fatalf("Could not generate primes up to 30: got %v", primes)
	}
	for i, p := range primes {
		if expected[i] != p {
			t.Fatalf("Could not generate primes up to 30: got %v", primes)
		}
	}
}

func TestFacorization64(t *testing.T) {
	num := int64(154297)
	expected := []factor64{{11, 1}, {13, 2}, {83, 1}}
	got := Factorization64(num)
	if len(got) != len(expected) {
		t.Fatal("Could not factor 154297")
	}
	for i, f := range expected {
		if f != got[i] {
			t.Fatal("Could not factor 15497")
		}
	}
}

func slowGenPrimes(n int64) {
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
