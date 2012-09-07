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
  "math/big"
  "fmt"
  "testing"
)


func TestFloatAdd(t *testing.T) {
  x := new(Float)
  y := new(Float)
  z := x.Add(y)
  fmt.Println(z)
  f := big.Int(z.mantissa)
  if f.Sign() == 0 {
    t.FailNow()
  }
}

type A int
type B A

func (x *A) f() int {
  return 42
}

func g(b B) int {
  a := A(b)
  if a.f() == 1 {
    return 2
  }
  if A(b).f() == 1 {
    return 2
  }
  return 1
}

/*func TestGG(t *testing.T) {
  b := new(B)
  if A(b).f() != 42 {
    t.FailNow()
  }
}*/
