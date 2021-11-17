// Package question_7
// 面试题 7：重建二叉树
package question_7

import (
	"strconv"
	"strings"
)

// 题目：输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历
// 的结果都不含重复的数字。例如，输入前序遍历序列 {1, 2, 4, 7, 3, 5, 6, 8} 和中序遍历序列
// {4, 7, 2, 1, 5, 3, 8, 6}，则重建如图 2.6 所示的二叉树并输出它的头节点。二叉树的定义如
// 下：
// struct BinaryTreeNode
// {
//     int             m_nValue;
//     BinaryTreeNode* m_pLeft;
//     BinaryTreeNode* m_pRight;
// }
//
//     1
//    / \
//   2   3
//  /   / \
// 4   5   6
//  \     /
//   7   8
//
// 图 2.6 根据前序遍历序列 {1, 2, 4, 7, 3, 5, 6, 8} 和中序遍历序列 {4, 7, 2, 1, 5,
// 3, 8, 6} 重建的二叉树

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func GetBinaryTree(pre, mid []int) *BinaryTreeNode {
	switch len(pre) {
	case 0:
		return nil
	case 1:
		return &BinaryTreeNode{Value: pre[0]}
	}

	value := pre[0]
	pos := 0
	for length := len(mid); pos < length; pos++ {
		if mid[pos] == value {
			break
		}
	}
	preL, preR := pre[1:pos+1], pre[pos+1:]
	midL, midR := mid[0:pos], mid[pos+1:]

	return &BinaryTreeNode{
		Value: value,
		Left:  GetBinaryTree(preL, midL),
		Right: GetBinaryTree(preR, midR),
	}
}

func (node *BinaryTreeNode) String() string {
	builder := &strings.Builder{}
	builder.WriteString("[")
	builder.WriteString(strconv.Itoa(node.Value))
	builder.WriteString(" ")
	if node.Left != nil {
		builder.WriteString(node.Left.String())
	} else {
		builder.WriteString("_")
	}
	builder.WriteString(" ")
	if node.Right != nil {
		builder.WriteString(node.Right.String())
	} else {
		builder.WriteString("_")
	}
	builder.WriteString("]")
	return builder.String()
}
