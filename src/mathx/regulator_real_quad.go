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

package mathx

import (
	"math"
)

// cohen, 5.7.2,  p. 270
func (poly *IntPolynomial) regulatorRealQuad() float64 {
	D := poly.Discriminant().Int64()
	f := math.Sqrt(float64(D))
	c := -poly.coeffs[0].Int64()
	// try all of the small elements
	for x := int64(1); x < 100; x++ {
		for y := int64(1); y < 100; y++ {
			z := x*x - c*y*y
			if z == 4 || z == -4 {
				u := x
				v := y
				return math.Log((float64(u) + float64(v)*math.Sqrt(float64(c))) / 2.0)
			}
		}
	}
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
	return math.Log((math.Abs(float64(u)) + math.Abs(float64(v))*f) / 2.0)
}
