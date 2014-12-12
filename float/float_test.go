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

	. "github.com/swenson/mathx"
)

var floatAddTestCases = []struct {
	a        *Float
	b        *Float
	c        *Float
	sign     bool
	exp      int64
	mantissa *Int
}{
	{NewFloat(1.125), NewFloat(40), NewFloat(41.125), true, -3, NewInt(329)},
	{NewFloat(2), NewFloat(0.0009765625), NewFloat(2.0009765625), true, -10, NewInt(2049)},
	{NewFloat(0.0087890625), NewFloat(1.75), NewFloat(1.7587890625), true, -10, NewInt(1801)},
	{NewFloat(0.28125), NewFloat(0.46875), NewFloat(0.75), true, -2, NewInt(3)},
	{NewFloat(5000), NewFloat(-4099), NewFloat(901), true, 0, NewInt(901)},
	{NewFloat(2), NewFloat(-0.05859375), NewFloat(1.94140625), true, -8, NewInt(497)},
	{NewFloat(0.00244140625), NewFloat(-20), NewFloat(-19.99755859375), false, -11, NewInt(40955)},
	{NewFloat(0.03125), NewFloat(-0.375), NewFloat(-0.34375), false, -5, NewInt(11)},
	{NewFloat(-1.1), NewFloat(2.6), NewFloat(1.5), true, -1, NewInt(3)},
	{NewFloat(-500.25), NewFloat(0.0015869140625), NewFloat(-500.2484130859375), false, -13, NewInt(4098035)},
	{NewFloat(-0.0234375), NewFloat(56.25), NewFloat(56.2265625), true, -7, NewInt(7197)},
	{NewFloat(-0.0546875), NewFloat(0.34375), NewFloat(0.2890625), true, -7, NewInt(37)},
	{NewFloat(-556), NewFloat(-66.42), NewFloat(-622.42), false, -46, NewInt(43798913751061627)},
	{NewFloat(-5.5), NewFloat(-0.002685546875), NewFloat(-5.502685546875), false, -12, NewInt(22539)},
	{NewFloat(-0.001708984375), NewFloat(-48.75), NewFloat(-40.751708984375), false, -12, NewInt(199687)},
	{NewFloat(-0.875), NewFloat(-0.25), NewFloat(-1.125), false, -3, NewInt(9)},
	{NewFloat(0.0), NewFloat(1.0), NewFloat(1), true, 0, NewInt(1)},
	{NewFloat(1.0), NewFloat(-1.0), NewFloat(0.0), false, 0, NewInt(0)},
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
	ident string
	a     *Float
	b     *Float
	c     *Float
}{
	{"1", NewFloat(2.0), NewFloat(0.5), NewFloat(1.0)},
	{"2", NewFloat(14.25), NewFloat(3.87), NewFloat(55.1475)},
	{"3", NewFloat(34.98), NewFloat(0.63), NewFloat(22.0374)},
	{"4", NewFloat(3.65), NewFloat(0.115), NewFloat(0.41975)},
	{"5", NewFloat(0.76767676), NewFloat(500.45), NewFloat(384.18383454)},
	{"6", NewFloat(0.214), NewFloat(4.164), NewFloat(0.891096)},
	{"7", NewFloat(0.00134), NewFloat(0.81), NewFloat(0.0010854)},
	{"8", NewFloat(923.83), NewFloat(-6.253), NewFloat(-5776.70899)},
	{"9", NewFloat(92.71), NewFloat(-0.045), NewFloat(-4.17195)},
	{"10", NewFloat(67.77), NewFloat(-0.0017), NewFloat(-0.115209)},
	{"11", NewFloat(0.176), NewFloat(-52.0), NewFloat(-9.152)},
	{"12", NewFloat(0.00095), NewFloat(-9.45), NewFloat(-0.0089775)},
	{"13", NewFloat(0.012), NewFloat(-0.0075), NewFloat(-0.00009)},
	{"14", NewFloat(-93.0), NewFloat(1000.0), NewFloat(-93000.0)},
	{"15", NewFloat(-1.5), NewFloat(0.8), NewFloat(-1.2)},
	{"16", NewFloat(-45.34), NewFloat(0.0097), NewFloat(-0.439798)},
	{"17", NewFloat(-0.077), NewFloat(7800.3333333), NewFloat(-600.6256666641)},
	{"18", NewFloat(-0.0000056), NewFloat(10.99), NewFloat(-0.000061544)},
	{"19", NewFloat(-0.257), NewFloat(0.45), NewFloat(-0.11565)},
	{"20", NewFloat(-75.903), NewFloat(-999.99), NewFloat(75902.24097)},
	{"21", NewFloat(-5044.5), NewFloat(-0.0035), NewFloat(17.65575)},
	{"22", NewFloat(-427.272727), NewFloat(-0.0003), NewFloat(0.1281818181)},
	{"23", NewFloat(-0.0057), NewFloat(-784.665), NewFloat(4.4725905)},
	{"24", NewFloat(-0.000075), NewFloat(-50.89), NewFloat(0.00381675)},
	{"25", NewFloat(-0.99999), NewFloat(-0.000001), NewFloat(0.00000099999)},
	{"26", NewFloat(0.0), NewFloat(1.0), NewFloat(0.0)},
	{"27", NewFloat(10.0), NewFloat(0.0), NewFloat(0.0)},
}

