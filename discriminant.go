package ntag

import (
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

func Factorization64(n int64) []factor64 {
  factors := []factor64{}
  if n < 0 {
    factors = append(factors, factor64{-1, 1})
  }
  if n <= 1 {
    return factors
  }
  twos := 0
  for twos := 0; n % 2 == 0; twos++ {
    n = n >> 1
  }
  if twos > 0 {
    factors = append(factors, factor64{2, twos})
  }
  for p := int64(3); n > 1; p += 2 {
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
