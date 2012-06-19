package ntag

import (
  "math"
  "math/big"
 )

type factor64 struct {
  prime int64
  exponent int
}

func (k *NumberField) Discriminant() *big.Int {
  return k.polynomial.Discriminant()
}

func (p *IntPolynomial) Discriminant() *big.Int {
  if p.Degree() == 2 {
    c, b, a := p.coeffs[0], p.coeffs[1], p.coeffs[2]
    b2 := big.NewInt(0).Mul(&b, &b)
    ac4 := big.NewInt(4)
    ac4 = ac4.Mul(ac4, &a)
    ac4 = ac4.Mul(ac4, &c)
    return b2.Sub(b2, ac4)
  }
  return nil
}

var primes = []int64{2, 3}

func genPrimes(n int64) {
  if primes[len(primes) - 1] >= n {
    return
  }
  // Use the primes we already generated.
  for q := primes[len(primes) - 1] + 2; q <= n; q += 2 {
    prime := true
    for _, p := range primes {
      if q % p == 0 {
        prime = false
        break
      }
    }
    if prime {
      primes = append(primes, q)
    }
  }
}

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
    for ; n % p == 0; x++ {
      n = n / p
    }
    if x > 0 {
      factors = append(factors, factor64{p, x})
    }
  }
  return factors
}

func IsSquareFree64(n int64) bool {
  factors := Factorization64(n)
  for _, f := range factors {
    if f.exponent >= 2 {
      return false
    }
  }
  return true
}

// Always return the positive modulus.
func PosMod(a, b int64) int64 {
  m := a % b
  if m < 0 {
    return m + b
  }
  return m
}

func IsFundamentalDiscriminant(D *big.Int) bool {
  d := D.Int64()
  absd := d
  if d < 0 {
    absd = -absd
  }
  if d == 1 {
    return true
  }
  if PosMod(d, 4) == 1 && IsSquareFree64(absd) {
    return true
  }
  if PosMod(d, 4) == 0 {
    if IsSquareFree64(absd/4) {
      return PosMod(d/4, 4) == 2 || PosMod(d/4, 4) == 3
    }
  }
  return false
}

func makeFundamentalDiscriminant(D int64) int64 {
  if IsFundamentalDiscriminant(big.NewInt(D)) {
    return D
  }
  factors := Factorization64(D)
  for _, f := range factors {
    ex := f.exponent
    for ex >= 2 {
      D = D / (f.prime * f.prime)
      ex -= 2
    }
  }
  return D
}
