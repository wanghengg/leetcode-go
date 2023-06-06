/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
package leetcodego

import "sort"

func nextPermutation(nums []int) {
	first, second := -1, 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] < nums[i+1] {
			first = i
			for j := i + 1; j < n; j++ {
				if nums[i] < nums[j] {
					second = j
				}
			}
		}
	}
	if first == -1 {
		sort.Ints(nums)
	} else {
		nums[first], nums[second] = nums[second], nums[first]
		sort.Ints(nums[first+1:])
	}
}

// @lc code=end
