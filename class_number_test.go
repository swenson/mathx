package ntag

import (
  "testing"
)

func TestMain(t *testing.T) {
  class_number_1_nums := []int{-3, -4, -7, -8, -11, -19, -43, -67, -163}
  class_number_1 := IntSetFromSlice(class_number_1_nums)
  for a := 1; a < 10; a++ {
    for b := -10; b < 10; b++ {
      for c := -10; c < 10; c++ {
        p := MakeIntPolynomial64(int64(c), int64(b), int64(a))
        if !p.IsIrreducible() {
          continue
        }
        if p.Discriminant().Sign() > 0 {
          continue
        }
        if !IsFundamentalDiscriminant(p.Discriminant()) {
          continue
        }
        k := MakeNumberField(p)

        if k.ClassNumber() == 1 && !class_number_1.Contains(int(p.Discriminant().Int64())) {
          t.Fail()
        }
      }
    }
  }
}
