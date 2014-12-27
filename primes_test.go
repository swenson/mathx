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
