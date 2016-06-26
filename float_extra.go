package mathx

import "math"

// Sqrt returns the square root of this number, computing using Newton's method.
func (z *Float) Sqrt() *Float {
	if z.Sign() == 0 {
		return z.copy()
	}
	if z.Sign() < 0 {
		panic("square root of a negative number is undefined")
	}
	if z.IsInf() {
		return z
	}

	number := z.SetPrec(z.Prec() + 53) // this will make sure that the loop has higher accuracy
	_, numberExp := number.MantExp()

	accuracy := NewFloat(1.0).SetPrec(number.Prec())
	accuracy = accuracy.SetExp(numberExp - int(number.Prec()))

	// start with a good guess
	zf, _ := z.Float64()
	var x *Float
	if math.IsInf(zf, 0) || math.IsNaN(zf) {
		x = NewFloat(1.0).SetPrec(number.Prec())
	} else {
		x = NewFloat(math.Sqrt(zf)).SetPrec(number.Prec())
	}
	two := NewFloat(2.0).SetPrec(number.Prec())
	denominator := NewFloat(1.0).SetPrec(number.Prec())
	delta := x.Mul(x).Sub(number).Abs()
	for delta.Cmp(accuracy) == 1 { // if the difference between the correct answer and the current guess is larger than the required accuracy, repeat
		prez := x
		denominator = two.Mul(prez)
		x = prez.Mul(prez)
		x = x.Sub(number)
		x = x.Quo(denominator)
		x = prez.Sub(x)
		delta = x.Mul(x).Sub(number).Abs()
	}
	return x
}
