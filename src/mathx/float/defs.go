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

func (x *Float) Add(y *Float) *Float {
	z := new(Float)
	if x.sign && y.sign {
		z.sign = true
		for x.exp > y.exp {
			y.exp++
			y.mantissa = y.mantissa.Lsh(1)
		}
		for y.exp > x.exp {
			x.exp++
			x.mantissa = y.mantissa.Lsh(1)
		}
		z.exp = x.exp
		z.mantissa = x.mantissa.Add(y.mantissa)
	}
	return z.normalize()
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

	m := (*big.Int)(z.mantissa)
	return fmt.Sprintf("%s%se%d", sign, m.String(), z.exp)
}