func TestFloatMul(t *testing.T) {
	precision := NewFloat(2)
	precision.exp = precision.exp - 27 //this needs to be changed eventually to a larger number
	for _, testCase := range floatMulTestCases {
		x := testCase.a
		y := testCase.b
		w := testCase.c
		z := x.Mul(y)
		z, w = z.denormalize(w)
		diff := z.Sub(w).Abs()
		bad := diff.Cmp(precision)
		if z.sign != w.sign || z.exp != w.exp || bad == 1 {
			fmt.Printf("ident %v\n%t and %v and %v \n%t and %v and %v\n\n%v dne %v\n", testCase.ident, z.sign, z.exp, z.mantissa, w.sign, w.exp, w.mantissa, z, w)
			fmt.Printf("bad %v, ident %v\n%v\n", bad, testCase.ident, w.Sub(z))
			t.FailNow()
		}
	}
}

var floatDivTestCases = []struct {
	a *Float
	b *Float
	c *Float
}{
	{NewFloat(4.0), NewFloat(2.0), NewFloat(2.0)},
	{NewFloat(20.0), NewFloat(2.0), NewFloat(10.0)},
	{NewFloat(10.0), NewFloat(0.5), NewFloat(20.0)},
	{NewFloat(10.0), NewFloat(12.0), NewFloat(0.83333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333)},
	{NewFloat(0.125), NewFloat(20.0), NewFloat(0.0062500)},
	{NewFloat(0.75), NewFloat(0.25), NewFloat(3.0)},
	{NewFloat(0.25), NewFloat(0.5), NewFloat(0.5)},
	{NewFloat(3454534), NewFloat(-987), NewFloat(-3500.03444782168186423505572441742654508611955420466058763931104356636271529888551165146909827760891590)},
	{NewFloat(92), NewFloat(-500), NewFloat(-0.184)},
	{NewFloat(87), NewFloat(-0.33), NewFloat(-263.6363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636)},
	{NewFloat(0.9045), NewFloat(-45.0), NewFloat(-0.0201)},
	{NewFloat(0.875), NewFloat(-0.2), NewFloat(-4.375)},
	{NewFloat(0.01), NewFloat(-0.74), NewFloat(-0.013513513513513513513513513513513513513513513513513513513513)},
	{NewFloat(-14.0), NewFloat(9.0), NewFloat(-1.55555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555)},
	{NewFloat(-55.55), NewFloat(110.0), NewFloat(-0.505)},
	{NewFloat(-12.5), NewFloat(0.8), NewFloat(-15.625)},
	{NewFloat(-0.00000001), NewFloat(40), NewFloat(-0.00000000025)},
	{NewFloat(-0.625), NewFloat(0.125), NewFloat(-5.0)},
	{NewFloat(-0.45), NewFloat(0.9), NewFloat(-0.5)},
	{NewFloat(-400.0), NewFloat(-87.0), NewFloat(4.597701149425287356321839080459770114942528735632183908045977011494252873563218390804597701149425287356)},
	{NewFloat(-87), NewFloat(-500), NewFloat(0.174)},
	{NewFloat(-8284.45), NewFloat(-0.98), NewFloat(8453.5204081632653061224489795918367346938775510204081632653061224489795918367346938775510204081632653)},
	{NewFloat(-0.23232323), NewFloat(-4.125), NewFloat(0.056320783030303030303030303030303030303030303030303030303030303030303030303030303030303030303030)},
	{NewFloat(-0.7823), NewFloat(-0.11287), NewFloat(6.930982546292194560113404801984584034730220607778860636130061)},
	{NewFloat(-0.234234678), NewFloat(-0.789879234), NewFloat(0.296544924739722933392119028666602469511180996562317525111693)},
}

