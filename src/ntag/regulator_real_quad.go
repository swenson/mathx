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
	"fmt"
	"math"
	"math/big"

//	"time"
)

func discTest(A, B, D *big.Int) bool {
	a2 := big.NewInt(0).Mul(A, A)
	b2 := big.NewInt(0).Mul(B, B)
	b2.Mul(b2, D)
	a2.Sub(a2, b2)
	fmt.Printf("%v^2 - %v * %v^2 = %v\n", A, D, B, a2)
	if a2.BitLen() <= 10 {
		aint := a2.Int64()
		return aint == 4 || aint == -4 || aint == 1 || aint == -1
	}
	return false
}

func (poly *IntPolynomial) fundamentalUnitRealQuad() (*big.Int, *big.Int) {
	D := poly.Discriminant()
	fmt.Printf("D = %v\n", D)
	if D.BitLen() <= 10 {
		dint := D.Int64()
		if dint == -3 || dint == 5 {
			return big.NewInt(1), big.NewInt(1)
		}
	}

	A0 := big.NewInt(1)
	B0 := big.NewInt(1)

	A1 := big.NewInt(1) // 2 + D - 1
	A1.Add(A1, D)
	B1 := big.NewInt(2)

	if discTest(A1, B1, D) {
		return A1, B1
	}

	An := big.NewInt(0)
	Bn := big.NewInt(0)
	D1 := big.NewInt(0).Sub(D, big.NewInt(1))
	two := big.NewInt(2)
	temp := big.NewInt(0)

	for i := 0; i < 10; i++ {
		//An := 2*A1 + (D-1)*A0
		//Bn := 2*B1 + (D-1)*B0
		An.Mul(A1, two)
		temp.Mul(D1, A0)
		An.Add(An, temp)
		Bn.Mul(B1, two)
		temp.Mul(D1, B0)
		Bn.Add(Bn, temp)

		r, _ := big.NewRat(1, 1).SetFrac(An, Bn).Float64()

		fmt.Printf("An = %s, Bn = %s, An/Bn = %f, %f\n", An.String(), Bn.String(), r, math.Sqrt(float64(D.Int64())))
		//time.Sleep(time.Duration(10000000))
		if discTest(An, Bn, D) {
			return An, Bn
		}
		A0.Set(A1)
		A1.Set(An)
		B0.Set(B1)
		B1.Set(Bn)
	}
	return An, Bn
}

func toFloat(a *big.Int) float64 {
	x, _ := big.NewRat(1, 1).SetInt(a).Float64()
	return x
}

// cohen, 5.7.1,  p. 270
func (poly *IntPolynomial) regulatorRealQuad() float64 {
	D := poly.Discriminant().Int64()
	fmt.Printf("D = %d\n", D)
	f := math.Sqrt(float64(poly.Discriminant().Int64()))
	d := int64(math.Floor(f))
	b, a := poly.coeffs[1].Int64(), poly.coeffs[2].Int64()
	u1 := -b
	u2 := int64(2) * a
	v1 := int64(1)
	v2 := int64(0)
	p := b
	q := 2 * a
step2:
	A := (p + d) / q
	fmt.Println("A", A)
	p = A*q - p
	q = (D - p*p) / q
	//q = (D - p*p) / q

	t := A*u2 + u1
	u1 = u2
	u2 = t

	t = A*v2 + v1
	v1 = v2
	v2 = t

	if q == 2*a && PosMod(p, 2*a) == PosMod(b, 2*a) {
		u := u2 / a
		if u < 0 {
			u = -u
		}
		v := v2 / a
		if v < 0 {
			v = -v
		}
		return math.Log((float64(u) + float64(v)*f) / 2.0)
	}
	goto step2
}
