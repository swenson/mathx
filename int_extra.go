package mathx

// This file is for Int operations that are not in the standard library.

import (
	"math"
	"math/big"
)

// Float returns this converted to a Float.
func (z *Int) Float() *Float {
	return (*Float)(new(big.Float).SetInt((*big.Int)(z)))
}

// Sqrt computes the square root of this number.
// Uses Newton's Method.
func (z *Int) Sqrt() *Int {
	if z.BitLen() <= 52 {
		sqrt := int64(math.Sqrt(float64(z.Int64())))
		return NewInt(sqrt)
	}
	if z.Sign() < 0 {
		return nil
	} else if z.Sign() == 0 {
		return NewInt(0)
	} else if z.Cmp(NewInt(1)) == 0 {
		return NewInt(1)
	}

	// initial guess
	s := z.Rsh(uint(z.BitLen() / 2))
	t := NewInt(0)

	for i := 0; s.Cmp(t) != 0 && i < z.BitLen()/2+10; i++ {
		// compute iteration
		t = z.Div(s)
		t = t.Add(s)
		t = t.Rsh(1)
		s, t = t, s
	}
	return s
}

// IsSquare returns true if this number is a perfect square.
func IsSquare(z *Int) bool {
	if z.Sign() < 0 {
		return false
	}
	s := z.Sqrt()
	s = s.Mul(s)
	return s.Cmp(z) == 0
}
