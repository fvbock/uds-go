package set

import (
	"fmt"
)

type StringSet struct {
	m       map[string]int
	members []string
}

func (s *StringSet) String() string {
	return fmt.Sprintf("<%v>", s.members)
}

func NewStringSet(elements ...string) (s *StringSet) {
	s = &StringSet{m: make(map[string]int)}
	s.Add(elements...)
	return
}

func (s *StringSet) Len() int {
	return len(s.members)
}

func (s *StringSet) Members() []string {
	return s.members
}

func (s *StringSet) Add(elements ...string) {
	for _, e := range elements {
		if _, notSet := s.m[e]; !notSet {
			s.members = append(s.members, e)
			s.m[e] = len(s.members) - 1
		}
	}
	return
}

func (s *StringSet) Remove(elements ...string) {
	for _, e := range elements {
		if idx, isSet := s.m[e]; isSet {
			if idx == 0 {
				s.members = s.members[idx+1:]
			} else if idx == len(s.members)-1 {
				s.members = s.members[:idx]
			} else {
				s.members = append(s.members[:idx], s.members[idx+1:]...)
			}
			for i := idx; i < len(s.members); i++ {
				s.m[s.members[i]] = s.m[s.members[i]] - 1
			}
			delete(s.m, e)
		}
	}
	return
}

func (s *StringSet) Union(s2 *StringSet) *StringSet {
	us := NewStringSet()
	us.Add(s.members...)
	us.Add(s2.members...)
	return us
}

func (s *StringSet) Intersect(s2 *StringSet) *StringSet {
	is := NewStringSet()
	if s.Len() == 0 || s2.Len() == 0 {
		return is
	}

	if s.Len() > s2.Len() {
		for _, elm := range s.members {
			if _, exists := s2.m[elm]; exists {
				is.Add(elm)
			}
		}
	} else {
		for _, elm := range s2.members {
			if _, exists := s.m[elm]; exists {
				is.Add(elm)
			}
		}
	}
	return is
}

func (s *StringSet) Difference(s2 *StringSet) *StringSet {
	ds := NewStringSet()
	if s.Len() == 0 {
		ds.Add(s2.members...)
		return ds
	}
	if s2.Len() == 0 {
		ds.Add(s.members...)
		return ds
	}

	for _, elm := range s.members {
		if !s2.HasMember(elm) {
			ds.Add(elm)
		}
	}
	for _, elm := range s2.members {
		if !s.HasMember(elm) {
			ds.Add(elm)
		}
	}
	return ds
}

func (s *StringSet) IsSubset(s2 *StringSet) bool {
	for _, elm := range s.members {
		if _, notExists := s2.m[elm]; !notExists {
			return false
		}
	}
	return true
}

func (s *StringSet) IsSuperset(s2 *StringSet) bool {
	return s2.IsSubset(s)
}

func (s *StringSet) IsEqual(s2 *StringSet) bool {
	if s.Len() != s2.Len() {
		return false
	}
	return s.IsSubset(s2)
}

func (s *StringSet) HasMember(elm string) bool {
	if _, exists := s.m[elm]; exists {
		return true
	}
	return false
}

func (s *StringSet) HasMembers(elements ...string) bool {
	ms := NewStringSet(elements...)
	return ms.IsSubset(s)
}
