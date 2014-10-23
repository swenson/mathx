// Copyright (c) 2014 Christopher Swenson.
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
	x.normalize()
	return x
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
	x := _x.Copy()
	y := _y.Copy()
	z := new(Float)

	z.precision = x.precision
	if z.precision > y.precision {
		z.precision = y.precision
	}

	if x.sign == y.sign {
		for x.exp < y.exp {
			y.exp--
			y.mantissa = y.mantissa.Lsh(1)
		}
		for y.exp < x.exp {
			x.exp--
			x.mantissa = x.mantissa.Lsh(1)
		}
		z.exp = x.exp
		z.mantissa = x.mantissa.Add(y.mantissa)
		z.sign = x.sign
	} else if x.sign == true {
		for x.exp < y.exp {
			y.exp--
			y.mantissa = y.mantissa.Lsh(1)
			z.sign = y.sign
		}
		for y.exp < x.exp {
			x.exp--
			x.mantissa = x.mantissa.Lsh(1)
			z.sign = x.sign
		}
		z.exp = x.exp
		z.mantissa = x.mantissa.Sub(y.mantissa)
	} else if y.sign == true {
		for x.exp < y.exp {
			y.exp--
			y.mantissa = y.mantissa.Lsh(1)
			z.sign = y.sign
		}
		for y.exp < x.exp {
			x.exp--
			x.mantissa = x.mantissa.Lsh(1)
			z.sign = x.sign
		}
		z.exp = x.exp
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

	for x.exp < y.exp {
		y.exp--
		y.mantissa = y.mantissa.Lsh(1)
	}
	for y.exp < x.exp {
		x.exp--
		x.mantissa = x.mantissa.Lsh(1)
	}
	z.exp = x.exp

	z.mantissa = x.mantissa.Mul(y.mantissa)
	z.exp = 2 * z.exp

	return z.normalize()
}

func (_x *Float) Div(_y *Float) *Float {
	//Theoretically this is the Newton Raphson method
	x := _x.Copy()
	y := _y.Copy()
	z := new(Float)

	if y.mantissa.Sign() == 0 {
		panic("Can not divide by zero")
	}

	z.precision = x.precision
	if z.precision > y.precision {
		z.precision = y.precision
	}

	//creating an accurate enough first guess
	thirtytwo := NewFloat(-32.0)
	fortyeight := NewFloat(48.0)
	//i := y.mantissa.BitLen() //the problem is in here somewhere
	//y.exp = 0 - int64(i)
	z = y.Mul(thirtytwo).Add(fortyeight)
	seventeen := NewFloat(0.0)
	seventeen.precision = 0
	repeatingchunk := NewFloat(15)
	for seventeen.precision < z.precision {
		seventeen.mantissa = seventeen.mantissa.Lsh(8).Add(repeatingchunk.mantissa)
		seventeen.precision = seventeen.precision + 8
		seventeen.exp = seventeen.exp - 8
	}
	z.Mul(seventeen)

	//create stopping point
	var stop float64
	stop = math.Log2((float64(z.precision) + 1) / (math.Log2(17))) //casting z.precision as a float64 should work up to 2^52 bits, hopefully
	stopp := int(math.Ceil(stop))

	one := NewFloat(1.0)
	prez := new(Float)
	for i := 0; i < stopp; i++ {
		prez = z
		z = prez.Mul(y)
		z = one.Sub(prez).Mul(prez).Add(prez)
		fmt.Printf("this is mantissa %v, this is exp %v, this is precision %v\n", z.mantissa, z.exp, z.precision)
	}
	z.Mul(x)
	fmt.Printf("this is z.mantissa %v, this is z.exp %v, this is z.precision %v\n", z.mantissa, z.exp, z.precision)

	return z.normalize()
}

func MakeSeventeen() *Float {
	seventeen := NewFloat(0.0)
	fmt.Printf("intialized at %v, precision is %v, exp is %v, mantissa is %v\n", seventeen, seventeen.precision, seventeen.exp, seventeen.mantissa)
	seventeen.precision = 0
	repeatingchunk := NewFloat(15)
	fmt.Printf("immediately before for loop%v\n", seventeen)
	for seventeen.precision < 64 {
		seventeen.mantissa = seventeen.mantissa.Lsh(8).Add(repeatingchunk.mantissa) //HERE is the problem
		seventeen.precision = seventeen.precision + 8
		fmt.Printf("in for loop %v seventeen, %v mantissa, %v precision, %v exp\n", seventeen, seventeen.mantissa, seventeen.precision, seventeen.exp)
		seventeen.exp = seventeen.exp - 8
	}
	fmt.Printf("after for loop, before return %v\n", seventeen.mantissa)
	return seventeen
}

func (_x *Float) Cmp(_y *Float) int {
	x := _x.Copy()
	y := _y.Copy()
	x.Sub(y)
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

func (z *Float) normalize() *Float {
	if z.mantissa.Sign() == 0 { //why no copy?
		return z
	}

	for z.mantissa.Bit(0) == 0 {
		z.mantissa = z.mantissa.Rsh(1)
		z.exp++
	}
	return z
}

func (x *Float) denormalize(y *Float) (*Float, *Float) {
	for x.exp < y.exp {
		y.exp--
		y.mantissa = y.mantissa.Lsh(1)
	}
	for y.exp < x.exp {
		x.exp--
		x.mantissa = x.mantissa.Lsh(1)
	}
	return x, y
}

func (z Float) String() string { // why is it Float not *Float
	sign := "+"
	if !z.sign {
		sign = "-"
	}

	var whole *Int
	var fraction *Int

	if z.exp >= 0 {
		whole = z.mantissa.Lsh(uint(z.exp))
		return fmt.Sprintf("%s%s", sign, whole.String())
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
