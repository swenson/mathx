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
	"math/big"
)

type RoundingMode int

const (
	_ = iota
	RoundUp RoundingMode = 1 * iota
	RoundDown
)

type Int big.Int

type Float struct {
	sign     bool
	exp      int64
  mantissa *Int
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

func (z *Int) Lsh(x *Int, n uint) *Int {
  return (*Int)((*big.Int)(z).Lsh((*big.Int)(x), n))
}

func (z *Int) Rsh(x *Int, n uint) *Int {
  return (*Int)((*big.Int)(z).Rsh((*big.Int)(x), n))
}

func (z *Int) Add(x *Int, y *Int) *Int {
  return (*Int)((*big.Int)(z).Add((*big.Int)(x), (*big.Int)(y)))
}

func (x *Float) Add(y *Float) *Float {
  z := new(Float)
  z.mantissa = (*Int)(big.NewInt(0))
  if x.sign && y.sign {
    z.sign = true
    for x.exp > y.exp {
      y.exp++
      y.mantissa.Lsh(y.mantissa, 1)
    }
    for y.exp > x.exp {
      x.exp++
      x.mantissa.Lsh(y.mantissa, 1)
    }
    z.exp = x.exp
    z.mantissa.Add(x.mantissa, y.mantissa)
  }
	return z.normalize()
}

func (z *Float) normalize() *Float {
  if z.mantissa.Sign() == 0 {
    return z
  }

  for z.mantissa.Bit(0) == 0 {
    z.mantissa.Rsh(z.mantissa, 1)
    z.exp++
  }
  return z
}

func (z Float) String() string {
  sign := "+"
  if !z.sign {
    sign = "-"
  }

  m := (*big.Int)(z.mantissa)
  return fmt.Sprintf("%s%se%d", sign, m.String(), z.exp)
}
