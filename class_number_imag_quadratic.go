package ntag

import (
	"math"
	"math/big"
 )

var intOne = big.NewInt(1)

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
