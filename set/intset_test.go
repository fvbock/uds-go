package set

import (
	"testing"
)

func TestAdd(t *testing.T) {

	s1 := NewIntSet()
	t.Log(s1)
	s1.Add(1)
	t.Log(s1)

	s2 := NewIntSet(2, 3)
	t.Log(s2)

	s3 := NewIntSet([]int{2, 4, 2}...)
	t.Log(s3)
	s3.Remove(1)
	t.Log(s3)
	s3.Remove(2)
	t.Log(s3)

	s4 := NewIntSet([]int{1, 2, 3, 4}...)
	t.Log(s4)
	s4.Remove(2)
	t.Log(s4)
	s4.Remove(4)
	t.Log(s4)
	s4.Remove(1)
	t.Log(s4)
	s4.Remove(3)
	t.Log(s4)

	s5 := s1.Union(s2)
	t.Log("S5", s5)

	s6 := s2.Union(s1)
	t.Log("S6", s6)

	t.Log("S5 == S6", s5.IsEqual(s6))
	t.Log("S6 == S5", s6.IsEqual(s5))

	sA := NewIntSet([]int{2, 4}...)
	sB := NewIntSet([]int{1, 2, 3, 6}...)
	s7 := sA.Intersect(sB)
	t.Log("INTERSECT", sA, sB)
	t.Log("S7", s7)

	sC := NewIntSet([]int{}...)
	sD := NewIntSet([]int{}...)
	s8 := sC.Intersect(sD)
	t.Log("INTERSECT", sC, sD)
	t.Log("S8", s8)

	sE := NewIntSet([]int{2, 4}...)
	sF := NewIntSet([]int{1, 3, 6, 7}...)
	s9 := sE.Intersect(sF)
	t.Log("INTERSECT", sE, sF)
	t.Log("S9", s9)

	s10 := sF.Intersect(sE)
	t.Log("INTERSECT", sF, sE)
	t.Log("S10", s10)

	// s11 := sG.Intersect(sH)
	// t.Log("INTERSECT", sG.Len(), sH.Len())
	// t.Log("S11", s11, s11.Len())

	sI := NewIntSet([]int{2, 4}...)
	sJ := NewIntSet([]int{1, 2, 4, 7}...)
	s12 := sI.Difference(sJ)
	t.Log("DIFFERENCE", sI, sJ)
	t.Log("S12", s12)

	sK := NewIntSet([]int{}...)
	sL := NewIntSet([]int{1, 2, 4, 7}...)
	s13 := sK.Difference(sL)
	t.Log("DIFFERENCE", sK, sL)
	t.Log("S13", s13)

	sM := NewIntSet([]int{1, 2}...)
	sN := NewIntSet([]int{1, 2, 3}...)
	sO := NewIntSet([]int{1, 2, 4, 7}...)

	t.Log("SM", sM)
	t.Log("SN", sN)
	t.Log("SO", sO)
	t.Log("sM isSubset of sN", sM.IsSubset(sN))
	t.Log("sM isSubset of sO", sM.IsSubset(sO))
	t.Log("sN isSubset of sO", sN.IsSubset(sO))
	t.Log("sN isSuperset of sM", sN.IsSuperset(sM))
	t.Log("sO isSuperset of sM", sM.IsSuperset(sM))

	t.Log("sO hasMember 2", sO.HasMember(2))
	t.Log("sO hasMember 6", sO.HasMember(6))
}
