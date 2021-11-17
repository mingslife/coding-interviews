package question3

import "testing"

func TestFindDuplicate(t *testing.T) {
	result := FindDuplicate([]int{2, 3, 1, 0, 2, 5, 3})
	if result == 2 || result == 3 {
		return
	}
	t.Fail()
}

func TestFindDuplicateWithoutModifying(t *testing.T) {
	result := FindDuplicateWithoutModifying([]int{2, 3, 5, 4, 3, 2, 6, 7})
	if result == 2 || result == 3 {
		return
	}
	t.Fail()
}
