/*
 * @lc app=leetcode.cn id=162 lang=golang
 *
 * [162] 寻找峰值
 */

// @lc code=start
package leetcodego

import "math"

func findPeakElement(nums []int) int {
	n := len(nums)
	get := func(ind int) int {
		if ind == -1 || ind == n {
			return math.MinInt
		} else {
			return nums[ind]
		}
	}
	left, right := 0, n-1
	for {
		mid := (left + right) >> 1
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		} else {
			// 如果递增，右侧必然有峰值
			if get(mid) < get(mid+1) {
				left = mid + 1
			} else {
				// 同理，如果递减，则左侧必然有峰值
				right = mid - 1
			}
		}
	}
}

// @lc code=end
