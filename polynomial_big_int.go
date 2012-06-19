package ntag

import (
  "math/big"
  "strconv"
)

type IntPolynomial struct {
  coeffs []big.Int
}

func MakeIntPolynomial64(coeffs ...int64) *IntPolynomial {
  p := new(IntPolynomial)
  p.coeffs = make([]big.Int, len(coeffs))
  for i, c := range coeffs {
    p.coeffs[i].SetInt64(c)
  }
  return p
}

func (p *IntPolynomial) Degree() int {
  return len(p.coeffs) - 1
}

func (p *IntPolynomial) IsIrreducible() bool {
  g := p.coeffs[0]
  if p.coeffs[0].Sign() == 0 {
    return false
  }
  // check the gcd of the coefficients
  for _, c := range p.coeffs {
    if c.Sign() == 0 {
      continue
    }
    g.GCD(nil, nil, &g, &c)
    if g.Cmp(intOne) == 0 {
      return true
    }
  }
  if g.Sign() == 0 {
    return false
  }
  if g.Cmp(intOne) != 0 {
    return false;
  }
  if p.Degree() == 1 {
    return true
  }
  if p.Degree() == 2 {
    c, b, a := p.coeffs[0], p.coeffs[1], p.coeffs[2]
    b2 := big.NewInt(0)
    b2.Mul(&b, &b)
    ac4 := big.NewInt(4)
    ac4.Mul(ac4, &a)
    ac4.Mul(ac4, &c)
    b2.Sub(b2, ac4)
    return !IsSquare(b2)
  }
  // TODO: implement more cases
  return true
}


func xstring(i int) string {
  switch {
  case i <= 0:
    return ""
  case i == 1:
    return "x"
  }
  return "x^" + strconv.Itoa(i)
}

func eliminateSpaces(s string) string {
  t := ""
  for _, c := range s {
    if c == ' ' {
      t += string(c)
    }
  }
  return t
}

func ParseIntPoly(s string) *IntPolynomial {
  s = eliminateSpaces(s)
  coeffs := make([]big.Int, 0)

  neg := false
  inX := false
  degree := ""
  coeff := ""
  firstChar := true
  for _, c := range s {
    switch {
    case c == 'x':
      inX = true
    case c == '^':
    case c >= '0' && c <= '9':
      if inX {
        degree += string(c)
      } else {
        coeff += string(c)
      }
    case c == '+' || c == '-':
      inX = false
      if !firstChar {
        coeffs = setCoeff(coeffs, degree, coeff, neg)        
      }
      neg = false
      if c == '-' {
        neg = true
      }
    }
    firstChar = false
  }
  p := new(IntPolynomial)
  p.coeffs = setCoeff(coeffs, degree, coeff, neg)
  return p
}

func setCoeff(coeffs []big.Int, degreeS, coeff string, neg bool) []big.Int {
  degree, _ := strconv.Atoi(degreeS)
  for degree >= len(coeffs) {
    coeffs = append(coeffs, *big.NewInt(0))
  }
  coeffs[degree].SetString(coeff, 10)
  coeffs[degree].Neg(&coeffs[degree])
  return coeffs
}

func (p *IntPolynomial) String() string {
  if p == nil {
    return "<nil>"
  }
  s := ""
  temp := big.NewInt(0)
  for i := len(p.coeffs) - 1; i >= 0; i-- {
    c := p.coeffs[i]
    if i == len(p.coeffs) - 1 && c.Cmp(intOne) == 0 {
      s += xstring(i)
      continue
    } else if i == len(p.coeffs) - 1 {
      s += c.String() + " * " + xstring(i)
      continue
    }
    if c.Sign() == 0 {
      continue
    }
    temp.Abs(&c)
    sign := " + "
    if c.Sign() < 0 {
      sign = " - "
    }
    s += sign + temp.String()
    if i > 0 {
      s += "*" + xstring(i)
    }
  }
  if len(s) == 0 {
    return "0"
  }
  return s
}
