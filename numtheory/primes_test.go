package numtheory

import (
	"testing"
)

func TestMod60(t *testing.T) {
	for n := 0; n < 10000000; n++ {
		if fastmod60(int64(n)) != (n % 60) {
			t.Fatalf("fastmod60(%d) = %d, but should be %d", n, fastmod60(int64(n)), n%60)
		}
	}
}

func TestMod12(t *testing.T) {
	for n := 0; n < 10000000; n++ {
		if fastmod12(int64(n)) != (n % 12) {
			t.Fatalf("fastmod12(%d) = %d, but should be %d", n, fastmod12(int64(n)), n%12)
		}
	}
}

func TestMod5(t *testing.T) {
	for n := 0; n < 10000000; n++ {
		if fastmod5(int64(n)) != (n % 5) {
			t.Fatalf("fastmod5(%d) = %d, but should be %d", n, fastmod5(int64(n)), n%5)
		}
	}
}

func TestPrimeGeneration(t *testing.T) {
	primes = []int64{2, 3}
	genPrimes(200)
	expected := []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}
	if len(primes) != len(expected) {
		t.Fatalf("Could not generate primes up to 200: got %v", primes)
	}
	for i, p := range primes {
		if expected[i] != p {
			t.Fatalf("Could not generate primes up to 200: got %v", primes)
		}
	}

	primes = []int64{2, 3}
	genPrimesAtkin(200)
	if len(primes) != len(expected) {
		t.Fatalf("Could not generate primes up to 200: got %v", primes)
	}
	for i, p := range primes {
		if expected[i] != p {
			t.Fatalf("Could not generate primes up to 200: got %v", primes)
		}
	}
}

func BenchmarkPrimeSieve100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes = []int64{2, 3}
		genPrimes(100)
	}
}

func BenchmarkPrimeSieve10000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes = []int64{2, 3}
		genPrimes(10000000)
	}
}

func BenchmarkPrimeAtkinSieve100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes = []int64{2, 3}
		genPrimesAtkin(100)
	}
}

func BenchmarkPrimeAtkinSieve10000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes = []int64{2, 3}
		genPrimesAtkin(10000000)
	}
}

func BenchmarkPrimeDiv100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes = []int64{2, 3}
		slowGenPrimes(100)
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
