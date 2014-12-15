// Copyright (c) 2014 Christopher Swenson.

package mathx

import (
	"math/big"
)

// Int is an immutable arbitrary-precision integer type, wrapping
// the built-in math/big.Int (which is mutable). This package
// also a simpler, two-argument API. For example:
//
//   a := mathx.NewInt(0)
//   b := a.Add64(123)
type Int big.Int

// NewInt returns the arbitrary-precision version of its 64-bit argument.
func NewInt(n int64) *Int {
	b := (*Int)(big.NewInt(n))
	return b
}

func NewIntFromString(s string, base int) (*Int, bool) {
	b, err := new(big.Int).SetString(s, base)
	return (*Int)(b), err
}

// internal only: return a copy. (Since the interface is immutable,
// this function is not useful externally.)
func (z *Int) copy() *Int {
	x := NewInt(0)
	(*big.Int)(x).Set((*big.Int)(z))
	return x
}

// String returns a string representation of this integer in base 10.
func (z *Int) String() string {
	return (*big.Int)(z).String()
}

// Cmp compares this integer (this) to the argument (x), returning something < 0 if
// this < x, == 0 if this == x, and > 0 if this > x.
func (z *Int) Cmp(x *Int) int {
	return (*big.Int)(z).Cmp((*big.Int)(x))
}

// Sign returns the sign of this integer (-1 if < 0, 1 if > 0, and 0 if == 0).
func (z *Int) Sign() int {
	return (*big.Int)(z).Sign()
}

// Bit returns the value of the n-th bit of this integer.
func (z *Int) Bit(n int) uint {
	return (*big.Int)(z).Bit(n)
}

// BitLen returns the size of this integer in bits.
func (z *Int) BitLen() int {
	return (*big.Int)(z).BitLen()
}

// Lsh returns this integer shifted left by n (that is, multiplied by 2^n).
func (z *Int) Lsh(n uint) *Int {
	t := big.NewInt(0)
	return (*Int)((*big.Int)(t).Lsh((*big.Int)(z), n))
}

// Rsh returns this integer shifted right by n (that is, divided by 2^n and truncated).
func (z *Int) Rsh(n uint) *Int {
	t := big.NewInt(0)
	return (*Int)((*big.Int)(t).Rsh((*big.Int)(z), n))
}

// Add returns this number added to the argument.
func (z *Int) Add(y *Int) *Int {
	t := big.NewInt(0)
	return (*Int)((*big.Int)(t).Add((*big.Int)(z), (*big.Int)(y)))
}

// Add64 returns this number added to the 64-bit argument.
func (z *Int) Add64(y int64) *Int {
	t := big.NewInt(y)
	return (*Int)((*big.Int)(t).Add((*big.Int)(z), t))
}

// Sub returns this number minus the argument.
func (z *Int) Sub(y *Int) *Int {
	t := big.NewInt(0)
	return (*Int)((*big.Int)(t).Sub((*big.Int)(z), (*big.Int)(y)))
}

// Sub64 returns this number minus the 64-bit argument.
func (z *Int) Sub64(y int64) *Int {
	t := big.NewInt(y)
	return (*Int)((*big.Int)(t).Sub((*big.Int)(z), t))
}

// Mul returns this number multiplied by the argument.
func (z *Int) Mul(y *Int) *Int {
	t := big.NewInt(0)
	return (*Int)((*big.Int)(t).Mul((*big.Int)(z), (*big.Int)(y)))
}

// Mul64 returns this number multiplied by the 64-bit argument.
func (z *Int) Mul64(y int64) *Int {
	t := big.NewInt(y)
	return (*Int)((*big.Int)(t).Mul((*big.Int)(t), (*big.Int)(z)))
}
