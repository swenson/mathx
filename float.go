package mathx

import "math/big"

// Float is an immutable arbitrary-precision floating-point type, wrapping
// the built-in math/big.Float (which is mutable). This package
// also has a simpler, two-argument API. For example:
//
//   a := mathx.NewFloat(0.0)
//   b := a.AddFloat64(123.0)
type Float big.Float

func NewFloat(x float64) *Float {
	return (*Float)(big.NewFloat(x))
}

func ParseFloat(s string, base int, prec uint, mode big.RoundingMode) (*Float, int, error) {
	f, b, err := big.ParseFloat(s, base, prec, mode)
	return (*Float)(f), b, err
}

func (z *Float) Abs() *Float {
	return (*Float)(new(big.Float).Abs((*big.Float)(z)))
}

func (z *Float) Acc() big.Accuracy {
	return (*big.Float)(z).Acc()
}

func (z *Float) Add(y *Float) *Float {
	return (*Float)(new(big.Float).Add((*big.Float)(z), (*big.Float)(y)))
}

func (z *Float) Cmp(y *Float) int {
	return (*big.Float)(z).Cmp((*big.Float)(y))
}

func (z *Float) copy() *Float {
	return (*Float)(new(big.Float).Set((*big.Float)(z)))
}

func (z *Float) Float32() (float32, big.Accuracy) {
	return (*big.Float)(z).Float32()
}

func (z *Float) Float64() (float64, big.Accuracy) {
	return (*big.Float)(z).Float64()
}

// func (z *Float) Format(s fmt.State, format rune) {
// 	(*big.Float)(z).Format(s, format)
// }

func (z *Float) Int() (*Int, big.Accuracy) {
	i := new(big.Int)
	i, acc := (*big.Float)(z).Int(i)
	return (*Int)(i), acc
}

func (z *Float) Int64() (int64, big.Accuracy) {
	return (*big.Float)(z).Int64()
}

func (z *Float) IsInf() bool {
	return (*big.Float)(z).IsInf()
}

func (z *Float) IsInt() bool {
	return (*big.Float)(z).IsInt()
}

func (z *Float) MantExp() (*Float, int) {
	y := new(big.Float)
	exp := (*big.Float)(z).MantExp(y)
	return (*Float)(y), exp
}

func (z *Float) MarshalText() (text []byte, err error) {
	return (*big.Float)(z).MarshalText()
}

func (z *Float) MinPrec() uint {
	return (*big.Float)(z).MinPrec()
}

func (z *Float) Mode() big.RoundingMode {
	return (*big.Float)(z).Mode()
}

func (z *Float) Mul(y *Float) *Float {
	return (*Float)(new(big.Float).Mul((*big.Float)(z), (*big.Float)(y)))
}

func (z *Float) Neg() *Float {
	return (*Float)(new(big.Float).Neg((*big.Float)(z)))
}

// func (z *Float) Parse(s string, base int) (f *Float, b int, err error) {
// }

func (z *Float) Prec() uint {
	return (*big.Float)(z).Prec()
}

func (z *Float) Quo(y *Float) *Float {
	return (*Float)(new(big.Float).Quo((*big.Float)(z), (*big.Float)(y)))
}

func (z *Float) Rat(r *big.Rat) (*big.Rat, big.Accuracy) {
	return (*big.Float)(z).Rat(r)
}

// func (z *Float) SetInf(signbit bool) *Float {
// }
//
//

func FloatFromInt64(x int64) *Float {
	return (*Float)(new(big.Float).SetInt64(x))
}

func (z *Float) SetExp(exp int) *Float {
	return (*Float)(new(big.Float).SetMantExp((*big.Float)(z), exp))
}

func (z *Float) SetMode(mode big.RoundingMode) *Float {
	return (*Float)((*big.Float)(z.copy()).SetMode(mode))
}

func (z *Float) SetPrec(prec uint) *Float {
	return (*Float)((*big.Float)(z.copy()).SetPrec(prec))
}

func FloatFromRat(x *big.Rat) *Float {
	return (*Float)(new(big.Float).SetRat(x))
}

func (z *Float) SetString(s string) (*Float, bool) {
	f, b := new(big.Float).SetString(s)
	return (*Float)(f), b
}

func FloatFromUint64(x uint64) *Float {
	return (*Float)(new(big.Float).SetUint64(x))
}

func (z *Float) Sign() int {
	return (*big.Float)(z).Sign()
}

func (z *Float) Signbit() bool {
	return (*big.Float)(z).Signbit()
}

func (z *Float) String() string {
	return (*big.Float)(z).String()
}

func (z *Float) Sub(y *Float) *Float {
	return (*Float)(new(big.Float).Sub((*big.Float)(z), (*big.Float)(y)))
}

func (z *Float) Text(format byte, prec int) string {
	return (*big.Float)(z).Text(format, prec)
}

func (z *Float) Uint64() (uint64, big.Accuracy) {
	return (*big.Float)(z).Uint64()
}

func (z *Float) UnmarshalText(text []byte) error {
	return (*big.Float)(z).UnmarshalText(text)
}
