// Copyright (c) 2015 Christopher Swenson.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Package decimal is for arbitrary-precision decimal arithmetic.

Currently supported:

* String input and output

TODO: multiply, divide
TODO: roots, rounding, logarithms, exponentiation, and everything else.
TODO: optimization
TODO: use base-1000 packed in 10 bits
TODO: subtraction is really, really slow. Should use faster algorithm.
*/
package decimal

import (
	"fmt"
	"regexp"
)

// RoundingMode represents the rounding requested.
type RoundingMode int

const (
	_ = iota
	// RoundUp means take the ceiling after the operation.
	RoundUp RoundingMode = 1 * iota
	// RoundDown means take the floor after the operation.
	RoundDown
)

// Decimal is the basic type of our arbitrary-precision decimal numbers.
type Decimal struct {
	neg      bool
	whole    []int8
	fraction []int8
}

var zero = &Decimal{
	neg:      false,
	whole:    []int8{},
	fraction: []int8{},
}

var decimalRe = regexp.MustCompile(`-?([0-9]*)(\.[0-9]*)?`)

func parseDigits(s string) []int8 {
	if s == "" {
		return []int8{}
	}
	if s[0] == '.' {
		s = s[1:]
	}
	arr := make([]int8, len(s), len(s))
	for i := range s {
		arr[i] = int8(s[i] - '0')
	}
	return arr
}

// New constructs a new Decimal from a string.
func New(s string) (*Decimal, error) {
	d := new(Decimal)
	if !decimalRe.MatchString(s) {
		return nil, fmt.Errorf("Unknown format for decimal number")
	}
	match := decimalRe.FindStringSubmatch(s)
	d.neg = false
	if s[0] == '-' {
		d.neg = true
	}
	d.whole = parseDigits(match[1])
	d.fraction = parseDigits(match[2])
	return d, nil
}

func digitsToString(digits []int8) string {
	out := make([]byte, len(digits), len(digits))
	for i := range digits {
		out[i] = byte(digits[i] + '0')
	}
	return string(out)
}

func (d *Decimal) String() string {
	sign := ""
	if d.neg {
		sign = "-"
	}
	whole := digitsToString(d.whole)
	point := ""
	if len(d.fraction) > 0 {
		point = "."
	}
	frac := digitsToString(d.fraction)
	s := fmt.Sprintf("%s%s%s%s", sign, whole, point, frac)
	if s == "" {
		return "0"
	}
	return s
}

func imax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func leftExtend(a []int8, length int) []int8 {
	if length <= len(a) {
		return a
	}
	b := make([]int8, length, length)
	copy(b[length-len(a):], a)
	return b
}

func rightExtend(a []int8, length int) []int8 {
	if length <= len(a) {
		return a
	}
	b := make([]int8, length, length)
	copy(b, a)
	return b
}

// Neg returns the negative of this number.
func (d *Decimal) Neg() *Decimal {
	x := new(Decimal)
	x.neg = !d.neg
	x.whole = make([]int8, len(d.whole), len(d.whole))
	copy(x.whole, d.whole)
	x.fraction = make([]int8, len(d.fraction), len(d.fraction))
	copy(x.fraction, d.fraction)
	return x
}

// Abs returns the absolute value of d, i.e., if d < 0, then it returns -d.
func (d *Decimal) Abs() *Decimal {
	x := new(Decimal)
	x.neg = false
	x.whole = make([]int8, len(d.whole), len(d.whole))
	copy(x.whole, d.whole)
	x.fraction = make([]int8, len(d.fraction), len(d.fraction))
	copy(x.fraction, d.fraction)
	return x
}

