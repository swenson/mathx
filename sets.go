package ntag

type Set struct {
  elements map[int]bool
}

func IntSetFromSlice(nums []int) *Set {
  s := new(Set)
  s.elements = make(map[int]bool)
  for _, n := range nums {
    s.elements[n] = true
  }
  return s
}

func (s *Set) Contains(num int) bool {
  _, ok := s.elements[num]
  return ok
}
