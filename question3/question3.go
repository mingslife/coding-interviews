// Package question3
// 面试题 3：数组中重复的数字
package question3

// 题目一：找出数组中重复的数字。
// 在一个长度为 n 的数组里的所有数字都在 0~n-1的范围内。数组中某些数据是重复的，但不知道有几
// 个数字重复了，也不知道每个数据重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长
// 度为 7 的数组 {2, 3, 1, 0, 2, 5, 3}，那么对应的输出的重复的数字 2 或者 3。

func FindDuplicate(nums []int) int {
	for i, n := range nums {
		if i != n {
			if nums[n] == n {
				return n
			}
			nums[n], nums[i] = n, nums[n]
		}
	}
	return -1
}

// 题目二：不修改数组找出重复的数字。
// 在一个长度为 n+1 的驻足里的所有数字都在 1~n 的范围内，所以数组中至少有一个数字是重复的。
// 请找出数组中任已一个重复的数字，但不能修改输入的数组。例如，如果输入长度为 8 的数组 {2,
// 3, 5, 4, 3, 2, 6, 7}，那么对应的输出是重复的数字 2 或者 3。

func FindDuplicateWithoutModifying(nums []int) int {
	l := len(nums)
	counts := make([]int, l)
	for _, n := range nums {
		counts[n]++
		if counts[n] > 1 {
			return n
		}
	}
	return -1
}
