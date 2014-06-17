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
  x.sign = true
  x.exp = 1
  x.mantissa = (*Int)(big.NewInt(1))
  y := new(Float)
  y.sign = true
  y.exp = 1
  y.mantissa = (*Int)(big.NewInt(1))
  z := x.Add(y)

  one := (*Int)(big.NewInt(1))
  if z.mantissa.Cmp(one) != 0 || z.exp != 2 || !z.sign {
    fmt.Printf("%s + %s = %s\n", x, y, z)
    t.FailNow()
  }
}
