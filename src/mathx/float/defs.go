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
	} else {
		x.sign = false
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

	if (x.sign && y.sign) || (!x.sign && !y.sign) {
		for x.exp > y.exp {
			y.exp++
			y.mantissa = y.mantissa.Lsh(1)
		}
		for y.exp > x.exp {
			x.exp++
			x.mantissa = x.mantissa.Lsh(1)
		}
		z.exp = x.exp
		z.mantissa = x.mantissa.Add(y.mantissa)
		z.sign = x.sign
	} else if x.sign == true {
		for x.exp > y.exp {
			y.exp++
			y.mantissa = y.mantissa.Lsh(1)
			z.sign = x.sign
		}
		for y.exp > x.exp {
			x.exp++
			x.mantissa = x.mantissa.Lsh(1)
			z.sign = y.sign
		}
		z.exp = x.exp
		z.mantissa = x.mantissa.Sub(y.mantissa)
	} else if y.sign == true {
		for x.exp > y.exp {
			y.exp++
			y.mantissa = y.mantissa.Lsh(1)
			z.sign = x.sign
		}
		for y.exp > x.exp {
			x.exp++
			x.mantissa = x.mantissa.Lsh(1)
			z.sign = y.sign
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

func (z *Float) normalize() *Float {
	if z.mantissa.Sign() == 0 {
		return z
	}

	for z.mantissa.Bit(0) == 0 {
		z.mantissa = z.mantissa.Rsh(1)
		z.exp++
	}
	return z
}

func (z Float) String() string {
	sign := "+"
	if !z.sign {
		sign = "-"
	}

	var whole *Int
	var fraction *Int

	if z.exp < 0 {
		whole = z.mantissa.Rsh(uint(-z.exp))
		fraction = z.mantissa.Sub(whole.Lsh(uint(-z.exp)))
	} else {
		whole = NewInt(0)
		fraction = z.mantissa
	}
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
