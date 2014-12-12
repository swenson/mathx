// Copyright (c) 2014 Christopher Swenson.
package mathx

import (
	"math/big"
)

// Int is a wrapper around the builtin math/big.Int that provides
// a simpler, two-argument API. For example:
// 	a := mathx.NewInt(0)
//  b := a.Add64(123)
type Int big.Int

func NewInt(n int64) *Int {
	b := (*Int)(big.NewInt(n))
	return b
}

func (z *Int) Copy() *Int {
	x := NewInt(0)
	(*big.Int)(x).Set((*big.Int)(z))
	return x
}

func (z *Int) String() string {
	return (*big.Int)(z).String()
}

func (z *Int) Cmp(x *Int) int {
	return (*big.Int)(z).Cmp((*big.Int)(x))
}

func (z *Int) Sign() int {
	return (*big.Int)(z).Sign()
}

func (z *Int) Bit(n int) uint {
	return (*big.Int)(z).Bit(n)
}

func (z *Int) BitLen() int {
	return (*big.Int)(z).BitLen()
}

func (x *Int) Lsh(n uint) *Int {
	z := big.NewInt(0)
	return (*Int)((*big.Int)(z).Lsh((*big.Int)(x), n))
}

func (x *Int) Rsh(n uint) *Int {
	z := big.NewInt(0)
	return (*Int)((*big.Int)(z).Rsh((*big.Int)(x), n))
}

func (x *Int) Add(y *Int) *Int {
	z := big.NewInt(0)
	return (*Int)((*big.Int)(z).Add((*big.Int)(x), (*big.Int)(y)))
}

func (x *Int) Add64(y int64) *Int {
	z := big.NewInt(y)
	return (*Int)((*big.Int)(z).Add((*big.Int)(x), z))
}

func (x *Int) Sub(y *Int) *Int {
	z := big.NewInt(0)
	return (*Int)((*big.Int)(z).Sub((*big.Int)(x), (*big.Int)(y)))
}

func (x *Int) Sub64(y int64) *Int {
	z := big.NewInt(y)
	return (*Int)((*big.Int)(z).Sub((*big.Int)(x), z))
}

func (x *Int) Mul(y *Int) *Int {
	z := big.NewInt(0)
	return (*Int)((*big.Int)(z).Mul((*big.Int)(x), (*big.Int)(y)))
}

func (x *Int) Mul64(y int64) *Int {
	z := big.NewInt(y)
	return (*Int)((*big.Int)(z).Mul((*big.Int)(z), (*big.Int)(x)))
}
