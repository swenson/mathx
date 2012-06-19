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
