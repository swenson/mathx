package mathx

import (
	"fmt"
	"math/big"
	"testing"
)

var floatSqrtTestCases = []struct {
	a string
	b string
}{
	{"4.0", "2.0"},
	{"2.0", "1.414213562373095048801688724209698078569671875376948073176679"},
	{"0.5", "0.707106781186547524400844362104849039284835937688474036588339"},
	{"389023489345.2349823489", "623717.4755810798847757521684870209731127345501774299345835044"},
	{"0.0", "0.0"},
	{"100.0", "10.0"},
	{"234901236105.21378943509845", "484666.1078569593867853431923603526948650657327099065112700745"},
	{"1.0", "1.0"},
	{"0.999999999999999999", "0.999999999999999999499999999999999999874999999999999999937500"},
	{"1.000000000000000000000000000001", "1.00000000000000000000000000000049999999999999999999999999999987500000000000000000000000000006250000000000000000000000000"},
	{"0.25", "0.5"},
}

func TestFloatSqrt(t *testing.T) {
	for _, testCase := range floatSqrtTestCases {
		x, _, _ := ParseFloat(testCase.a, 10, 256, big.ToNearestEven)
		rootx, _, _ := ParseFloat(testCase.b, 10, 256, big.ToNearestEven)
		z := x.Sqrt()
		diff := z.Sub(rootx).Abs()
		_, xexp := x.MantExp()
		accuracy := NewFloat(1.0).SetExp(53 + xexp - int(x.Prec()) + 53)
		bad := diff.Cmp(accuracy)
		if z.Sign() != rootx.Sign() || bad == 1 {
			fmt.Printf("%v dne %v\n", z, rootx)
			fmt.Printf("accu %v\ndiff %v\n", accuracy, diff)
			fmt.Printf("testcase.b %v\n", testCase.b)
			t.FailNow()
		}
	}
}
