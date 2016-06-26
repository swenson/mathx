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
	"github.com/swenson/mathx"
	"github.com/swenson/mathx/poly"
)

type factor64 struct {
	prime    int64
	exponent int
}

// Discriminant returns the discriminant of the polynomial defining this
// number field.
func (k *NumberField) Discriminant() *mathx.Int {
	return Discriminant(k.polynomial)
}

// Discriminant returns the discriminant of this polynomial.
func Discriminant(p *poly.IntPolynomial) *mathx.Int {
	if p.Degree() == 2 {
		c, b, a := p.Coeff(0), p.Coeff(1), p.Coeff(2)
		ac4 := mathx.NewInt(4).Mul(a).Mul(c)
		return b.Mul(b).Sub(ac4)
	}
	return nil
}

// IsFundamentalDiscriminant returns true if the given discriminant is
// fundamental.
func IsFundamentalDiscriminant(D *mathx.Int) bool {
	d := D.Int64()
	absd := d
	if d < 0 {
		absd = -absd
	}
	if d == 1 {
		return true
	}
	if PosMod(d, 4) == 1 && IsSquareFree64(absd) {
		return true
	}
	if PosMod(d, 4) == 0 {
		if IsSquareFree64(absd / 4) {
			return PosMod(d/4, 4) == 2 || PosMod(d/4, 4) == 3
		}
	}
	return false
}

func makeFundamentalDiscriminant(D int64) int64 {
	if IsFundamentalDiscriminant(mathx.NewInt(D)) {
		return D
	}
	factors := Factorization64(D)
	for _, f := range factors {
		ex := f.exponent
		for ex >= 2 {
			D = D / (f.prime * f.prime)
			ex -= 2
		}
	}
	return D
}
