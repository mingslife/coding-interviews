// Package question_6
// 面试 6：从尾到头打印链表
package question_6

import (
	"strconv"
)

// 题目：输入一个链表的头节点，从尾到头反过来打印出每个节点的值。链表节点定义如下：
// struct ListNode
// {
//     int       m_nKey;
//     ListNode* m_pNext;
// }

type ListNode struct {
	Key  int
	Next *ListNode
}

func ReverseKey(node *ListNode) string {
	if node.Next == nil {
		return strconv.Itoa(node.Key)
	}

	return ReverseKey(node.Next) + strconv.Itoa(node.Key)
}
