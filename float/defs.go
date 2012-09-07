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
)

type RoundingMode int

const (
	_ = iota
	RoundUp RoundingMode = 1 * iota
	RoundDown
)

type Int big.Int

type Float struct {
	sign     bool
	exp      int64
  mantissa Int
}

func (x *Float) Add(y *Float) *Float {
  z := new(Float)
	return z
}
