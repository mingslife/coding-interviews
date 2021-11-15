// Package question_4
// 面试题 4：二维数组中的查找
package question_4

// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序
// 排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func Search(matrix [][]int, num int) bool {
	for _, ys := range matrix {
		for _, current := range ys {
			if current == num {
				return true
			} else if current > num {
				break
			}
		}
	}
	return false
}
