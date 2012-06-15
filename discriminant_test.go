package ntag

import (
  "fmt"
  "math/big"
  "testing"
)

func TestFundamentalDiscriminantsSmall(t *testing.T) {
  fundamentals := []int{1, 5, 8, 12, 13, 17, 21, 24, 28, 29, 33, -3, -4, -7, -8, -11, -15, -19, -20, -23, -24, -31}
  good := IntSetFromSlice(fundamentals)

  for d := -31; d <= 33; d++ {
    a := IsFundamentalDiscriminant(big.NewInt(int64(d)))
    b := good.Contains(d)
    if a != b {
      fmt.Println("Fundamental discriminant failed for ", d)
      t.Fail()
    }
  }
}
