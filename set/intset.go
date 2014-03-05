package set

import (
	"fmt"
)

type IntSet struct {
	m          map[int]int
	memberList []int
}

func (s *IntSet) String() string {
	return fmt.Sprintf("<%v>", s.memberList)
}

func NewIntSet(elements ...int) (s *IntSet) {
	s = &IntSet{m: make(map[int]int)}
	s.Add(elements...)
	return
}

func (s *IntSet) Len() int {
	return len(s.memberList)
}

func (s *IntSet) Members() []int {
	return s.memberList
}

func (s *IntSet) Add(elements ...int) {
	for _, e := range elements {
		if _, notSet := s.m[e]; !notSet {
			s.memberList = append(s.memberList, e)
			s.m[e] = len(s.memberList) - 1
		}
	}
	return
}

func (s *IntSet) Remove(elements ...int) {
	for _, e := range elements {
		if idx, isSet := s.m[e]; isSet {
			if idx == 0 {
				s.memberList = s.memberList[idx+1:]
			} else if idx == len(s.memberList)-1 {
				s.memberList = s.memberList[:idx]
			} else {
				s.memberList = append(s.memberList[:idx], s.memberList[idx+1:]...)
			}
			for i := idx; i < len(s.memberList); i++ {
				s.m[s.memberList[i]] = s.m[s.memberList[i]] - 1
			}
			delete(s.m, e)
		}
	}
	return
}

func (s *IntSet) Union(s2 *IntSet) *IntSet {
	us := NewIntSet()
	us.Add(s.memberList...)
	us.Add(s2.memberList...)
	return us
}

func (s *IntSet) Intersect(s2 *IntSet) *IntSet {
	is := NewIntSet()
	if s.Len() == 0 || s2.Len() == 0 {
		return is
	}

	if s.Len() > s2.Len() {
		for _, elm := range s.memberList {
			if _, exists := s2.m[elm]; exists {
				is.Add(elm)
			}
		}
	} else {
		for _, elm := range s2.memberList {
			if _, exists := s.m[elm]; exists {
				is.Add(elm)
			}
		}
	}
	return is
}

func (s *IntSet) Difference(s2 *IntSet) *IntSet {
	ds := NewIntSet()
	if s.Len() == 0 {
		ds.Add(s2.memberList...)
		return ds
	}
	if s2.Len() == 0 {
		ds.Add(s.memberList...)
		return ds
	}

	for _, elm := range s.memberList {
		if !s2.HasMember(elm) {
			ds.Add(elm)
		}
	}
	for _, elm := range s2.memberList {
		if !s.HasMember(elm) {
			ds.Add(elm)
		}
	}
	return ds
}

func (s *IntSet) IsSubset(s2 *IntSet) bool {
	for _, elm := range s.memberList {
		if _, notExists := s2.m[elm]; !notExists {
			return false
		}
	}
	return true
}

func (s *IntSet) IsSuperset(s2 *IntSet) bool {
	return s2.IsSubset(s)
}

func (s *IntSet) IsEqual(s2 *IntSet) bool {
	if s.Len() != s2.Len() {
		return false
	}
	return s.IsSubset(s2)
}

func (s *IntSet) HasMember(elm int) bool {
	if _, exists := s.m[elm]; exists {
		return true
	}
	return false
}

func (s *IntSet) HasMembers(elements ...int) bool {
	ms := NewIntSet(elements...)
	return ms.IsSubset(s)
}
