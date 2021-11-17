package question4

import "testing"

func TestSearch(t *testing.T) {
	matrix := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	num1, num2 := 7, 5
	if !Search(matrix, num1) {
		t.Fail()
	}
	if Search(matrix, num2) {
		t.Fail()
	}
}
