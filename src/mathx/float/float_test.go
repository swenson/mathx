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
	"testing"
)

var floatAddTestCases = []struct {
	a *Float
	b *Float
	c *Float
}{
	{NewFloat(-1.0), NewFloat(2.5), NewFloat(1.5)},
}

func TestFloatAdd(t *testing.T) {
	x := new(Float)
	y := new(Float)
	for _, testCase := range floatAddTestCases {
		x = testCase.a
		y = testCase.b
		z := x.Add(y)
		if z != testCase.c { //much broken so fail
			fmt.Printf("%v + %v =%v\n", x, y, z)
			t.FailNow()
		}
	}
}

var floatStringTestCases = []struct {
	num float64
	str string
}{
	{0.0, "-0.0"},
	{2.3, "+2.29999999999999982236431605997495353221893310546875"},
}

func TestFloatPrint(t *testing.T) {
	for _, testCase := range floatStringTestCases {
		if NewFloat(testCase.num).String() != testCase.str {
			fmt.Printf("Expected %s got %s\n", testCase.str, NewFloat(testCase.num).String())
			t.FailNow()
		}
	}
}
