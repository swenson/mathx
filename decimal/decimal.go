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

TODO: add, subtract, multiply, divide
TODO: roots, rounding, logarithms, exponentiation, and everything else.
TODO: optimization
TODO: use base-1000 packed in 10 bits
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
	whole    []uint8
	fraction []uint8
}

var decimalRe = regexp.MustCompile(`-?([0-9]*)(\.[0-9]*)?`)

func parseDigits(s string) []uint8 {
	if s == "" {
		return []uint8{}
	}
	if s[0] == '.' {
		s = s[1:]
	}
	arr := make([]uint8, len(s), len(s))
	for i := range s {
		arr[i] = s[i] - '0'
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

func digitsToString(digits []uint8) string {
	out := make([]byte, len(digits), len(digits))
	for i := range digits {
		out[i] = digits[i] + '0'
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
	return fmt.Sprintf("%s%s%s%s", sign, whole, point, frac)
}

func imax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func leftExtend(a []uint8, length int) []uint8 {
	if length <= len(a) {
		return a
	}
	b := make([]uint8, length, length)
	copy(b[length-len(a):], a)
	return b
}

func rightExtend(a []uint8, length int) []uint8 {
	if length <= len(a) {
		return a
	}
	b := make([]uint8, length, length)
	copy(b, a)
	return b
}

func (d *Decimal) Add(e *Decimal) *Decimal {
	s := new(Decimal)
	if d.neg == e.neg {
		s.neg = d.neg
		carry := uint8(0)
		flen := imax(len(d.fraction), len(e.fraction))
		df := rightExtend(d.fraction, flen)
		ef := rightExtend(e.fraction, flen)
		s.fraction = make([]uint8, flen, flen)
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
		s.whole = make([]uint8, wlen, wlen)
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
