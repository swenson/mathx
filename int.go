// Copyright (c) 2014 Christopher Swenson.

package mathx

import (
	"fmt"
	"math/big"
	"math/rand"
)

// Int is an immutable arbitrary-precision integer type, wrapping
// the built-in math/big.Int (which is mutable). This package
// also has a simpler, two-argument API. For example:
//
//   a := mathx.NewInt(0)
//   b := a.Add64(123)
type Int big.Int

// NewInt returns the arbitrary-precision version of its 64-bit argument.
func NewInt(n int64) *Int {
	b := (*Int)(big.NewInt(n))
	return b
}

// NewIntFromString creates a new arbitrary-precision integer from
// a string in the given base.
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

// Int64 returns this number as a 64-bit integer, if it is able to.
// If it does not fit in 64 bits, the result is undefined.
func (z *Int) Int64() int64 {
	return (*big.Int)(z).Int64()
}

// Div returns this divided by y.
func (z *Int) Div(y *Int) *Int {
	x := (*big.Int)(z.copy())
	return (*Int)(x.Div(x, (*big.Int)(y)))
}

// Abs returns the absolute value of this integer.
func (z *Int) Abs() *Int {
	return (*Int)(new(big.Int).Abs((*big.Int)(z)))
}

// And returns the value of this bitwise-ANDed to the argument.
func (z *Int) And(x *Int) *Int {
	return (*Int)(new(big.Int).And((*big.Int)(z), (*big.Int)(x)))
}

// AndNot returns the value of this bitwise-AND-NOTed to the argument.
func (z *Int) AndNot(x, y *Int) *Int {
	return (*Int)(new(big.Int).AndNot((*big.Int)(z), (*big.Int)(x)))
}

// Append returns the value of this with the given string (and base) appended.
func (z *Int) Append(buf []byte, base int) []byte {
	x := z.copy()
	return (*big.Int)(x).Append(buf, base)
}

// Binomial returns n choose k.
func Binomial(n, k int64) *Int {
	return (*Int)(new(big.Int).Binomial(n, k))
}

// Bits returns (a copy of) the underlying raw data.
func (z *Int) Bits() []big.Word {
	return (*big.Int)(z.copy()).Bits()
}

// Bytes returns (a copy of) the underlying raw data.
func (z *Int) Bytes() []byte {
	return (*big.Int)(z.copy()).Bytes()
}

// DivMod returns the quotient and remainder when divided by y, modulo m, using Euclidean
// (non-Go standard) division.
func (z *Int) DivMod(y, m *Int) (*Int, *Int) {
	a, b := new(big.Int).DivMod((*big.Int)(z), (*big.Int)(y), (*big.Int)(m))
	return (*Int)(a), (*Int)(b)
}

// Exp returns this**y, modulo m if m != 0.
func (z *Int) Exp(y, m *Int) *Int {
	return (*Int)(new(big.Int).Exp((*big.Int)(z), (*big.Int)(y), (*big.Int)(m)))
}

// Format sets the state to this formatted as specified by the conversion character.
func (z *Int) Format(s fmt.State, ch rune) {
	(*big.Int)(z).Format(s, ch)
}

// GCD returns the GCD of this and y.
func (z *Int) GCD(y *Int) *Int {
	return (*Int)(new(big.Int).GCD((*big.Int)(z), (*big.Int)(y), nil, nil))
}

// ExtendedGCD returns the GCD (g) of this and b, and also returns
// x and y such that this * x + b * y = g.
func (z *Int) ExtendedGCD(b *Int) (*Int, *Int, *Int) {
	x := new(Int)
	y := new(Int)
	g := (*Int)(new(big.Int).GCD((*big.Int)(z), (*big.Int)(b), (*big.Int)(x), (*big.Int)(y)))
	return g, x, y
}

// GobDecode decodes the data from the buffer, and changes this.
// WARNING: this breaks immutability since it can change the underlying data.
// This is unavoidable with the way that Gob works.
func (z *Int) GobDecode(buf []byte) error {
	return (*big.Int)(z).GobDecode(buf)
}

// GobEncode encodes this as a gob byte array.
func (z *Int) GobEncode() ([]byte, error) {
	return (*big.Int)(z).GobEncode()
}

// MarshalJSON marshals this to a JSON byte array.
func (z *Int) MarshalJSON() ([]byte, error) {
	return (*big.Int)(z).MarshalJSON()
}

// MarshalText marshals this to a text string.
func (z *Int) MarshalText() (text []byte, err error) {
	return (*big.Int)(z).MarshalText()
}

