package slice

import "testing"

func TestDeleteFirstInt(t *testing.T) {
	is := []int{1, 2, 3}
	t.Log(is)
	is = DeleteIntSlice(is, 0)
	t.Log(is)

	if len(is) != 2 {
		t.Errorf("Expected slice to be 2 elements long. Got %v elements: %v", len(is), is)
	}
	if is[0] != 2 {
		t.Errorf("Expected new first element to be 2. Got %v", is[0])
	}
	if is[1] != 3 {
		t.Errorf("Expected new first element to be 3. Got %v", is[1])
	}

	is = DeleteIntSlice(is, 0)
	t.Log(is)

	if len(is) != 1 {
		t.Errorf("Expected slice to be 1 elements long. Got %v elements: %v", len(is), is)
	}
	if is[0] != 3 {
		t.Errorf("Expected new first element to be 3. Got %v", is[0])
	}

	is = DeleteIntSlice(is, 0)
	t.Log(is)

	if len(is) != 0 {
		t.Errorf("Expected slice to be empty. Got %v elements: %v", len(is), is)
	}
}

func TestDeleteLastInt(t *testing.T) {
	is := []int{1, 2, 3}
	t.Log(is)
	is = DeleteIntSlice(is, 2)
	t.Log(is)

	if len(is) != 2 {
		t.Errorf("Expected slice to be 2 elements long. Got %v elements: %v", len(is), is)
	}
	if is[0] != 1 {
		t.Errorf("Expected new first element to be 1. Got %v", is[0])
	}
	if is[1] != 2 {
		t.Errorf("Expected new first element to be 2. Got %v", is[1])
	}

	is = DeleteIntSlice(is, 1)
	t.Log(is)

	if len(is) != 1 {
		t.Errorf("Expected slice to be 1 elements long. Got %v elements: %v", len(is), is)
	}
	if is[0] != 1 {
		t.Errorf("Expected new first element to be 1. Got %v", is[0])
	}

	is = DeleteIntSlice(is, 0)
	t.Log(is)

	if len(is) != 0 {
		t.Errorf("Expected slice to be empty. Got %v elements: %v", len(is), is)
	}
}

func TestDeleteInnerInt(t *testing.T) {
	is := []int{1, 2, 3}
	t.Log(is)
	is = DeleteIntSlice(is, 1)
	t.Log(is)

	if len(is) != 2 {
		t.Errorf("Expected slice to be 2 elements long. Got %v elements: %v", len(is), is)
	}
	if is[0] != 1 {
		t.Errorf("Expected new first element to be 1. Got %v", is[0])
	}
	if is[1] != 3 {
		t.Errorf("Expected new first element to be 3. Got %v", is[1])
	}
}

func TestDeleteIntNonexistentIndex(t *testing.T) {
	is := []int{1, 2, 3}
	t.Log(is)
	is = DeleteIntSlice(is, -1)
	t.Log(is)

	// make sure is is unaffected deletion with invalid indexes should do nothing
	if len(is) != 3 {
		t.Errorf("Expected slice to be unaffected - and 3 elements long. Got %v elements: %v", len(is), is)
	}
	if is[0] != 1 {
		t.Errorf("Expected the first element to be 1. Got %v", is[0])
	}
	if is[1] != 2 {
		t.Errorf("Expected the second element to be 2. Got %v", is[1])
	}
	if is[2] != 3 {
		t.Errorf("Expected the third element to be 3. Got %v", is[2])
	}

	t.Log(is)
	is = DeleteIntSlice(is, 3)
	t.Log(is)

	// make sure is is unaffected deletion with invalid indexes should do nothing
	if len(is) != 3 {
		t.Errorf("Expected slice to be unaffected - and 3 elements long. Got %v elements: %v", len(is), is)
	}
	if is[0] != 1 {
		t.Errorf("Expected the first element to be 1. Got %v", is[0])
	}
	if is[1] != 2 {
		t.Errorf("Expected the second element to be 2. Got %v", is[1])
	}
	if is[2] != 3 {
		t.Errorf("Expected the third element to be 3. Got %v", is[2])
	}

}
