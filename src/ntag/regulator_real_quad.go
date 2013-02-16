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
	"fmt"
	"math"
)

func (poly *IntPolynomial) regulatorRealQuad() float64 {
	D := poly.Discriminant().Int64()
	f := math.Sqrt(float64(poly.Discriminant().Int64()))
	d := int64(math.Floor(f))
	e := 0
	R := 1.0
	b, a := poly.coeffs[1].Int64(), poly.coeffs[2].Int64()
	p := b
	q := a << 1
	q1 := (D - p*p) / q
	fmt.Println("q1", q1)
	L := 1000.0 // TODO: fix this
	L2 := math.Exp2(L)
	fmt.Printf("L = %f\n", L)
	fmt.Printf("L2 = %f\n", L2)

step2:
	A := (p + d) / q
	fmt.Println("A", A)
	r := PosMod(p+d, q)
	fmt.Println("r", r)
	p1 := p
	p = d - r
	t := q
	q = q1 - A*(p-p1)
	q1 = t
	R = R * (float64(p) + f) / float64(q)
	fmt.Println("R", R)
	if R > L2 {
		R = R / L2
		e++
	}
	a2 := a * 2
	if q == a2 && PosMod(p, a2) == PosMod(b, a2) {
		return math.Log(R) + float64(e)*L*math.Ln2
	}
	goto step2
}