// Cmp compares d to e and returns 1 if d > e, 0 if d == e, and -1 if d < e.
func (d *Decimal) Cmp(e *Decimal) int {
	if !d.neg && e.neg {
		return 1
	} else if d.neg && !e.neg {
		return -1
	}
	if d.neg && e.neg {
		// both negative
		return -d.Neg().Cmp(e.Neg())
	}
	// both positive
	if len(d.whole) > len(e.whole) {
		return 1
	} else if len(e.whole) > len(d.whole) {
		return -1
	}
	for i := 0; i < len(d.whole); i++ {
		if d.whole[i] > e.whole[i] {
			return 1
		} else if d.whole[i] < e.whole[i] {
			return -1
		}
	}
	// have to check fractions
	for i := 0; ; i++ {
		if i < len(d.fraction) && i < len(e.fraction) {
			if d.fraction[i] < e.fraction[i] {
				return 1
			} else if d.fraction[i] > e.fraction[i] {
				return -1
			} else {
				continue
			}
		}
		if i >= len(d.fraction) && i >= len(e.fraction) {
			return 0 // equal
		} else if i >= len(d.fraction) && i < len(e.fraction) && e.fraction[i] != 0 {
			return -1
		} else if i >= len(d.fraction) && i < len(e.fraction) && e.fraction[i] == 0 {
			continue // have to keep going
		} else if i < len(d.fraction) && i >= len(e.fraction) && d.fraction[i] != 0 {
			return 1
		} else {
			continue // have to keep going
		}
	}
}

// Sub returns this minus its argument.
func (d *Decimal) Sub(e *Decimal) *Decimal {
	cmp := d.Cmp(e)
	if cmp == 0 {
		return zero
	} else if cmp < 0 {
		return e.Sub(d).Neg()
	}
	// we can now be sure that d > e

	if d.neg != e.neg {
		// this is really addition, and e is negative
		return d.Add(e.Neg())
	}

	if d.neg && e.neg {
		// both negative, so we need to swap
		return e.Neg().Sub(d.Neg())
	}

	// both positive, d > e
	s := new(Decimal)
	s.neg = d.neg
	borrow := int8(0)
	flen := imax(len(d.fraction), len(e.fraction))
	df := rightExtend(d.fraction, flen)
	ef := rightExtend(e.fraction, flen)
	s.fraction = make([]int8, flen, flen)
	for i := flen - 1; i >= 0; i-- {
		a := df[i]
		b := ef[i]
		sum := a - b + borrow
		borrow = sum / 10
		sum = sum % 10
		if sum < 0 {
			sum += 10
			borrow--
		}
		s.fraction[i] = sum
	}
	// borrow into the whole part
	wlen := imax(len(d.whole), len(e.whole)) + 1
	s.whole = make([]int8, wlen, wlen)
	dw := leftExtend(d.whole, wlen)
	ew := leftExtend(e.whole, wlen)
	for i := wlen - 1; i >= 0; i-- {
		a := dw[i]
		b := ew[i]
		sum := a - b + borrow
		borrow = sum / 10
		sum = sum % 10
		if sum < 0 {
			sum += 10
			borrow--
		}
		s.whole[i] = sum
	}
	if borrow != 0 {
		panic("Subtraction failed")
	}

	// normalize
	for len(s.whole) > 1 && s.whole[0] == 0 {
		s.whole = s.whole[1:]
	}
	return s
}

// Add returns the sum of this plus its argument.
func (d *Decimal) Add(e *Decimal) *Decimal {
	s := new(Decimal)
	if d.neg == e.neg {
		s.neg = d.neg
		carry := int8(0)
		flen := imax(len(d.fraction), len(e.fraction))
		df := rightExtend(d.fraction, flen)
		ef := rightExtend(e.fraction, flen)
		s.fraction = make([]int8, flen, flen)
		for i := flen - 1; i >= 0; i-- {
			a := df[i]
			b := ef[i]
			sum := a + b + carry
			carry = sum / 10
			sum = sum % 10
			s.fraction[i] = sum
		}
		// carry into the whole part
		wlen := imax(len(d.whole), len(e.whole)) + 1
		s.whole = make([]int8, wlen, wlen)
		dw := leftExtend(d.whole, wlen)
		ew := leftExtend(e.whole, wlen)
		for i := wlen - 1; i >= 0; i-- {
			a := dw[i]
			b := ew[i]
			sum := a + b + carry
			carry = sum / 10
			sum = sum % 10
			s.whole[i] = sum
		}
		if carry != 0 {
			panic("Addition didn't work correctly")
		}

		// normalize
		for len(s.whole) > 1 && s.whole[0] == 0 {
			s.whole = s.whole[1:]
		}
	} else {
		panic("Not implemented yet")
	}
	return s
}
