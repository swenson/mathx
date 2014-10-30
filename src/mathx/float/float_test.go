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
	//Can not use Cmp() for testing because it relies on Sub() and consquently Add(), which would be circular logic
	{NewFloat(1.125), NewFloat(40), NewFloat(41.125), true, -3, NewInt(329)},
	{NewFloat(2), NewFloat(0.0009765625), NewFloat(2.0009765625), true, -10, NewInt(2049)},
	{NewFloat(0.0087890625), NewFloat(1.75), NewFloat(1.7587890625), true, -10, NewInt(1801)},
	{NewFloat(0.28125), NewFloat(0.46875), NewFloat(0.75), true, -2, NewInt(3)},
	{NewFloat(5000), NewFloat(-4099), NewFloat(901), true, 0, NewInt(901)},
	{NewFloat(2), NewFloat(-0.05859375), NewFloat(1.94140625), true, -8, NewInt(497)},
	{NewFloat(0.00244140625), NewFloat(-20), NewFloat(-19.99755859375), false, -11, NewInt(40955)},
	{NewFloat(0.03125), NewFloat(-0.375), NewFloat(-0.34375), false, -5, NewInt(11)},
	{NewFloat(-1.1), NewFloat(2.6), NewFloat(1.5), true, -1, NewInt(3)},
	//{NewFloat(-), NewFloat(0.0015869140625), NewFloat(), , ,NewInt()},
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
	a *Float
	b *Float
	c *Float //I don't think this variable does anything
}{
	{NewFloat(4.0), NewFloat(2.0), NewFloat(2.0)},
	{NewFloat(20.0), NewFloat(2.0), NewFloat(10.0)},
	{NewFloat(10.0), NewFloat(0.5), NewFloat(20.0)},
	{NewFloat(10.0), NewFloat(12.0), NewFloat(0.833333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333)},
	{NewFloat(0.125), NewFloat(20.0), NewFloat(0.0062500)}, //aBc ++
	{NewFloat(0.75), NewFloat(0.25), NewFloat(3.0)},        //abC ++
	{NewFloat(0.25), NewFloat(0.5), NewFloat(0.5)},         //abc ++
	{NewFloat(3454534), NewFloat(-987), NewFloat(-3500.034447821681864235055724417426545086119554204660587639311043566362715298885511651469098277608915906788247213779128672745694022289766970618034447821681864235055724417426545086119554204660587639311043566362715298885511651469098277608915906788247213779128672745694022289766970618034447821681864235055724417426545086119554204660587639311043566362715298885511651469098277608915906788247213779128672745694022289766970618034447821681864235055724417426545086119554204660587639311043566362)}, //ABC +-
	{NewFloat(92), NewFloat(-500), NewFloat(-0.184)}, //ABc +-
	{NewFloat(87), NewFloat(-0.33), NewFloat(-263.636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636363636)}, //AbC +-
	{NewFloat(0.9045), NewFloat(-45.0), NewFloat(-0.0201)}, //aBc +-
	{NewFloat(0.875), NewFloat(-0.2), NewFloat(-4.375)},    //abC +-
	{NewFloat(0.74), NewFloat(-0.01), NewFloat(-0.0074)},   //abc +-
	{NewFloat(-14.0), NewFloat(9.0), NewFloat(-1.55555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555555)}, //ABC -+
	{NewFloat(-55.55), NewFloat(110.0), NewFloat(-0.505)},           //ABc -+
	{NewFloat(-12.5), NewFloat(0.8), NewFloat(-15.625)},             //AbC -+
	{NewFloat(-0.00000001), NewFloat(40), NewFloat(-0.00000000025)}, //aBc -+
	{NewFloat(-0.625), NewFloat(0.125), NewFloat(-5.0)},             //abC -+
	{NewFloat(-0.45), NewFloat(0.9), NewFloat(-0.5)},                //abc -+
	{NewFloat(-400.0), NewFloat(-87.0), NewFloat(4.59770114942528735632183908045977011494252873563218390804597701149425287356321839080459770114942528735632183908045977011494252873563218390804597701149425287356321839080459770114942528735632183908045977011494252873563218390804597701149425287356321839080459770114942528735632183908045977011494252873563218390804597701149425287356321839080459770114942528735632183908045977011494252873563218390804597701149425287356321839080459770114942528735632183908045977011494252873563218390804597701149)}, //ABC --
	{NewFloat(-87), NewFloat(-500), NewFloat(0.174)}, //ABc --
	{NewFloat(-8284.45), NewFloat(-0.98), NewFloat(8453.52040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244897959183673469387755102040816326530612244)},           //AbC --
	{NewFloat(-0.23232323), NewFloat(-4.125), NewFloat(0.05632078303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303)},       //aBc --
	{NewFloat(-0.7823), NewFloat(-0.11287), NewFloat(6.92832462124568087179941525648976698857092229999114024984495437228670151501727651280233897404093204571631080003543900062018251085319393993089394879064410383627181713475679985824399751926995658722424027642420483742358465491273146097280056702400992292017365110303889430318065030566138034907415610879773190396030831930539558784442278727739877735447860370337556480907238415876672277841764862230885089040489058208558518649774076371046336493310888632940551076459643838043767165765925400903694)},         //abC --
	{NewFloat(-0.234234678), NewFloat(-0.789879234), NewFloat(0.2965787173543568460443207307634684724000681933833501711166652849056182500456824408928319222948536672506731080611311548974596430115430011040135297666009967413660642530358928645512531764898684349500768201152764764073752871642714744830264424698399978442856312726086109196064326194651572067390323530290107752975523593931390561345636170067114386621279266513779801662882631798447634954720084219329837053717042691417594076750861357018700510673205758082035339570101053061455127406813955126666)}, //abc --
}

//don't foreget to add a divide by zero

func TestFloatDiv(t *testing.T) {
	precision := NewFloat(2)
	precision.exp = precision.exp - 50
	for _, testCase := range floatDivTestCases { //floatDivTestCases or floatMulTestCAses?
		fmt.Printf("\u001b[2J") //this will clear stdout so the failure will only print the failed iteration and not all pervious testcases
		x := testCase.a
		y := testCase.b
		w := testCase.c
		z := x.Div(y)
		z, w = z.denormalize(w)
		diff := z.Sub(w).Abs()
		yes := diff.Cmp(precision)
		if z.sign != w.sign || z.exp != w.exp || yes <= 0 {
			fmt.Printf("%t and %v and %v \n%t and %v and %v\n\n%v dne %v\n", z.sign, z.exp, z.mantissa, w.sign, w.exp, w.mantissa, z, w)
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
