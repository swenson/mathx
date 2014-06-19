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
	. "mathx"
	"testing"
)

func TestFloatAdd(t *testing.T) {
	x := new(Float)
	x.sign = true
	x.exp = 1
	x.mantissa = NewInt(1)
	y := new(Float)
	y.sign = true
	y.exp = 1
	y.mantissa = NewInt(1)
	z := x.Add(y)

	one := NewInt(1)
	if z.mantissa.Cmp(one) != 0 || z.exp != 2 || !z.sign {
		fmt.Printf("%s + %s = %s\n", x, y, z)
		t.FailNow()
	}
}

func TestPrint23(t *testing.T) {
	if NewFloat(2.3).String() != "+2.29999999999999982236431605997495353221893310546875" {
		fmt.Printf("Expected +2.29999999999999982236431605997495353221893310546875 got %s\n", NewFloat(2.3).String())
		t.FailNow()
	}
}
