/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
package leetcodego

import "sort"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	first := sort.Search(n, func(i int) bool {
		return nums[i] >= target
	})
	if first == n || nums[first] != target { // 未查找到target
		return []int{-1, -1}
	}
	second := sort.Search(n, func(i int) bool {
		return nums[i] > target
	})
	return []int{first, second - 1}
}

// @lc code=end
