// Copyright (c) 2014 Christopher Swenson.
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

package mathx

// Set is a basic set of ints.
type Set struct {
	elements map[int]bool
}

// NewIntSet creates a new set of integers from the given slice.
func NewIntSet(nums []int) *Set {
	s := new(Set)
	s.elements = make(map[int]bool)
	for _, n := range nums {
		s.elements[n] = true
	}
	return s
}

// Contains returns true of the number is in the set.
func (s *Set) Contains(num int) bool {
	_, ok := s.elements[num]
	return ok
}
