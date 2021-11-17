package question7

import (
	"testing"
)

func TestGetBinaryTreeNode(t *testing.T) {
	tree := GetBinaryTree([]int{1, 2, 4, 7, 3, 5, 6, 8}, []int{4, 7, 2, 1, 5, 3, 8, 6})
	if tree.String() != "[1 [2 [4 _ [7 _ _]] _] [3 [5 _ _] [6 [8 _ _] _]]]" {
		t.Fail()
	}
}
