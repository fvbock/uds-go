package set

import (
	"fmt"
)

type Int64Set struct {
	m          map[int64]int64
	memberList []int64
}

func (s *Int64Set) String() string {
	return fmt.Sprintf("<%v>", s.memberList)
}

func NewInt64Set(elements ...int64) (s *Int64Set) {
	s = &Int64Set{m: make(map[int64]int64)}
	s.Add(elements...)
	return
}

func (s *Int64Set) Len() int {
	return len(s.memberList)
}

func (s *Int64Set) Members() []int64 {
	return s.memberList
}

func (s *Int64Set) Add(elements ...int64) {
	for _, e := range elements {
		if _, notSet := s.m[e]; !notSet {
			s.memberList = append(s.memberList, e)
			s.m[e] = int64(len(s.memberList) - 1)
		}
	}
	return
}

func (s *Int64Set) Remove(elements ...int64) {
	for _, e := range elements {
		if idx, isSet := s.m[e]; isSet {
			if idx == 0 {
				s.memberList = s.memberList[idx+1:]
			} else if idx == int64(len(s.memberList)-1) {
				s.memberList = s.memberList[:idx]
			} else {
				s.memberList = append(s.memberList[:idx], s.memberList[idx+1:]...)
			}
			for i := idx; i < int64(len(s.memberList)); i++ {
				s.m[s.memberList[i]] = s.m[s.memberList[i]] - 1
			}
			delete(s.m, e)
		}
	}
	return
}

func (s *Int64Set) Union(s2 *Int64Set) *Int64Set {
	us := NewInt64Set()
	us.Add(s.memberList...)
	us.Add(s2.memberList...)
	return us
}

func (s *Int64Set) Intersect(s2 *Int64Set) *Int64Set {
	is := NewInt64Set()
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

func (s *Int64Set) Difference(s2 *Int64Set) *Int64Set {
	ds := NewInt64Set()
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

func (s *Int64Set) IsSubset(s2 *Int64Set) bool {
	for _, elm := range s.memberList {
		if _, notExists := s2.m[elm]; !notExists {
			return false
		}
	}
	return true
}

func (s *Int64Set) IsSuperset(s2 *Int64Set) bool {
	return s2.IsSubset(s)
}

func (s *Int64Set) IsEqual(s2 *Int64Set) bool {
	if s.Len() != s2.Len() {
		return false
	}
	return s.IsSubset(s2)
}

func (s *Int64Set) HasMember(elm int64) bool {
	if _, exists := s.m[elm]; exists {
		return true
	}
	return false
}

func (s *Int64Set) HasMembers(elements ...int64) bool {
	ms := NewInt64Set(elements...)
	return ms.IsSubset(s)
}
