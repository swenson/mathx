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
	"github.com/swenson/mathx/poly"
)

// NumberField represents the field of rational numbers adjoined with the root
// of a polynomial. (WIP)
type NumberField struct {
	polynomial *poly.IntPolynomial
}

// MakeNumberField creates the number field defined by the root of the given
// number field.
func MakeNumberField(poly *poly.IntPolynomial) *NumberField {
	k := new(NumberField)
	k.polynomial = poly
	return k
}

// Degree gives the degree of the polynomial defining the number field.
func (k *NumberField) Degree() int {
	return k.polynomial.Degree()
}

// ClassNumber computes the class number of the number field.
// Currently only supports imaginary quadratic number fields.
func (k *NumberField) ClassNumber() int {
	if k.Degree() == 2 {
		if Discriminant(k.polynomial).Sign() < 0 {
			return classNumberImagQuadSlow(k)
		}
	}
	return -1
}