func TestFloatDiv(t *testing.T) {
	precision := NewFloat(2)
	precision.exp = precision.exp - 30
	for _, testCase := range floatDivTestCases {
		x := testCase.a
		y := testCase.b
		w := testCase.c
		z := x.Div(y)
		z, w = z.denormalize(w)
		diff := z.Sub(w).Abs()
		bad := diff.Cmp(precision)
		if z.sign != w.sign || z.exp != w.exp || bad == 1 {
			fmt.Printf("%t and %v and %v \n%t and %v and %v\n\n%v dne %v\n", z.sign, z.exp, z.mantissa, w.sign, w.exp, w.mantissa, z, w)
			fmt.Printf("%v\n%v\n%v\n", z.Sub(w), bad, precision)
			t.FailNow()
		}
	}
}

func TestFloatDivZero(t *testing.T) {
	x := NewFloat(10.0)
	y := NewFloat(0.0)
	defer func() {
		if r := recover(); r == "division by zero is undefined\n" {
			fmt.Printf("Panicking because %v\n", r)
			fmt.Print("TestFloatDivZero recovers from panic\n")
		} else {
			t.FailNow()
		}
	}()
	x.Div(y)
}

var floatSqrtTestCases = []struct {
	a *Float
	b *Float
}{
	{NewFloat(4.0), NewFloat(2.0)},
	{NewFloat(2.0), NewFloat(1.414213562373095048801688724209698078569671875376948073176679)},
	{NewFloat(0.5), NewFloat(0.707106781186547524400844362104849039284835937688474036588339)},
	{NewFloat(389023489345.2349823489), NewFloat(623717.4755810798847757521684870209731127345501774299345835044)},
	{NewFloat(0.0), NewFloat(0.0)},
	{NewFloat(100.0), NewFloat(10.0)},
	{NewFloat(234901236105.21378943509845), NewFloat(484666.1078569593867853431923603526948650657327099065112700745)},
	{NewFloat(1.0), NewFloat(1.0)},
	{NewFloat(0.999999999999999999), NewFloat(0.999999999999999999499999999999999999874999999999999999937500)},
	{NewFloat(1.000000000000000000000000000001), NewFloat(1.00000000000000000000000000000049999999999999999999999999999987500000000000000000000000000006250000000000000000000000000)},
	{NewFloat(0.25), NewFloat(0.5)},
}

func TestFloatSqrt(t *testing.T) {
	for _, testCase := range floatSqrtTestCases {
		x := testCase.a
		rootx := testCase.b
		z := x.Sqrt()
		z, rootx = z.denormalize(rootx)
		diff := z.Sub(rootx).Abs()
		accuracy := NewFloat(1.0)
		accuracy.exp += x.exp - int64(x.precision) + int64(x.mantissa.BitLen())
		bad := diff.Cmp(accuracy)
		if z.sign != rootx.sign || bad == 1 {
			fmt.Printf("%v dne %v\n", z, rootx)
			fmt.Printf("accu %v\ndiff %v\n", accuracy, diff)
			fmt.Printf("testcase.b %v\n", testCase.b)
			t.FailNow()
		}
	}
}

func TestFloatSqrtNeg(t *testing.T) {
	x := NewFloat(-10.0)
	defer func() {
		if r := recover(); r == "square root of a negative number is undefined\n" {
			fmt.Printf("Panicking because %v\n", r)
			fmt.Print("TestFloatSqrtNeg recovers from panic\n")
		} else {
			t.FailNow()
		}
	}()
	x.Sqrt()
}

/*func TestTheTest(t *testing.T) {
	t.FailNow()
} */

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
