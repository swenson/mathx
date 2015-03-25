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

package numtheory

import (
	"math"
	"math/big"
)

var intOne = big.NewInt(1)

// Compute the class number of an imaginary, quadratic number field.
// Based on Henri Cohen, _A Course in Algebraic Number Theory_, Alg 5.3.5.
func classNumberImagQuadSlow(k *NumberField) int {
	D := k.Discriminant().Int64()
	D = makeFundamentalDiscriminant(D)
	h := 1
	aD := D
	if aD < 0 {
		aD = -aD
	}
	B := int64(math.Floor(math.Sqrt(float64(aD) / 3.0)))

	for b := PosMod(D, 2); b <= B; b += 2 {
		q := (b*b - D) / 4
		a := b
		if a <= 1 {
			a = 1
		}
		for a == 1 || a*a <= q {
			if a != 1 && q%a == 0 {
				if a == b || a*a == q || b == 0 {
					h++
				} else {
					h += 2
				}
			}
			a++
		}
	}
	return h
}
