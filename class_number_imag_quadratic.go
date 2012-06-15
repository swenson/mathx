package ntag

import (
	"math"
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

type NumberField struct {
	polynomial *IntPolynomial
}

func MakeNumberField(poly *IntPolynomial) *NumberField {
	k := new(NumberField)
	k.polynomial = poly
	return k
}

func (k *NumberField) Degree() int {
	return len(k.polynomial.coeffs) - 1
}

func (k *NumberField) ClassNumber() int {
	if k.Degree() == 2 {
		return classNumberImagQuad(k)
	}
	return -1
}

func classNumberImagQuad(k *NumberField) int {
	D := k.Discriminant().Int64()
	h := 1
	aD := D
	if aD < 0 {
		aD = -aD
	}
	B := int64(math.Floor(math.Sqrt(float64(aD) / 3.0)))

	for b := D % 2; b <= B; b += 2 {
		q := (b * b - D) / 4
		a := b
		if a <= 1 {
			a = 1
		}
		for a == 1 || a * a <= q {
			if a != 1 && q % a == 0 {
				if a == b || a * a == q || b == 0 {
					h++
				} else {
					h += 2
				}
			}
			a++
		}
	}
	return h
}

func Sqrt(z *big.Int) *big.Int {
	if z.Sign() < 0 {
		return nil
	} else if z.Sign() == 0 {
		return big.NewInt(0)
	} else if z.Cmp(intOne) == 0 {
		return big.NewInt(1)
	}


  // initial guess
	s := big.NewInt(0).Rsh(z, 1)
	t := big.NewInt(0)

	for s.Cmp(t) != 0 {
		// compute iteration
		t.Div(z, s)
		t.Add(t, s)
		t.Rsh(t, 1)
		s, t = t, s
	}
	return s
}

func IsSquare(z *big.Int) bool {
	s := Sqrt(z)
	s.Mul(s, s)
	return s.Cmp(z) == 0
}

func (p *IntPolynomial) Degree() int {
	return len(p.coeffs) - 1
}

var intOne = big.NewInt(1)

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