// Mod returns this modulo the argument.
func (z *Int) Mod(y *Int) *Int {
	return (*Int)(new(big.Int).Mod((*big.Int)(z), (*big.Int)(y)))
}

// ModInverse returns the inverse of this modulo the argument.
func (z *Int) ModInverse(n *Int) *Int {
	return (*Int)(new(big.Int).ModInverse((*big.Int)(z), (*big.Int)(n)))
}

// ModSqrt returns the square root of this modulo the argument.
func (z *Int) ModSqrt(p *Int) *Int {
	return (*Int)(new(big.Int).ModSqrt((*big.Int)(z), (*big.Int)(p)))
}

// MulRange returns the product of all integers between a and b (inclusive). argument.
func MulRange(a, b int64) *Int {
	return (*Int)(new(big.Int).MulRange(a, b))
}

// Neg returns this with the sign flipped.
func (z *Int) Neg() *Int {
	return (*Int)(new(big.Int).Neg((*big.Int)(z)))
}

// Not returns this with all bits inverted.
func (z *Int) Not() *Int {
	return (*Int)(new(big.Int).Not((*big.Int)(z)))
}

// Or returns this bitwise-ORed with the argument.
func (z *Int) Or(y *Int) *Int {
	return (*Int)(new(big.Int).Or((*big.Int)(z), (*big.Int)(y)))
}

// ProbablyPrime retusn whether or not this is probably prime (with certainty 1 - ¼ⁿ).
func (z *Int) ProbablyPrime(n int) bool {
	return (*big.Int)(z).ProbablyPrime(n)
}

// Quo returns this divided by y with truncated (Go) division.
func (z *Int) Quo(y *Int) *Int {
	return (*Int)(new(big.Int).Quo((*big.Int)(z), (*big.Int)(y)))
}

// QuoRem returns the quotient and remainder of this divided by y with truncated (Go) division.
func (z *Int) QuoRem(y *Int) (*Int, *Int) {
	r := new(big.Int)
	a, b := new(big.Int).QuoRem((*big.Int)(z), (*big.Int)(y), r)
	return (*Int)(a), (*Int)(b)
}

// Rand returns a random number between 0 (inclusive) and n (exclusive).
func Rand(rnd *rand.Rand, n *Int) *Int {
	return (*Int)(new(big.Int).Rand(rnd, (*big.Int)(n)))
}

// Rem returns the remainder (using truncated Go division) of this divided by y.
func (z *Int) Rem(y *Int) *Int {
	return (*Int)(new(big.Int).Rem((*big.Int)(z), (*big.Int)(y)))
}

// Scan implements the fmt.Scanner interface, and changes the underling big.Int.
// WARNING: this breaks immutability since it can change the underlying data.
// This is unavoidable with the way that scanning works.
func (z *Int) Scan(s fmt.ScanState, ch rune) error {
	return (*big.Int)(z).Scan(s, ch)
}

// SetBit returns this with the ith bit set to b.
func (z *Int) SetBit(i int, b uint) *Int {
	return (*Int)(new(big.Int).SetBit((*big.Int)(z), i, b))
}

// SetBits returns a new (copy) set from the given data.
func SetBits(abs []big.Word) *Int {
	return (*Int)(new(big.Int).SetBits(abs)).copy()
}

// SetBytes returns a new (copy) set from the given data.
func (z *Int) SetBytes(buf []byte) *Int {
	return (*Int)(new(big.Int).SetBytes(buf)).copy()
}

// Text returns this printed as a string in the given base.
func (z *Int) Text(base int) string {
	return (*big.Int)(z).Text(base)
}

// Uint64 retusn this casted to a uint64.
func (z *Int) Uint64() uint64 {
	return (*big.Int)(z).Uint64()
}

// UnmarshalJSON unmarshals the JSON buffer into this.
// WARNING: this breaks immutability since it can change the underlying data.
// This is unavoidable with the way that unmarshaling works.
func (z *Int) UnmarshalJSON(text []byte) error {
	return (*big.Int)(z).UnmarshalJSON(text)
}

// UnmarshalText unmarshals the text buffer into this.
// WARNING: this breaks immutability since it can change the underlying data.
// This is unavoidable with the way that unmarshaling works.
func (z *Int) UnmarshalText(text []byte) error {
	return (*big.Int)(z).UnmarshalText(text)
}

// Xor returns this bitwise-XORed with the argument.
func (z *Int) Xor(y *Int) *Int {
	return (*Int)(new(big.Int).Xor((*big.Int)(z), (*big.Int)(y)))
}
