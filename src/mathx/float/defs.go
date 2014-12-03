// Copyright (c) 2014 Christopher Swenson.
// Copyright (c) 2014 Georgia Reh.
// Copyright (c) 2012 Google, Inc. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//teeeeeeeeeest
package float

import (
	"fmt"
	"math"
	. "mathx"
)

type RoundingMode int

const (
	_                    = iota
	RoundUp RoundingMode = 1 * iota
	RoundDown
)

type Float struct {
	sign      bool
	precision uint64
	exp       int64
	mantissa  *Int
}

func NewFloat(f float64) *Float {
	x := new(Float)
	x.precision = 52
	// Convert from IEEE 754 double
	bits := math.Float64bits(f)
	s := bits >> 63
	e := int64((bits >> 52) & 0x7ff)
	m := int64(bits & uint64((int64(1)<<52)-1))
	if s == 0 && e == 0 && m == 0 {
		x.sign = false
		x.exp = 0
		x.mantissa = NewInt(0)
		return x
	}

	if s == 0 {
		x.sign = true
	}
	x.exp = e - 1023 - 52
	x.mantissa = NewInt((int64(1) << 52) | m)
	return x.normalize()
}

func (x *Float) Copy() *Float {
	y := NewFloat(0.0)
	y.sign = x.sign
	y.precision = x.precision
	y.exp = x.exp
	y.mantissa = x.mantissa.Copy()
	return y
}

func (_x *Float) Add(_y *Float) *Float {
	if _x.mantissa.Sign() == 0 {
		return _y
	} else if _y.mantissa.Sign() == 0 {
		return _x
	}

	x := _x.Copy()
	y := _y.Copy()
	z := new(Float)

	z.precision = x.precision
	if z.precision > y.precision {
		z.precision = y.precision
	}
	x, y = x.denormalize(y)
	z.exp = x.exp
	if x.sign == y.sign {
		z.sign = x.sign
		z.mantissa = x.mantissa.Add(y.mantissa)
	} else if x.mantissa.Cmp(y.mantissa) == 1 {
		z.sign = x.sign
		z.mantissa = x.mantissa.Sub(y.mantissa)
	} else if x.mantissa.Cmp(y.mantissa) < 1 {
		z.sign = y.sign
		z.mantissa = y.mantissa.Sub(x.mantissa)
	}
	return z.normalize()
}

func (_x *Float) Sub(_y *Float) *Float {
	x := _x.Copy()
	y := _y.Copy()
	z := new(Float)
	y.sign = !y.sign
	z = x.Add(y)
	return z
}

func (_x *Float) Mul(_y *Float) *Float {
	if _x.mantissa.Sign() == 0 || _y.mantissa.Sign() == 0 {
		return NewFloat(0.0)
	}

	x := _x.Copy()
	y := _y.Copy()
	z := new(Float)

	z.precision = x.precision
	if z.precision > y.precision {
		z.precision = y.precision
	}

	z.sign = true
	if x.sign != y.sign {
		z.sign = false
	}

	x, y = x.denormalize(y)
	z.exp = x.exp
	z.mantissa = x.mantissa.Mul(y.mantissa)
	z.exp = 2 * z.exp

	return z.normalize()
}

func (_x *Float) Div(_y *Float) *Float {
	//Div implements division by calculating the recipipricol of of the denominator and multiplying by the numerator using the Newton Raphson Method. (http://en.wikipedia.org/wiki/Division_algorithm)
	x := _x.Copy()
	y := _y.Copy()
	z := new(Float)

	if y.mantissa.Sign() == 0 {
		panic("division by zero is undefined\n")
	}
	if y.sign == x.sign {
		y.sign = true
		x.sign = true
	}
	if y.sign == false {
		y.sign = true
		x.sign = false
	}
	z.precision = x.precision
	if z.precision > y.precision {
		z.precision = y.precision
	}

	thirtytwo := NewFloat(-32.0)
	fortyeight := NewFloat(48.0)
	one := NewFloat(1.0)
	thirtytwo.precision = z.precision
	fortyeight.precision = z.precision
	one.precision = z.precision

	//Create an accurate enough initial guess
	i := y.mantissa.BitLen()
	tempexp := 0 - int64(i)
	x.exp = x.exp + (tempexp - y.exp)
	y.exp = tempexp
	z = y.Mul(thirtytwo)
	z = z.Add(fortyeight)
	seventeen := NewFloat(0.0)
	seventeen.sign = true
	seventeen.precision = 0
	repeatingchunk := NewFloat(15)
	for seventeen.precision < z.precision {
		seventeen.mantissa = seventeen.mantissa.Lsh(8).Add(repeatingchunk.mantissa)
		seventeen.precision = seventeen.precision + 8
		seventeen.exp = seventeen.exp - 8
	}
	z = z.Mul(seventeen)

	//Create stopping point for the for loop
	var stop float64
	stop = math.Log2((float64(z.precision) + 1))
	stopp := int(math.Ceil(stop))
	prez := new(Float)

	for i := 0; i < stopp; i++ {
		prez = z
		z = prez.Mul(y)
		z = one.Sub(z)
		z = z.Mul(prez)
		z = z.Add(prez)
	}
	z = z.Mul(x)
	return z.normalize()
}

