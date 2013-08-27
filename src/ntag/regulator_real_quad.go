// Copyright (c) 2013 Christopher Swenson.
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

package ntag

import (
	"math"
	"math/big"
)

// cohen, 5.7.2,  p. 270
func (poly *IntPolynomial) regulatorRealQuad() float64 {
	D := poly.Discriminant().Int64()
	f := math.Sqrt(float64(D))
	d := int64(math.Floor(f))
	var b int64
	if d&1 == D&1 {
		b = d
	} else {
		b = d - 1
	}
	var u, v int64
	u1 := -b
	u2 := int64(2)
	v1 := int64(1)
	v2 := int64(0)
	p := b
	q := int64(2)
	for {
		A := (p + d) / q
		t := p
		p = A*q - p
		if t == p && v2 != 0 {
			u = (u2*u2 + v2*v2*D) / q
			v = (2 * u2 * v2) / q
			break
		}
		u1, u2 = u2, A*u2+u1
		v1, v2 = v2, A*v2+v1
		t = q
		q = (D - p*p) / q

		if q == t && v2 != 0 {
			u = (u1*u2 + D*v1*v2) / q
			v = (u1*v2 + u2*v1) / q
			break
		}
	}
	squareparts := dumbSquareFactors(D)
	g := big.NewInt(0).GCD(nil, nil, big.NewInt(u), big.NewInt(v*squareparts)).Int64()
	return math.Log((math.Abs(float64(u/g)) + math.Abs(float64(v/g))*f) / 2.0)
}

func dumbSquareFactors(n int64) int64 {
	out := int64(1)
	for n%4 == 0 {
		out <<= 1
		n >>= 2
	}
	if n%1 == 0 {
		n >>= 1
	}
	s := math.Sqrt(float64(n))
	for p := 3; p <= s; p += 2 {
		for n%p == 0 {
			if (n/p)%p == 0 {
				out *= p
				n /= p
			}
			n /= p
		}
		if n == 1 {
			break
		}
	}
	return out
}
