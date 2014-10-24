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

var floatAddTestCases = []struct {
	a        *Float
	b        *Float
	c        *Float
	sign     bool
	exp      int64
	mantissa *Int
}{
	//16 cases, positive negative, between zero and one and greater than one
	{NewFloat(1.125), NewFloat(40), NewFloat(41.125), true, -3, NewInt(329)},
	{NewFloat(2), NewFloat(0.0009765625), NewFloat(2.0009765625), true, -10, NewInt(2049)},
	{NewFloat(0.0087890625), NewFloat(1.75), NewFloat(1.7587890625), true, -10, NewInt(1801)},
	{NewFloat(0.28125), NewFloat(0.46875), NewFloat(0.75), true, -2, NewInt(3)},
	{NewFloat(5000), NewFloat(-4099), NewFloat(901), true, 0, NewInt(901)},
	{NewFloat(2), NewFloat(-0.05859375), NewFloat(1.94140625), true, -8, NewInt(497)},
	{NewFloat(0.00244140625), NewFloat(-20), NewFloat(-19.99755859375), false, -11, NewInt(-40955)},
	{NewFloat(0.03125), NewFloat(-0.375), NewFloat(-0.34375), false, -5, NewInt(-11)},
	/*{NewFloat(-1.1), NewFloat(2.6), NewFloat(1.5), true, -1, NewInt(3)},
	/*{NewFloat(-), NewFloat(0.0015869140625), NewFloat(), , ,NewInt()},
	/*{NewFloat(-0.0234375), NewFloat(), NewFloat(), , ,NewInt()},
	/*{NewFloat(-0.0546875), NewFloat(0.34375), NewFloat(), , ,NewInt()},
	/*{NewFloat(-), NewFloat(-), NewFloat(), , ,NewInt()},
	/*{NewFloat(-), NewFloat(-0.002685546875), NewFloat(), , ,NewInt()},
	{NewFloat(-0.001708984375), NewFloat(-), NewFloat(), , ,NewInt()},
	{NewFloat(-0.875), NewFloat(-0.25), NewFloat(), , ,NewInt()},
	{NewFloat(-0.0), NewFloat(1), NewFloat(1), true, 1, NewInt(1)}, this bit fails */
}

func TestFloatAdd(t *testing.T) {
	x := new(Float)
	y := new(Float)
	for _, testCase := range floatAddTestCases {
		x = testCase.a
		y = testCase.b
		z := x.Add(y)
		mantissaCmp := testCase.mantissa.Cmp(z.mantissa)
		if z.sign != testCase.sign || z.exp != testCase.exp || mantissaCmp != 0 {
			fmt.Printf("%v + %v =%v\n", x, y, z)
			fmt.Printf("%t and %v and %v also %t and %v and %v\n", z.sign, z.exp, z.mantissa, testCase.sign, testCase.exp, testCase.mantissa)
			t.FailNow()
		}
	}
}

func TestFloatSub(t *testing.T) {
	x := new(Float)
	y := new(Float)
	for _, testCase := range floatAddTestCases {
		x = testCase.a
		y = testCase.b
		y.sign = !y.sign
		z := x.Sub(y)
		mantissaCmp := testCase.mantissa.Cmp(z.mantissa)
		if z.sign != testCase.sign || z.exp != testCase.exp || mantissaCmp != 0 {
			fmt.Printf("%v - %v =%v\n", x, y, z)
			fmt.Printf("%t and %v and %v also %t and %v and %v\n", z.sign, z.exp, z.mantissa, testCase.sign, testCase.exp, testCase.mantissa)
			t.FailNow()
		}
	}
}

var floatMulTestCases = []struct {
	a        *Float
	b        *Float
	c        *Float //I don't think this variable does anything
	sign     bool
	exp      int64
	mantissa *Int
}{
	{NewFloat(2), NewFloat(0.5), NewFloat(1), true, 0, NewInt(1)},
}

func TestFloatMul(t *testing.T) {
	x := new(Float)
	y := new(Float)
	for _, testCase := range floatMulTestCases {
		x = testCase.a
		y = testCase.b
		z := x.Mul(y)
		mantissaCmp := testCase.mantissa.Cmp(z.mantissa)
		if z.sign != testCase.sign || z.exp != testCase.exp || mantissaCmp != 0 {
			fmt.Printf("%t and %v and %v also %t and %v and %v\n", z.sign, z.exp, z.mantissa, testCase.sign, testCase.exp, testCase.mantissa)
			t.FailNow()
		}
	}
}

/*func TestFloatMakeSeventeen(t *testing.T) {
	z := NewFloat(1.0)
	z = MakeSeventeen()
	fmt.Printf("this is z %v\n", z)
	t.FailNow()
}*/

var floatDivTestCases = []struct {
	a        *Float
	b        *Float
	c        *Float //I don't think this variable does anything
	sign     bool
	exp      int64
	mantissa *Int
}{
	{NewFloat(2.0), NewFloat(2.0), NewFloat(1.0), true, 1, NewInt(1)},
	{NewFloat(20.0), NewFloat(2.0), NewFloat(10.0), true, 1, NewInt(10)},
	//{NewFloat(10.0), NewFloat(0.5), NewFloat(20.0), true, 1, NewInt(20)},
}

func TestFloatDiv(t *testing.T) {
	precision := NewFloat(2)
	precision.exp = precision.exp - 50
	for _, testCase := range floatDivTestCases { //floatDivTestCases or floatMulTestCAses?
		x := testCase.a
		y := testCase.b
		w := testCase.c
		z := x.Div(y)
		z, w = z.denormalize(w)
		diff := z.Sub(w)
		yes := diff.Cmp(precision)
		if z.sign != w.sign || z.exp != w.exp || yes >= 0 {
			fmt.Printf("%t and %v and %v \n%t and %v and %v\n", z.sign, z.exp, z.mantissa, w.sign, w.exp, w.mantissa)
			t.FailNow()
		}
	}
}

/*var floatStringTestCases = []struct {
	num float64
	str string
}{
	{0.0, "-0"},
	{1 / 17.0, "+0.05882352941176470506601248189326724968850612640380859375"},
	{1.0, "+1"},
	{1.1, "+1.100000000000000088817841970012523233890533447265625"},
	{-1.1, "-1.100000000000000088817841970012523233890533447265625"},
	{2.3, "+2.29999999999999982236431605997495353221893310546875"},
	{-6568408355712890880, "-6568408355712890880"},
}

func TestFloatPrint(t *testing.T) {
	for _, testCase := range floatStringTestCases {
		if NewFloat(testCase.num).String() != testCase.str {
			fmt.Printf("Expected %s got %s\n", testCase.str, NewFloat(testCase.num).String())
			t.FailNow()
		}
	}
}*/