func (_z *Float) Sqrt() *Float {
	//Sqrt uses Newton's Method
	if _z.mantissa.Sign() == 0 {
		return _z
	}
	if _z.sign == false {
		panic("square root of a negative number is undefined\n")
	}
	number := _z.Copy()
	accuracy := NewFloat(1.0)
	accuracy.exp += number.exp - int64(number.precision) + int64(number.mantissa.BitLen())
	accuracy.precision = 2 * number.precision
	accuracy.exp = accuracy.exp - int64(accuracy.precision) //this will make sure that the loop compares z^2 to accuracy^2
	number.precision = 2 * number.precision
	z := NewFloat(1.0)
	z.precision = 2 * z.precision
	two := NewFloat(2.0)
	two.precision = number.precision
	denominator := NewFloat(1.0)
	denominator.precision = 2 * denominator.precision
	delta := z.Mul(z).Sub(number).Abs()
	for delta.Cmp(accuracy) == 1 { //if the difference between the correct answer and the current guess is larger than the required accuracy, repeat
		prez := z
		denominator = two.Mul(prez)
		z = prez.Mul(prez)
		z = z.Sub(number)
		z = z.Div(denominator)
		z = prez.Sub(z)
		delta = z.Mul(z).Sub(number).Abs()
	}
	return z.normalize()
}

func (_x *Float) Cmp(_y *Float) int {
	x := _x.Copy()
	y := _y.Copy()

	x = x.Sub(y)

	var z int
	if x.mantissa.Sign() == 0 {
		z = 0
	} else if x.sign == true {
		z = 1
	} else if x.sign == false {
		z = -1
	}
	return z
}

func (_z *Float) Neg() *Float {
	z := _z.Copy()
	z.sign = !z.sign
	return z
}

func (_z *Float) Abs() *Float {
	z := _z.Copy()
	z.sign = true
	return z
}

func (z *Float) truncate() *Float {
	chop := z.mantissa.BitLen() - (2 * int(z.precision))
	if chop > 0 {
		z.mantissa = z.mantissa.Rsh(uint(chop))
		z.exp += int64(chop)
	}
	return z
}

func (z *Float) normalize() *Float {
	if z.mantissa.Sign() == 0 {
		return z
	}

	z.truncate()

	for z.mantissa.Bit(0) == 0 {
		z.mantissa = z.mantissa.Rsh(1)
		z.exp++
	}
	return z
}

func (x *Float) denormalize(y *Float) (*Float, *Float) {
	for x.exp < y.exp {
		y.mantissa = y.mantissa.Lsh(uint(y.exp - x.exp))
		y.exp = x.exp
	}
	for y.exp < x.exp {
		x.mantissa = x.mantissa.Lsh(uint(x.exp - y.exp))
		x.exp = y.exp
	}
	return x, y
}

func (z Float) String() string {
	sign := "+"
	if !z.sign {
		sign = "-"
	}

	var whole *Int
	var fraction *Int

	if z.exp <= 0 {
		whole = z.mantissa.Rsh(uint(-z.exp))
		fraction = z.mantissa.Sub(whole.Lsh(uint(-z.exp)))
	} else {
		whole = NewInt(0)
		fraction = z.mantissa
	}

	whole = z.mantissa.Rsh(uint(-z.exp))
	fraction = z.mantissa.Sub(whole.Lsh(uint(-z.exp)))

	digits := ""
	for fraction.Sign() != 0 {
		fraction = fraction.Mul64(10)
		digit := fraction.Rsh(uint(-z.exp))
		fraction = fraction.Sub(digit.Lsh(uint(-z.exp)))
		digits += digit.String()
	}
	if digits == "" {
		digits = "0"
	}
	return fmt.Sprintf("%s%s.%s", sign, whole.String(), digits)
}
